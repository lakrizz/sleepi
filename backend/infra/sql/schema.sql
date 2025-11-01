-- alarms.sql
-- Table for main alarm objects
CREATE TABLE alarms (
    id TEXT PRIMARY KEY,                         -- UUID
    label TEXT NOT NULL,
    time TEXT NOT NULL,                          -- "HH:MM"
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    warmup_duration INTEGER DEFAULT 0,           -- seconds
    led_target TEXT,                             -- JSON or serialized RGB
    music_playlist_id TEXT,
    music_file_id TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Optional: trigger for automatic updated_at management
CREATE TRIGGER update_alarm_timestamp
AFTER UPDATE ON alarms
BEGIN
    UPDATE alarms SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;

-- Table for weekday repetitions (repeated Weekday field)
-- Assuming Weekday enum: MONDAY=1 .. SUNDAY=7
CREATE TABLE alarm_repeat_days (
    alarm_id TEXT NOT NULL,
    weekday INTEGER NOT NULL, -- 1..7
    PRIMARY KEY (alarm_id, weekday),
    FOREIGN KEY (alarm_id) REFERENCES alarms(id) ON DELETE CASCADE
);

-- Index for quick lookup
CREATE INDEX idx_alarm_repeat_days_alarm_id ON alarm_repeat_days(alarm_id);

-- Optional convenience view to join alarms and repeat days
CREATE VIEW alarm_with_days AS
SELECT a.*, group_concat(r.weekday, ',') AS repeat_days
FROM alarms a
LEFT JOIN alarm_repeat_days r ON a.id = r.alarm_id
GROUP BY a.id;

-- --- MEDIA FILES ---------------------------------------------------------

CREATE TABLE media_files (
    id TEXT PRIMARY KEY,                 -- UUID
    title TEXT NOT NULL,
    artist TEXT,
    size_bytes INTEGER NOT NULL,
    length_hms TEXT NOT NULL,            -- "MM:SS" or "HH:MM:SS"
    mime_type TEXT NOT NULL,
    type INTEGER NOT NULL,               -- FileType enum
    uploaded_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- --- PLAYLISTS -----------------------------------------------------------
CREATE TABLE playlists (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_playlist_timestamp
AFTER UPDATE ON playlists
BEGIN
    UPDATE playlists SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;

CREATE TABLE playlist_files (
    playlist_id TEXT NOT NULL,
    file_id TEXT NOT NULL,
    position INTEGER NOT NULL,
    PRIMARY KEY (playlist_id, file_id),
    FOREIGN KEY (playlist_id) REFERENCES playlists(id) ON DELETE CASCADE,
    FOREIGN KEY (file_id) REFERENCES media_files(id) ON DELETE CASCADE
);

CREATE INDEX idx_playlist_files_playlist_id ON playlist_files(playlist_id);

-- --- SLEEPSCAPES --------------------------------------------------------
CREATE TABLE sleepscapes (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    source_type INTEGER NOT NULL,       -- SourceType enum
    source_id TEXT NOT NULL,
    source_name TEXT,
    led_expression TEXT,
    author TEXT,
    description TEXT,
    tags TEXT,                          -- comma-joined or JSON
    verified BOOLEAN DEFAULT FALSE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_sleepscape_timestamp
AFTER UPDATE ON sleepscapes
BEGIN
    UPDATE sleepscapes SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;

-- --- SYSTEM SNAPSHOT CACHE ----------------------------------------------
CREATE TABLE system_info_cache (
    id INTEGER PRIMARY KEY CHECK (id = 1),
    os TEXT,
    arch TEXT,
    hostname TEXT,
    ip TEXT,
    network_iface TEXT,
    last_boot DATETIME,
    uptime_seconds INTEGER,
    current_version TEXT,
    latest_version TEXT,
    release_notes TEXT,     -- JSON or newline-joined
    disk_used TEXT,
    disk_total TEXT,
    disk_used_pct REAL,
    mem_used TEXT,
    mem_total TEXT,
    mem_used_pct REAL,
    cpu_load_pct REAL,
    temperature_c REAL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
