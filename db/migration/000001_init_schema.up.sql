-- check if uuid extention is available
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
SELECT uuid_generate_v4();

-- add an admin table

CREATE TABLE "users" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "first_name" varchar NOT NUll,
    "last_name" varchar,
    "user_name" varchar NOT NULL,
    "hashed_password" varchar NOT NULL,
    "address" varchar,
    "email" varchar,
    "contact_number" bigint UNIQUE NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "organizations" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" varchar NOT NULL,
    "creator_id" uuid NOT NULL,
    "status" smallint NOT NUll,
    "updater_id" uuid NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "organizations" ("name");
CREATE UNIQUE INDEX ON "users" ("user_name");