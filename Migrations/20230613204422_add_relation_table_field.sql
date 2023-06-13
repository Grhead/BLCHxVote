-- +goose Up
-- +goose StatementBegin
ALTER TABLE RelationPatterns ADD Master VARCHAR(64);
-- +goose StatementEnd

-- +goose Down

