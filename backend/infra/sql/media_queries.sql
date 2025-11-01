-- name: CreateMediaFile :exec
INSERT INTO media_files (id, title, artist, size_bytes, length_hms, mime_type, type)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: ListMediaFiles :many
SELECT * FROM media_files ORDER BY uploaded_at DESC LIMIT ? OFFSET ?;

-- name: SearchMediaFiles :many
SELECT * FROM media_files
WHERE title LIKE '%' || ? || '%' OR artist LIKE '%' || ? || '%'
ORDER BY uploaded_at DESC LIMIT ? OFFSET ?;

-- name: UpdateMediaFileMeta :exec
UPDATE media_files
SET title = COALESCE(?, title),
    artist = COALESCE(?, artist)
WHERE id = ?;

-- name: DeleteMediaFile :exec
DELETE FROM media_files WHERE id = ?;

