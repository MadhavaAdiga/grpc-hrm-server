CREATE TABLE "payrolls" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "employee" uuid NOT NULL,
    "ctc" integer,
    "allowance" integer,
    "create_by" uuid NOT NULL,
    "updated_by" uuid NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "payrolls" ("employee");

