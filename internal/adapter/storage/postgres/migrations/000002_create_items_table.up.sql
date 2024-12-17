
CREATE TABLE IF NOT EXISTS "items" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"name" varchar NOT NULL,
    "category_id" uuid NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL,
	CONSTRAINT items_pk PRIMARY KEY (id)
);

ALTER TABLE "items"
      ADD CONSTRAINT fk_items_category FOREIGN KEY (category_id) 
          REFERENCES "categories" (id);