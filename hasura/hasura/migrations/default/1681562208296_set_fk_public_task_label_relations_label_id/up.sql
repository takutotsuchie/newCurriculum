alter table "public"."task_label_relations"
  add constraint "task_label_relations_label_id_fkey"
  foreign key ("label_id")
  references "public"."labels"
  ("id") on update restrict on delete restrict;
