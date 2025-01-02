
CREATE TABLE IF NOT EXISTS "products" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"name" varchar NOT NULL,
    "category_id" uuid NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL,
	CONSTRAINT products_pk PRIMARY KEY (id)
);

ALTER TABLE "products"
      ADD CONSTRAINT fk_products_category FOREIGN KEY (category_id) 
          REFERENCES "categories" (id);