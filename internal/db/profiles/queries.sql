-- name: AddProfile :exec
INSERT INTO profiles (name, secret, challenges) VALUES ($1, $2, $3);

-- name: GetProfiles :many
SELECT * FROM profiles ORDER BY id asc;

-- name: UpdateProfile :exec
UPDATE profiles SET secret = $1, challenges = $2 WHERE name = $3;

-- name: DeleteProfile :exec
DELETE FROM profiles WHERE name = $1;

-- name: CheckProfileExists :one
SELECT EXISTS(SELECT 1 FROM profiles WHERE name = $1);