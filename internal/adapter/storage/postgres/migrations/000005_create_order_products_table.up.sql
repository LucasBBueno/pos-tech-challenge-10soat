CREATE TABLE IF NOT EXISTS "order_products" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "order_id" uuid NOT NULL,
	"product_id" uuid NOT NULL,
    "quantity" integer NOT NULL,
    "sub_total" numeric(10, 2) NOT NULL,
    "observation" varchar NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL,
	CONSTRAINT order_products_pk PRIMARY KEY (id)
);

ALTER TABLE "order_products"
      ADD CONSTRAINT fk_order_products_order FOREIGN KEY (order_id) 
          REFERENCES "orders" (id);

ALTER TABLE "order_products"
    ADD CONSTRAINT fk_order_products_product FOREIGN KEY (product_id) 
        REFERENCES "products" (id);