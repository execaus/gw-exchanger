-- name: GetAllCurrencies :many
SELECT * FROM app.currencies;

-- name: GetCurrency :one
SELECT * FROM app.currencies WHERE code=$1;
