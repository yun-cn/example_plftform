-- ================================================================================
-- users
-- --------------------------------------------------------------------------------

create table users (
  id uuid NOT NULL,
  email         character varying(255) not null default '',
  password_hash text                   not null,
  name          character varying(255) not null default '',
  avatar_url    character varying(255) not null default '',
  deleted_at    timestamp  without time zone,
  created_at     timestamp without time zone NOT NULL,
  updated_at     timestamp without time zone NOT NULL
);
ALTER TABLE users ADD CONSTRAINT users_pkey PRIMARY KEY (id);
create unique index users_i01 on users (email);

-- ================================================================================
-- json_web_tokens
-- --------------------------------------------------------------------------------

create table json_web_tokens (
  id uuid NOT NULL,
  resource_id    int not null default 0,
  resource_type  character varying(255) not null,
  token          text                   not null,
  expires_at     timestamp without time zone,
  created_at     timestamp without time zone NOT NULL,
  updated_at     timestamp without time zone NOT NULL
);
ALTER TABLE json_web_tokens_i01 ADD CONSTRAINT json_web_tokens_i01_pkey PRIMARY KEY (id);
create index json_web_tokens_i01 on json_web_tokens (resource_id, resource_type, expires_at);
