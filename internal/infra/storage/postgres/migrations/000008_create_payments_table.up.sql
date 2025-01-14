CREATE TYPE "payments_type_enum" AS ENUM ('PIX-QRCODE');

CREATE TABLE "payments" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "type" payments_type_enum NOT NULL,
  "provider" varchar NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL,
	CONSTRAINT payments_pk PRIMARY KEY (id)
);

ALTER TABLE "orders"
    ADD CONSTRAINT fk_orders_payment
      FOREIGN KEY ("payment_id")
    REFERENCES "payments" ("id")
      ON DELETE SET NULL
      ON UPDATE CASCADE;