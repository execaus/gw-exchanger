package pkg

// Currency код валюты, например "USD", "EUR", "JPY"
type Currency = string

// Rate курс валюты относительно базовой валюты
type Rate = float32

// ExchangeRatesMap — карта валютных курсов
type ExchangeRatesMap = map[Currency]Rate
