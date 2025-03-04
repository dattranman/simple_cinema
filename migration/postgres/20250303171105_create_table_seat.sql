-- +goose Up
-- +goose StatementBegin
create table seats (
    room_id integer,
    "row" integer,
    "column" integer,
    PRIMARY KEY (room_id, "row", "column"),
    FOREIGN KEY (room_id) REFERENCES rooms(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE seats;
-- +goose StatementEnd
