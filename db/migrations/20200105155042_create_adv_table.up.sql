CREATE TABLE IF NOT EXISTS adv.adv (
    id          SERIAL              NOT NULL,
    name        VARCHAR(1000)       NOT NULL,
    price       DECIMAL(10,0)    NOT NULL,
    photos      jsonb             NOT NULL,
    deleted_at  time,
    updated_at  time,
    created_at  time,
    PRIMARY KEY (id)
);
