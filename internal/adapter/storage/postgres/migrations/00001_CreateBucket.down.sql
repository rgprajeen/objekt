drop index if exists bucket_public_id_uindex;
drop index if exists bucket_name_uindex;

drop table if exists bucket;

drop type if exists bucket_type_enum;
drop type if exists bucket_region_enum;
