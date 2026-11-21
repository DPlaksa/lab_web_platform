CREATE TABLE IF NOT EXISTS button_stats (
    id SMALLINT PRIMARY KEY CHECK (id = 1),
    click_count BIGINT NOT NULL DEFAULT 0
);

INSERT INTO button_stats (id, click_count)
VALUES (1, 0)
ON CONFLICT (id) DO NOTHING;