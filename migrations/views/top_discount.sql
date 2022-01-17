create or replace view top_discount as
select sp.id,
       (100 - sp.price * 100 / op.price)::int discount,
       sp.sell_status,
       sp.updated_at,
       sp.title,
       c.title                                category_title,
       sp.price,
       op.price                               origin_price,
--        sp.href,
--        sp.brand,
--        substring(sp.docket from 52 for position('</' in sp.docket) - 52) reason,
--        sp.category_id,
--        sp.sell_status,
--        sp.docket,
--        op.id,
--        op.title,
--        op.category_id,
--        op.sell_status,
--        op.href,
--        op.old_price,
--        op.brand,
--        stars,
       op.comments_mark,
       op.comments_amount,
       sp.href
from sale_products sp
         join origin_products op on sp.origin_id = op.id
         join categories c on sp.category_id = c.id
where op.comments_amount > 17
  and op.comments_mark > 4
--   and sp.sell_status != 'out_of_stock'
  and (100-sp.price * 100 / op.price) > 30
order by 2 desc;


drop view top_discount;

