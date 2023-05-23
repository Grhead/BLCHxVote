-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Chains (
                                      Id CHAR(36) PRIMARY KEY,
                                      Hash VARCHAR(44) UNIQUE,
                                      Block TEXT

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
