INSERT INTO "products" ("name", "category_id") VALUES
  ('Lanche 1', (SELECT id FROM "categories" WHERE name = 'Lanche')),
  ('Lanche 2', (SELECT id FROM "categories" WHERE name = 'Lanche')),
  ('Lanche 3', (SELECT id FROM "categories" WHERE name = 'Lanche')),
  ('Acompanhamento 1', (SELECT id FROM "categories" WHERE name = 'Acompanhamento')),
  ('Acompanhamento 2', (SELECT id FROM "categories" WHERE name = 'Acompanhamento')),
  ('Acompanhamento 3', (SELECT id FROM "categories" WHERE name = 'Acompanhamento')),
  ('Bebida 1', (SELECT id FROM "categories" WHERE name = 'Bebida')),
  ('Bebida 2', (SELECT id FROM "categories" WHERE name = 'Bebida')),
  ('Bebida 3', (SELECT id FROM "categories" WHERE name = 'Bebida')),
  ('Sobremesa 1', (SELECT id FROM "categories" WHERE name = 'Sobremesa')),
  ('Sobremesa 2', (SELECT id FROM "categories" WHERE name = 'Sobremesa')),
  ('Sobremesa 3', (SELECT id FROM "categories" WHERE name = 'Sobremesa'));
