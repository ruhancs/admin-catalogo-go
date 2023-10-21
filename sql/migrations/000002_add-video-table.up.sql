CREATE TABLE "videos" (
  "id" varchar NOT NULL PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar,
  "duration" bigint,
  "year_launched" bigint NOT NULL,
  "is_published" boolean NOT NULL DEFAULT false,
  "banner_url" varchar,
  "video_url" varchar,
  "categories_id" varchar[],
  "created_at" timestamptz NOT NULL DEFAULT (now())
);