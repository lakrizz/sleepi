-- name: CreateSleepscape :exec
INSERT INTO sleepscapes
(id, name, source_type, source_id, source_name, led_expression, author, description, tags, verified)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: ListSleepscapes :many
SELECT * FROM sleepscapes ORDER BY name LIMIT ? OFFSET ?;

-- name: GetSleepscape :one
SELECT * FROM sleepscapes WHERE id = ?;

-- name: UpdateSleepscape :exec
UPDATE sleepscapes
SET name = ?, source_type = ?, source_id = ?, source_name = ?, led_expression = ?,
    author = ?, description = ?, tags = ?, verified = ?
WHERE id = ?;

-- name: DeleteSleepscape :exec
DELETE FROM sleepscapes WHERE id = ?;

