create type bucket_type_enum as enum ('aws', 'azure', 'oci');
create type bucket_region_enum as enum ('ashburn', 'frankfurt', 'london', 'phoenix', 'singapore', 'sydney');

create table bucket (
  id bigserial primary key,
  public_id uuid not null default (uuid_generate_v4()),
  name text not null,
  region bucket_region_enum not null,
  type bucket_type_enum not null,
  created_at timestamptz not null default (now()),
  updated_at timestamptz not null default (now())
);

create unique index bucket_public_id_uindex on bucket (public_id);
create unique index bucket_name_uindex on bucket (name);
