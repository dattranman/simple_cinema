-- +goose Up
-- +goose StatementBegin
INSERT INTO rooms ("id", "row", "column", "min_distance") VALUES (1, 4, 5, 6);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM rooms WHERE id = 1;
-- +goose StatementEnd
