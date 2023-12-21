CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" integer PRIMARY KEY,
  "balance" float DEFAULT 0,
  "currency" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "user_id" integer
);

CREATE TABLE "transfers" (
  "id" integer PRIMARY KEY,
  "amount" float NOT NULL,
  "reason" varchar,
  "account_id" integer
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
