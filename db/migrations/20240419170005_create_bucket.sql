-- migrate:up
create type bucket_type_enum as enum ('aws', 'azure', 'oci');
create type bucket_region_enum as enum ('ap-southeast-1', 'ap-southeast-2', 'eu-central-1', 'eu-west-2', 'us-east-1', 'us-west-1');

create table bucket (
  id bigserial primary key,
  public_id uuid not null default (uuid_generate_v4()),
  name text not null,
  region bucket_region_enum not null,
  type bucket_type_enum not null,
  created_at timestamp not null default (now()),
  updated_at timestamp not null default (now())
);

create unique index bucket_public_id_uindex on bucket (public_id);
create unique index bucket_name_uindex on bucket (name);

-- migrate:down
drop index if exists bucket_public_id_uindex;
drop index if exists bucket_name_uindex;

drop table if exists bucket;

drop type if exists bucket_type_enum;
drop type if exists bucket_region_enum;
