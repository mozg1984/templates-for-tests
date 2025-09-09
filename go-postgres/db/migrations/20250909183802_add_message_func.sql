-- +goose Up
-- +goose StatementBegin
create or replace function add_message(p_value text, p_author varchar(128)) returns integer as $$
    declare
        l_id bigint;
    begin
        insert into messages (value, author)
        values (p_value, p_author)
        returning id into l_id;

        return l_id;
    end
$$ language plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop function if exists add_message;
-- +goose StatementEnd
