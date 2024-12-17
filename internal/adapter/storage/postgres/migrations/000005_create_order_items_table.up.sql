CREATE TABLE IF NOT EXISTS "order_items" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "order_id" uuid NOT NULL,
	"item_id" uuid NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL,
	CONSTRAINT order_items_pk PRIMARY KEY (id)
);

ALTER TABLE "order_items"
      ADD CONSTRAINT fk_order_items_order FOREIGN KEY (order_id) 
          REFERENCES "orders" (id);

ALTER TABLE "order_items"
    ADD CONSTRAINT fk_order_items_item FOREIGN KEY (item_id) 
        REFERENCES "items" (id);