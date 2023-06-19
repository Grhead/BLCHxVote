-- +goose Up
-- +goose StatementBegin
ALTER TABLE TransactionQueue ADD Master VARCHAR(64);
-- +goose StatementEnd

-- +goose Down
