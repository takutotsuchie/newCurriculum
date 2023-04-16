alter table "public"."task_label_relations"
  add constraint "task_label_relations_task_id_fkey"
  foreign key ("task_id")
  references "public"."tasks"
  ("id") on update restrict on delete restrict;
