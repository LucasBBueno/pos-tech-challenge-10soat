CREATE TYPE "orders_status_enum" AS ENUM ('received', 'preparing', 'ready', 'completed');

CREATE TABLE IF NOT EXISTS "orders" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "status" orders_status_enum DEFAULT 'received',
	"client_id" uuid NULL,
	"total" numeric(10, 2) NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL,
	CONSTRAINT orders_pk PRIMARY KEY (id)
);

ALTER TABLE "orders"
      ADD CONSTRAINT fk_orders_client FOREIGN KEY (client_id) 
          REFERENCES "clients" (id);