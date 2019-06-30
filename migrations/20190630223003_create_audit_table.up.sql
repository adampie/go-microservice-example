CREATE TABLE audit
(
    id	INTEGER PRIMARY KEY,
    user_id	TEXT NOT NULL,
    org_id	TEXT,
    ip_address TEXT NOT NULL,
    target TEXT NOT NULL,
    action TEXT NOT NULL,
    action_type TEXT NOT NULL,
    event_name TEXT NOT NULL
);