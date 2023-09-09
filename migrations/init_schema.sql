create table "users"(
    "id" bigserial PRIMARY KEY,
    "login" varchar NOT NULL UNIQUE,
    "email" varchar NOT NULL UNIQUE,
    "mobile_phone" varchar UNIQUE default (NULL),
    "password_hash" varchar NOT NULL,
    "created_at" timestamptz not null default (now())
);