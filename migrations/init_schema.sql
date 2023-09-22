create table "users"(
    "id" bigserial PRIMARY KEY,
    "login" varchar NOT NULL UNIQUE,
    "email" varchar NOT NULL UNIQUE,
    "mobile_phone" varchar UNIQUE default (NULL),
    "password_hash" varchar NOT NULL,
    "created_at" timestamptz not null default (now())
);

create table "wardrobes_users"(
    "user_id" bigint not null references users (id),
    "wardrobe_id" bigint not null references wardrobes (id),
    "is_owner" boolean not null default (FALSE)
);

create table "wardrobes"(
    "id" bigserial PRIMARY KEY,
    "title" varchar not null,
    "description" varchar,
    "poster" bytea,
    "created_at" timestamptz not null default (now())
);

create table "clothes"(
    "id" bigserial PRIMARY KEY,
    "name" varchar not null,
    "type" varchar,
    "numerical_size" int,
    "unified_size" varchar,
    "wardrobe_id" bigint not null references wardrobes (id),
    "poster" bytea,
    "image" bytea,
    "created_at" timestamptz not null default (now())
)