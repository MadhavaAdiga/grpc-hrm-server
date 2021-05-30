-- check if uuid extention is available
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
SELECT uuid_generate_v4();

-- add an admin table

CREATE TABLE "organizations" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" varchar NOT NULL,
    "created_by" varchar NOT NULL,
    "creator_id" uuid NOT NULL,
    "status" integer NOT NUll,
    "updated_by" varchar,
    "updater_id" uuid,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "organizations" ("name");