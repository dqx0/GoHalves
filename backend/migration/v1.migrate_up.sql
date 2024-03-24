create table "accounts" (
  "id" bigserial primary key,
  "user_id" varchar(255) not null,
  "name" varchar(255) not null,
  "email" varchar(255) default null,
  "password" varchar(255) not null,
  "is_bot" boolean default false,
  "created_at" timestamp not null default current_timestamp,
  "updated_at" timestamp not null default current_timestamp
);

create table "events" (
  "id" bigserial primary key,
  "title" varchar(255) not null,
  "description" text not null,
  "created_at" timestamp not null default current_timestamp,
  "updated_at" timestamp not null default current_timestamp
);

create table "pays"(
  "id" bigserial primary key,
  "paid_user_id" bigint not null,
  "event_id" bigint not null,
  "amount" bigint not null,
  "created_at" timestamp not null default current_timestamp,
  "updated_at" timestamp not null default current_timestamp
);

create table "accounts_events"(
  "id" bigserial primary key,
  "account_id" bigint not null,
  "event_id" bigint not null,
  "authority_id" bigint not null,
  "created_at" timestamp not null default current_timestamp,
  "updated_at" timestamp not null default current_timestamp
);

create table "authorities"(
  "id" bigserial primary key,
  "add_pays" boolean not null default true,
  "edit_pays" boolean not null default true,
  "delete_pays" boolean not null default true,
  "add_user" boolean not null default true,
  "edit_event" boolean not null default true,
  "delete_user" boolean not null default true
);

create table "friends"(
  "id" bigserial primary key,
  "send_account_id" bigint not null,
  "received_account_id" bigint not null,
  "send_at" timestamp not null default current_timestamp,
  "accepted_at" timestamp not null default current_timestamp
);

create table "accounts_pays"(
  "id" bigserial primary key,
  "account_id" bigint not null,
  "pay_id" bigint not null,
  "created_at" timestamp not null default current_timestamp,
  "updated_at" timestamp not null default current_timestamp
);


create index on "accounts" ("user_id");

create index on "accounts" ("name");

create index on "events" ("title");

create index on "pays" ("paid_user_id");

create index on "pays" ("event_id");

create index on "accounts_events" ("account_id", "event_id");

create index on "friends" ("send_account_id");

create index on "friends" ("received_account_id");

create index on "friends" ("send_account_id", "received_account_id");

create index on "accounts_pays" ("account_id");

create index on "accounts_pays" ("pay_id");

alter table "accounts_events" add foreign key ("account_id") references "accounts" ("id");
alter table "accounts_events" add foreign key ("event_id") references "events" ("id");
alter table "accounts_events" add foreign key ("authority_id") references "authorities" ("id");
alter table "pays" add foreign key ("paid_user_id") references "accounts" ("id");
alter table "pays" add foreign key ("event_id") references "events" ("id");
alter table "accounts_pays" add foreign key ("account_id") references "accounts" ("id");
alter table "accounts_pays" add foreign key ("pay_id") references "pays" ("id");
alter table "friends" add foreign key ("send_account_id") references "accounts" ("id");
alter table "friends" add foreign key ("received_account_id") references "accounts" ("id");

