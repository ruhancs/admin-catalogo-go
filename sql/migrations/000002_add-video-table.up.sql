CREATE TABLE "videos" (
  "id" varchar NOT NULL PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar,
  "duration" bigint,
  "is_published" boolean NOT NULL DEFAULT false,
  "banner" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);