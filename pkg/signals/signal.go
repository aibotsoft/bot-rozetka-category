package signals

import (
	"os"
	"os/signal"
)

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func SetupSignalHandler() (stopCh <-chan os.Signal) {
	c := make(chan os.Signal)
	signal.Notify(c, shutdownSignals...)
	return c
}
