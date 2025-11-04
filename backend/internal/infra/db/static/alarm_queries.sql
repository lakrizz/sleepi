-- name: CreateAlarm :exec
INSERT INTO alarms (
    id, label, time, enabled, warmup_duration, led_target, playable_id, weekdays
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: GetAlarm :one
SELECT * FROM alarms WHERE id = ?;

-- name: ListAlarms :many
SELECT * FROM alarms ORDER BY time;

-- name: UpdateAlarm :exec
UPDATE alarms
SET label = ?, time = ?, enabled = ?, warmup_duration = ?, led_target = ?, playable_id = ?, weekdays = ?
WHERE id = ?;

-- name: DeleteAlarm :exec
DELETE FROM alarms WHERE id = ?;

-- name: ToggleAlarm :exec
UPDATE alarms SET enabled = ? WHERE id = ?;

