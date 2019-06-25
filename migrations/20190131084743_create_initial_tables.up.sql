CREATE SCHEMA app;
SET search_path TO app,public;

-- CREATE FUNCTIONS START
CREATE OR REPLACE FUNCTION app.trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- CREATE FUNCTIONS END

-- CREATE TABLES START
CREATE TABLE app.requests (
    id BIGSERIAL,
    platform_id BIGSERIAL,
    endpoint varchar NOT NULL,
    params varchar  NOT NULL,
    data json NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE app.platforms (
    id BIGSERIAL,
    name varchar NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- CREATE TABLES END

-- INDEXES AND CONSTRAINTS START
ALTER TABLE app."requests" ADD CONSTRAINT requests_endpoint_params UNIQUE (endpoint, params);
CREATE INDEX requests_platform_id ON app."requests" (platform_id);
-- INDEXES AND CONSTRAINTS END

-- CREATE TRIGGERS START
-- these triggers makes sure the created_at and updated_at columns are automatically set/updates on inserts/updates
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON app.requests
FOR EACH ROW
EXECUTE PROCEDURE app.trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON app.platforms
FOR EACH ROW
EXECUTE PROCEDURE app.trigger_set_timestamp();
-- CREATE TRIGGERS END

-- CREATE VIEWS START
CREATE VIEW app.stay_rooms_search_results_rooms_ids AS
  WITH rows AS (
  	  SELECT data
	  FROM app.requests
	  WHERE endpoint = 'searchStayRooms'
  ),
  results AS (
	  SELECT DISTINCT json_array_elements(data->'data'->'searchStayRooms'->'results')
	  AS room
	  FROM rows
  )
  SELECT DISTINCT (results.room ->> 'uid'::text) AS uid, (results.room ->> 'id'::text) as id
  from results;

CREATE VIEW app.rooms_search_results_rooms_ids AS
  WITH rows AS (
  	  SELECT data
	  FROM app.requests
	  WHERE endpoint = 'searchRooms'
  ),
  results AS (
	  SELECT DISTINCT json_array_elements(data->'data'->'searchRooms'->'results')
	  AS room
	  FROM rows
  )
  SELECT DISTINCT (results.room ->> 'uid'::text) AS uid, (results.room ->> 'id'::text) as id
  from results;
-- CREATE VIEWS END

-- SEED DATA
INSERT INTO app.platforms (id, name) VALUES (1, 'spacemarket'), (2, 'instabase'), (3, 'spacee');