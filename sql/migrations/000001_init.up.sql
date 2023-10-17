CREATE TABLE "categories" (
  "id" varchar NOT NULL PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar,
  "is_active" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);