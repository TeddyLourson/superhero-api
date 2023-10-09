-- migrate:up
-- :::::::::: TABLES :::::::::: --
-- :::::::::: TABLES :::::::::: --
-- :::::::::: TABLES :::::::::: --

-- ::::: ACCOUNT ::::: --
CREATE TABLE account(
  id UUID PRIMARY KEY UNIQUE NOT NULL,
  email TEXT UNIQUE NOT NULL,
  acquired_watch_time INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ
);

-- ::::: PROFILE ::::: --
CREATE TABLE profile(
  id UUID PRIMARY KEY UNIQUE NOT NULL,
  username TEXT UNIQUE NOT NULL,
  code TEXT,
  image_path TEXT,
  age_restriction INTEGER,
  synced_theme_id TEXT,

  account_id UUID REFERENCES account(id) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ
);

-- ::::: PROFILE_ROLE ::::: --
CREATE TABLE profile_role (
  id UUID NOT NULL,
  name TEXT NOT NULL,
  profile_id UUID NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ
);

-- ::::: WATCH_RECORD ::::: --
CREATE TABLE watch_record (
  id UUID PRIMARY KEY UNIQUE NOT NULL,
  -- The "account_id" is necessary because a profile can be deleted, however the watch record still exists for the account.
  account_id UUID NOT NULL,
  profile_id UUID,
  media_id UUID NOT NULL,
  duration INTEGER NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ
);

-- ::::: PROHIBITED_MEDIA ::::: --
CREATE TABLE prohibited_media (
  id UUID PRIMARY KEY UNIQUE NOT NULL,
  profile_id UUID NOT NULL,
  media_id UUID NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ
);

-- :::::::::: FUNCTIONS :::::::::: --
-- :::::::::: FUNCTIONS :::::::::: --
-- :::::::::: FUNCTIONS :::::::::: --

-- ::::: HANDLE_USER_UPSERT ::::: --
/*
  Whenever a user is created/updated, this function should be called.
  It can be called from another database (supabase for example, through a TRIGGER).
*/
CREATE FUNCTION handle_user_upsert(account_id UUID, digibear_email TEXT, first_profile_id UUID, first_profile_username TEXT)
RETURNS UUID
LANGUAGE PLPGSQL
SECURITY DEFINER
SET search_path = public AS
$$
  BEGIN
    INSERT INTO account (id, email)
    VALUES(
      account_id,
      digibear_email
    ) ON CONFLICT (id)
    DO UPDATE
      SET email = digibear_email;
    INSERT INTO profile (id, username, account_id)
    VALUES(
      first_profile_id,
      first_profile_username,
      account_id
    ) ON CONFLICT (id)
      DO NOTHING;
    RETURN account_id;
  END;
$$;

-- ::::: HANDLE_USER_DELETE ::::: --
/*
  Whenever a user is deleted, this function should be called.
  It can be called from another database (supabase for example, through a TRIGGER).
*/
CREATE FUNCTION handle_user_delete(account_id UUID)
RETURNS UUID
LANGUAGE PLPGSQL
SECURITY DEFINER
SET search_path = public AS
$$
  BEGIN
    DELETE FROM profile
    WHERE account_id = account_id;
    DELETE FROM account
    WHERE account.id = account_id;
    RETURN account_id;
  END;
$$;

-- :::::::::: VIEWS :::::::::: --
-- :::::::::: VIEWS :::::::::: --
-- :::::::::: VIEWS :::::::::: --
-- CREATE VIEW profile_watch_time_view AS
--   SELECT
--     prof.id AS profile_id,
--     SUM(watch_rec.duration) AS profile_watch_time
--     FROM watch_record watch_rec
--     JOIN profile prof ON prof.id = watch_rec.profile_id
--     GROUP BY prof.id;

-- CREATE VIEW account_watch_time_view AS
--   SELECT
--     acct.id AS account_id,
--     SUM(watch_rec.duration) AS account_watch_time
--     FROM watch_record watch_rec
--     JOIN account acct ON acct.id = watch_rec.account_id
--     GROUP BY acct.id;

-- CREATE VIEW account_remaining_watch_time_view AS
--   SELECT
--     acct.id AS account_id,
--     (acct.acquired_watch_time - account_watch_time_view.account_watch_time) AS account_remaining_watch_time
--     FROM account_watch_time_view
--     JOIN account acct ON acct.id = account_watch_time_view.account_id
--     GROUP BY acct.id, account_watch_time_view.account_watch_time;

-- migrate:down
DROP FUNCTION handle_user_delete;
DROP FUNCTION handle_user_upsert;

DROP TABLE prohibited_media;
DROP TABLE watch_record;
DROP TABLE profile_role;
DROP TABLE profile;
DROP TABLE account;

-- :::::::::: UTILS :::::::::: --
-- :::::::::: UTILS :::::::::: --
-- :::::::::: UTILS :::::::::: --
-- ::::: TO DROP FUNCTIONS THAT YOU DON'T KNOW THE ARGS/RETURN TYPE OF:
SELECT 'DROP FUNCTION ' || oid::regprocedure
FROM   pg_proc
WHERE  proname = 'my_function_name'
AND    pg_function_is_visible(oid);