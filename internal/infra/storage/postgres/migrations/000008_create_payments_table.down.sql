ALTER TABLE "orders"
    DROP CONSTRAINT fk_orders_payment

DROP TYPE IF EXISTS "payments_type_enum";

DROP TABLE IF EXISTS "payments";

