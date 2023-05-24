CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currencry" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "acccount_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_acccount_id" bigint NOT NULL,
  "to_acccount_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "accounts" ("balance");

CREATE INDEX ON "accounts" ("currencry");

CREATE INDEX ON "entries" ("acccount_id");

CREATE INDEX ON "transfers" ("to_acccount_id");

CREATE INDEX ON "transfers" ("from_acccount_id");

CREATE INDEX ON "transfers" ("to_acccount_id", "from_acccount_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("acccount_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_acccount_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_acccount_id") REFERENCES "accounts" ("id");