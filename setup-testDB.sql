CREATE TABLE "public"."users" ("id" uuid NOT NULL, "name" text NOT NULL, PRIMARY KEY ("id") );
CREATE TABLE "public"."tasks" ("id" uuid NOT NULL, "title" text NOT NULL, "explanation" text, "limit" timestamptz NOT NULL, "priority" integer NOT NULL, "status" text NOT NULL, "user_id" uuid NOT NULL, PRIMARY KEY ("id") );
CREATE TABLE "public"."task_label_relations" ("id" UUID NOT NULL, "task_id" uuid NOT NULL, "label_id" uuid NOT NULL, PRIMARY KEY ("id") );
CREATE TABLE "public"."labels" ("id" uuid NOT NULL, "value" integer NOT NULL, PRIMARY KEY ("id") );
alter table "public"."task_label_relations"
  add constraint "task_label_relations_task_id_fkey"
  foreign key ("task_id")
  references "public"."tasks"
  ("id") on update restrict on delete restrict;
  alter table "public"."tasks"
  add constraint "tasks_user_id_fkey"
  foreign key ("user_id")
  references "public"."users"
  ("id") on update restrict on delete restrict;
--  insert
INSERT INTO labels (id, value) VALUES
('ab997dba-8d92-5f05-b914-9c71d78f8ec9', 0),
('26404d20-c2cb-4016-d163-72fde3e52370', 1),
('67488461-a378-7793-d01c-6396d1bbef75', 2),
('e39d826d-16fd-692f-4755-2a704d581412', 3),
('e3fb0138-f67e-ed48-a198-1c6d9c247c16', 4),
('fe9a90cf-da49-dd43-4fa5-cfd44677846a', 5);
insert into users values ('3bdb5a00-7ac5-01e4-2b9a-64f787b698db', 'taku'); 