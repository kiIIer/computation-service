-- +goose Up
-- +goose StatementBegin
CREATE TABLE computation_tasks (
    id SERIAL PRIMARY KEY,
    status VARCHAR(50) NOT NULL,
    input TEXT,
    result TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
