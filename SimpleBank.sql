CREATE TABLE "accounts" (
  "id" serial PRIMARY KEY,
  "owner" text NOT NULL,
  "balance" numeric(20,2) NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" serial PRIMARY KEY,
  "from_account_id" integer NOT NULL,
  "to_account_id" integer NOT NULL,
  "amount" numeric(20,2) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" serial PRIMARY KEY,
  "account_id" integer NOT NULL,
  "amount" numeric(20,2) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "accounts" ("owner");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "entries" ("account_id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
