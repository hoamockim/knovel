CREATE TABLE task (
    id          SERIAL NOT NULL,
    name        text NOT NULL,
    userid      text NOT NULL,
    description text NULL,
    status      text NOT NULL,
    created_at  timestamptz DEFAULT NOW(),
    updated_at  timestamptz DEFAULT NOW(),
    deleted_at  timestamptz NULL,
    PRIMARY KEY(id)
)