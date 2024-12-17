ALTER TABLE "order_items"
    DROP CONSTRAINT order_items_pk

ALTER TABLE "order_items"
    DROP CONSTRAINT fk_order_items_order

ALTER TABLE "order_items"
    DROP CONSTRAINT fk_order_items_item

DROP TABLE IF EXISTS "order_items"