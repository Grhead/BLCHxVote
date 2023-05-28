-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Chains (
                                      Id CHAR(36) PRIMARY KEY,
                                      Hash VARCHAR(44) UNIQUE NOT NULL,
                                      Block TEXT NOT NULL

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
