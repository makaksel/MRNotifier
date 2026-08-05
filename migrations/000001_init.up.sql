CREATE TABLE merge_requests
(
    id              SERIAL PRIMARY KEY,

    project_path    TEXT    NOT NULL,
    mr_iid          INTEGER NOT NULL,

    title           TEXT,
    description     TEXT,

    state           TEXT,
    web_url         TEXT,

    author_username TEXT,

    source_branch   TEXT,
    target_branch   TEXT,

    draft           BOOLEAN DEFAULT FALSE,

    changes_count   INTEGER DEFAULT 0,

    created_at      TIMESTAMPTZ,
    updated_at      TIMESTAMPTZ,
    merged_at       TIMESTAMPTZ,

    UNIQUE (project_path, mr_iid)
);
CREATE TABLE notifications
(
    id                  SERIAL PRIMARY KEY,

    project_path        TEXT    NOT NULL,

    mr_iid              INTEGER NOT NULL,

    reply_to_message_id INTEGER,

    status              TEXT    NOT NULL,

    created_at          TIMESTAMPTZ DEFAULT now(),

    UNIQUE (project_path, mr_iid, status)
);
