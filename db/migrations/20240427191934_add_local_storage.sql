-- migrate:up
alter type bucket_type_enum add value if not exists 'local' before 'oci';
alter type bucket_region_enum add value if not exists 'local' before 'us-east-1';

-- migrate:down
-- cannot revert this upgrade and its safe to leave as it is
