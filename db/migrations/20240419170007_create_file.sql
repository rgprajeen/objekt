-- migrate:up
create table file (
  id bigserial primary key,
  public_id uuid not null default (uuid_generate_v4()),
  name text not null,
  size bigint not null check(size >= 0),
  mime_type text not null,
  bucket_id bigint not null,
  created_at timestamp not null default (now()),
  updated_at timestamp not null default (now()),
  unique (name, bucket_id),
  constraint fk_file_bucket foreign key (bucket_id) references bucket (id)
);

create unique index file_public_id_uindex on file (public_id);
create index idx_file_mimetype on file (mime_type);
create index idx_file_size on file (size);

-- migrate:down
drop index if exists file_public_id_uindex;
drop index if exists idx_file_mimetype;
drop index if exists idx_file_size;

drop table if exists file;
