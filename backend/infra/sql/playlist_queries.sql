  -- name: CreatePlaylist :exec
INSERT INTO playlists (id, name) VALUES (?, ?);

-- name: AddPlaylistFile :exec
INSERT INTO playlist_files (playlist_id, file_id, position)
VALUES (?, ?, ?);

-- name: ListPlaylists :many
SELECT * FROM playlists ORDER BY name;

-- name: GetPlaylist :one
SELECT * FROM playlists WHERE id = ?;

-- name: GetPlaylistFiles :many
SELECT f.* FROM media_files f
JOIN playlist_files pf ON f.id = pf.file_id
WHERE pf.playlist_id = ?
ORDER BY pf.position;

-- name: UpdatePlaylist :exec
UPDATE playlists SET name = ? WHERE id = ?;

-- name: DeletePlaylist :exec
DELETE FROM playlists WHERE id = ?;

-- name: ReorderPlaylistFiles :exec
UPDATE playlist_files SET position = ? WHERE playlist_id = ? AND file_id = ?;

