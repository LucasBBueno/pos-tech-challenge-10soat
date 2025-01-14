
CREATE TABLE IF NOT EXISTS "clients" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "cpf" varchar NULL,
	"name" varchar NULL,
    "email" varchar NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL,
	CONSTRAINT clients_pk PRIMARY KEY (id)
);