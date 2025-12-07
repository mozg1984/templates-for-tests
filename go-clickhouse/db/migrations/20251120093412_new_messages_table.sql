-- +goose Up
-- +goose StatementBegin
CREATE TABLE ch.messages (
     user_id UInt64,
     text String,
     date DateTime
)
ENGINE = MergeTree
ORDER BY date;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ch.messages;
-- +goose StatementEnd
