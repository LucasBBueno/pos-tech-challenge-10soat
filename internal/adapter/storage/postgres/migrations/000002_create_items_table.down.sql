ALTER TABLE "items"
    DROP CONSTRAINT items_pk

ALTER TABLE "items"
    DROP CONSTRAINT fk_items_category

DROP TABLE IF EXISTS "items"