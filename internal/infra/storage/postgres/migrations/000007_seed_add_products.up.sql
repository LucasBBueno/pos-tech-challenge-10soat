INSERT INTO "products" ("name", "description", "image", "value", "category_id") VALUES
  ('Lanche 1', 'Lanche com bacon', 'https://res.cloudinary.com/lucasbbueno/image/upload/v1735834118/Pos-tech%20images/lanche_vsnicd.webp', 15.90, (SELECT id FROM "categories" WHERE name = 'Lanche')),
  ('Lanche 2', 'Lanche com salada', 'https://res.cloudinary.com/lucasbbueno/image/upload/v1735834118/Pos-tech%20images/lanche_vsnicd.webp', 19.90, (SELECT id FROM "categories" WHERE name = 'Lanche')),
  ('Lanche 3', 'Lanche com queijo', 'https://res.cloudinary.com/lucasbbueno/image/upload/v1735834118/Pos-tech%20images/lanche_vsnicd.webp', 18.90, (SELECT id FROM "categories" WHERE name = 'Lanche')),
  ('Acompanhamento 1', 'Acompanhamento com queijo', 'https://res.cloudinary.com/lucasbbueno/image/upload/v1735834206/Pos-tech%20images/batata_yzkzfg.jpg', 9.90, (SELECT id FROM "categories" WHERE name = 'Acompanhamento')),
  ('Acompanhamento 2', 'Acompanhamento com cheddar', 'https://res.cloudinary.com/lucasbbueno/image/upload/v1735834206/Pos-tech%20images/batata_yzkzfg.jpg', 9.90, (SELECT id FROM "categories" WHERE name = 'Acompanhamento')),
  ('Acompanhamento 3', 'Acompanhamento com farofa bacon', 'https://res.cloudinary.com/lucasbbueno/image/upload/v1735834206/Pos-tech%20images/batata_yzkzfg.jpg', 9.90, (SELECT id FROM "categories" WHERE name = 'Acompanhamento')),
  ('Bebida 1', 'Suco laranja', 'https://res.cloudinary.com/lucasbbueno/image/upload/v1735834206/Pos-tech%20images/bebida_gqhxdh.webp', 12.90, (SELECT id FROM "categories" WHERE name = 'Bebida')),
  ('Bebida 2', 'Suco morango', 'https://res.cloudinary.com/lucasbbueno/image/upload/v1735834206/Pos-tech%20images/bebida_gqhxdh.webp', 12.90, (SELECT id FROM "categories" WHERE name = 'Bebida')),
  ('Bebida 3', 'Suco detox', 'https://res.cloudinary.com/lucasbbueno/image/upload/v1735834206/Pos-tech%20images/bebida_gqhxdh.webp', 12.90, (SELECT id FROM "categories" WHERE name = 'Bebida')),
  ('Sobremesa 1', 'Com chocolate', 'https://res.cloudinary.com/lucasbbueno/image/upload/v1735834206/Pos-tech%20images/sobremesa_qnejbz.jpg', 18.90, (SELECT id FROM "categories" WHERE name = 'Sobremesa')),
  ('Sobremesa 2', 'Com sorvete', 'https://res.cloudinary.com/lucasbbueno/image/upload/v1735834206/Pos-tech%20images/sobremesa_qnejbz.jpg', 18.90,(SELECT id FROM "categories" WHERE name = 'Sobremesa')),
  ('Sobremesa 3', 'Com frutas', 'https://res.cloudinary.com/lucasbbueno/image/upload/v1735834206/Pos-tech%20images/sobremesa_qnejbz.jpg', 18.90, (SELECT id FROM "categories" WHERE name = 'Sobremesa'));
