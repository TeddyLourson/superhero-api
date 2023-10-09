/**********  ACCOUNT **********/
/**********  ACCOUNT **********/
/**********  ACCOUNT **********/

-- name: CreateAccount :one
INSERT INTO account (id, email)
VALUES (@id, @email)
RETURNING id;

-- name: GetAccountById :one
SELECT id,
  email,
  acquired_watch_time,
  created_at,
  updated_at
FROM account
WHERE id = @id
LIMIT 1;

-- name: GetAccountIDByEmail :one
SELECT id
FROM account
WHERE email = @email
LIMIT 1;

-- name: GetAccount :one
SELECT id,
  acquired_watch_time,
  email,
  created_at,
  updated_at
FROM account
WHERE id = @id
ORDER BY id;

/**********  PROFILE **********/
/**********  PROFILE **********/
/**********  PROFILE **********/

-- name: GetProfile :one
SELECT id,
  username,
  image_path,
  account_id,
  code,
  age_restriction,
  created_at,
  updated_at
FROM profile
WHERE id = @id
LIMIT 1;

-- name: GetAccountForProfileSelection :many
SELECT acct.id,
  acct.email,
  acct.acquired_watch_time,
  acct.created_at,
  acct.updated_at,
  prfl.id AS profile_id,
  prfl.username AS profile_username,
  prfl.code AS profile_code,
  prfl.image_path AS profile_image_path,
  prfl.age_restriction AS profile_age_restriction,
  prfl.account_id AS profile_account_id,
  prfl.created_at AS profile_created_at,
  prfl.updated_at AS profile_updated_at
FROM account acct
  JOIN profile prfl ON prfl.account_id = acct.id
WHERE acct.id = @id
ORDER BY acct.id;

-- name: GetProfilesForAccount :many
SELECT id,
  username,
  image_path,
  account_id,
  code,
  age_restriction,
  synced_theme_id,
  created_at,
  updated_at
FROM profile
WHERE account_id = @id
ORDER BY id;

-- name: CreateProfile :one
INSERT INTO profile (id, username, code, image_path, age_restriction, synced_theme_id, account_id)
VALUES (@id, @username, @code, @image_path, @age_restriction, @synced_theme_id, @account_id)
RETURNING id;

-- name: UpdateProfile :one
UPDATE profile
SET username = coalesce(@username, username),
  code = coalesce(@code, code),
  image_path = coalesce(@image_path, image_path),
  age_restriction = coalesce(@age_restriction, age_restriction),
  synced_theme_id = coalesce(@synced_theme_id, synced_theme_id)
WHERE id = @id
AND coalesce(@username, @code, @image_path, @age_restriction, @synced_theme_id) IS NOT NULL
RETURNING id;

-- name: DeleteProfile :one
DELETE FROM profile
WHERE id = @id
RETURNING id;

-- name: GetRolesForProfile :many
SELECT id,
  name,
  profile_id,
  created_at,
  updated_at
FROM profile_role
WHERE profile_id = @profile_id;