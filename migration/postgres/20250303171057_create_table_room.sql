-- +goose Up
-- +goose StatementBegin
CREATE TABLE rooms (
    id serial PRIMARY KEY,
    "row" integer,
    "column" integer,
    "min_distance" integer
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rooms;
-- +goose StatementEnd
