ALTER TABLE "products"
    DROP CONSTRAINT products_pk

ALTER TABLE "products"
    DROP CONSTRAINT fk_products_category

DROP TABLE IF EXISTS "products"