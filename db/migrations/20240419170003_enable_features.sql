-- migrate:up
create extension if not exists "uuid-ossp";

-- migrate:down
drop extension if exists "uuid-ossp";
