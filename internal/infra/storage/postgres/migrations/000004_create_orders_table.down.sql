ALTER TABLE "orders"
    DROP CONSTRAINT orders_pk

ALTER TABLE "orders"
    DROP CONSTRAINT fk_orders_client

DROP TYPE IF EXISTS "orders_status_enum";

DROP TABLE IF EXISTS "orders"