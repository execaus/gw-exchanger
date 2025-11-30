-- name: GetAllCurrencies :many
SELECT * FROM app.currencies;

-- name: GetTwoCurrencies :many
SELECT * FROM app.currencies WHERE code IN ($1, $2);
