CREATE TABLE "roles" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" VARCHAR NOT NULL,
    "active" boolean DEFAULT false,
    "organization" uuid NOT NULL,
    "permissions " integer[],
    "created_by" uuid NOT NULL,
    "updated_by" uuid NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "roles" ("name");
CREATE INDEX ON "roles" ("organization");
CREATE INDEX ON "roles" ("name","organization");

ALTER TABLE "roles" ADD FOREIGN KEY ("organization") REFERENCES "organizations" ("id");