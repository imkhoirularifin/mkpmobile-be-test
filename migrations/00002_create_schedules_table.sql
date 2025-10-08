-- +goose Up
-- +goose StatementBegin
CREATE TABLE schedules (
    id SERIAL PRIMARY KEY,
    movie_title VARCHAR(255) NOT NULL,
    studio_name VARCHAR(100) NOT NULL,
    show_date DATE NOT NULL,
    show_time VARCHAR(10) NOT NULL,
    available_seats INTEGER NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_schedules_deleted_at ON schedules(deleted_at);
CREATE INDEX idx_schedules_show_date ON schedules(show_date);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS schedules;
-- +goose StatementEnd
