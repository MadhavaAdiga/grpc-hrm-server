CREATE TABLE "employees" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "user" uuid NOT NULL,
    "organization" uuid NOT NULL,
    "role" uuid NOT NULL, 
    "status" smallint NOT NULL,
    "create_by" uuid NOT NULL,
    "updated_by" uuid NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "employees" ("user");
CREATE INDEX ON "employees" ("role","organization");

ALTER TABLE "employees" ADD FOREIGN KEY ("user") REFERENCES "users" ("id");
ALTER TABLE "employees" ADD FOREIGN KEY ("organization") REFERENCES "organizations" ("id");
ALTER TABLE "employees" ADD FOREIGN KEY ("role") REFERENCES "roles" ("id");

ALTER TABLE "payrolls" ADD FOREIGN KEY ("employee") REFERENCES "employees" ("id");