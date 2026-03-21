CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE counter (
    count INTEGER NOT NULL
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20260321001535');
