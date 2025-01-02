ALTER TABLE "order_products"
    DROP CONSTRAINT order_products_pk

ALTER TABLE "order_products"
    DROP CONSTRAINT fk_order_products_order

ALTER TABLE "order_products"
    DROP CONSTRAINT fk_order_products_product

DROP TABLE IF EXISTS "order_products"