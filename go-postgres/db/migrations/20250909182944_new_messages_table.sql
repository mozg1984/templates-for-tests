-- +goose Up
-- +goose StatementBegin
create table if not exists messages (
    id serial primary key,
    value text not null,
    author varchar(128) not null,
    created_at timestamp(0) not null default current_timestamp,
    updated_at timestamp(0) not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists messages;
-- +goose StatementEnd
