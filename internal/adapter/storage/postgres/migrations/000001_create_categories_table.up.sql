CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "categories" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"name" varchar NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL,
	CONSTRAINT categories_pk PRIMARY KEY (id)
);