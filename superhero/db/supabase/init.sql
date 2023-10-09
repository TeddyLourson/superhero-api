CREATE EXTENSION dblink WITH SCHEMA public;
/**
 * HANDLE_USER_UPSERT
 * Triggered every time a new user signs up to initialize his profile with the information he provided.
 */
CREATE FUNCTION public.handle_user_upsert()
RETURNS TRIGGER
LANGUAGE PLPGSQL
SECURITY DEFINER
SET search_path = public, extensions AS
$$
  DECLARE
    r UUID;
  BEGIN
    SELECT * FROM dblink('postgres://TeddyLourson:DXsBE6TN2WQv@ep-fragrant-hall-19811233-pooler.eu-central-1.aws.neon.tech/usr', format('SELECT handle_user_upsert(%L::UUID, %L, %L::UUID, %L);', NEW.id, NEW.email, NEW.raw_user_meta_data->'digibearFirstProfile'->>'id', NEW.raw_user_meta_data->'digibearFirstProfile'->>'username')) AS x(x UUID) INTO r;
    RETURN NEW;
  END;
$$;

/**
 * HANDLE_USER_DELETE
 * Triggered every time a new user signs up to initialize his profile with the information he provided.
 */
CREATE FUNCTION public.handle_user_delete()
RETURNS TRIGGER
LANGUAGE PLPGSQL
SECURITY DEFINER
SET search_path = public, extensions AS
$$
  DECLARE
    r UUID;
  BEGIN
    SELECT * FROM dblink('postgres://TeddyLourson:DXsBE6TN2WQv@ep-fragrant-hall-19811233-pooler.eu-central-1.aws.neon.tech/usr', format('SELECT handle_user_delete(%L::UUID);', OLD.id)) AS x(x UUID) INTO r;
    RETURN OLD;
  END;
$$;

CREATE TRIGGER on_auth_user_upsert
  AFTER INSERT OR UPDATE
  ON auth.users
  FOR EACH ROW EXECUTE PROCEDURE public.handle_user_upsert();

CREATE TRIGGER on_auth_user_delete
  AFTER DELETE
  ON auth.users
  FOR EACH ROW EXECUTE PROCEDURE public.handle_user_delete();

-- migrate:down
DROP TRIGGER on_auth_user_delete ON auth.users;
DROP FUNCTION handle_user_delete;
DROP TRIGGER on_auth_user_upsert ON auth.users;
DROP FUNCTION handle_user_upsert;
DROP EXTENSION dblink;