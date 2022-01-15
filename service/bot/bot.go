package bot

import (
	"context"
	"fmt"
	"github.com/aibotsoft/bot-rozetka-category/pkg/config"
	"github.com/aibotsoft/bot-rozetka-category/pkg/store"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"net/url"
	"strconv"
)

type Bot struct {
	log   *zap.Logger
	cfg   *config.Config
	ctx   context.Context
	store *store.Store
	bot   *tgbotapi.BotAPI
}

const helpMsg = "Subscribe to receive notifications of new discounted products in online stores.\n" +
	"/subscribe - subscribe to receive notifications\n" +
	"/unsubscribe - unsubscribe and stop receive notifications\n" +
	"/help - get help\n" +
	"/status - get bot status\n"

func NewBot(log *zap.Logger, cfg *config.Config, ctx context.Context, sto *store.Store) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return nil, err
	}
	bot.Debug = cfg.TelegramDebug
	return &Bot{
		log:   log,
		cfg:   cfg,
		ctx:   ctx,
		bot:   bot,
		store: sto,
	}, nil
}

func (b *Bot) Send(id int64, text string) error {
	msg := tgbotapi.NewMessage(id, text)
	msg.ParseMode = tgbotapi.ModeHTML
	_, err := b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("send_error: %w", err)
	}
	return nil
}
func (b *Bot) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 120

	updates := b.bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		// Extract the command from the Message.
		switch update.Message.Command() {
		case "start":
			msg.Text = "Hello. Use /help to get help or /subscribe to subscribe"
			b.log.Info("start", zap.Any("chat", update.Message.Chat))
		case "subscribe":
			msg.Text = fmt.Sprintf("Hello %s. You are subscribed", update.Message.Chat.UserName)
			b.log.Info("subscribe", zap.Any("chat", update.Message.Chat))
			u := store.User{
				ID:        update.Message.Chat.ID,
				FirstName: update.Message.Chat.FirstName,
				UserName:  update.Message.Chat.UserName,
				Discount:  30,
				Rating:    4,
			}

			err := b.store.SaveUser(&u)
			if err != nil {
				b.log.Warn("save_user_error", zap.Error(err))
			}
		case "unsubscribe":
			msg.Text = "goodbye"
			b.log.Info("stop", zap.Any("chat", update.Message.Chat))
			err := b.store.DeleteUser(update.Message.Chat.ID)
			if err != nil {
				b.log.Warn("delete_user_error", zap.Error(err))
			}
		case "set":
			b.log.Info("discount", zap.Any("chat", update.Message.Text))

		case "discount":
			query, err := url.ParseQuery(update.Message.Text)
			if err != nil {
				msg.Text = "parse_discount_error"
				b.log.Info("parse_discount_error", zap.Any("msg", update.Message))
				continue
			}
			discountStr := query.Get("/discount")
			msg.Text = fmt.Sprintf("set discount to %s", discountStr)
			discount, err := strconv.ParseInt(discountStr, 10, 64)
			if err != nil {
				msg.Text = "parse_discount_error"
				b.log.Info("parse_discount_error", zap.Any("msg", update.Message))
				continue
			}
			b.log.Info("parse_discount", zap.String("discountStr", discountStr), zap.Int64("discount", discount))

		case "help":
			msg.Text = helpMsg
		case "status":
			msg.Text = "I'm ok."
		default:
			msg.Text = "I don't know that command, use /help to get help"
		}

		if _, err := b.bot.Send(msg); err != nil {
			b.log.Warn("send_error", zap.Error(err))
		}
	}
}
