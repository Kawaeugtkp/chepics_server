CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "timestamp" timestamptz NOT NULL DEFAULT (now()),
  "owner_id" bigint NOT NULL,
  "type" varchar NOT NULL,
  "is_root_opinion" boolean,
  "votes" integer NOT NULL DEFAULT 0,
  "topic" varchar NOT NULL,
  "description" varchar,
  "caption" varchar,
  "topic_id" bigint,
  "set_id" bigint,
  "category" varchar NOT NULL,
  "base_opinion_id" bigint,
  "post_image_url" varchar,
  "link" varchar
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "profile_image_url" varchar NOT NULL,
  "bio" varchar NOT NULL,
  "follower" integer NOT NULL DEFAULT 0,
  "following" integer NOT NULL DEFAULT 0
);

CREATE TABLE "picksets" (
  "id" bigserial PRIMARY KEY,
  "topic_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "votes" integer NOT NULL DEFAULT 0
);

ALTER TABLE "posts" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("set_id") REFERENCES "picksets" ("id");
