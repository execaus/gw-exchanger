-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS app.currencies (
    code VARCHAR(10) PRIMARY KEY,
    rate REAL NOT NULL
);

COMMENT ON TABLE app.currencies IS 'Справочник валют';
COMMENT ON COLUMN app.currencies.code IS 'Уникальный код валюты (ISO)';
COMMENT ON COLUMN app.currencies.rate IS 'Курс валюты относительно базовой';


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS app.currencies;

-- +goose StatementEnd
