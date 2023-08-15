CREATE TABLE public.currency (
    currency_id SERIAL PRIMARY KEY,
    currency_code CHAR(3) NOT NULL UNIQUE,  -- ISO 4217 currency codes like USD, EUR, GBP
    currency_name TEXT NOT NULL,
    symbol CHAR(1),   -- $, €, £, etc.
    exchange_rate DECIMAL(15, 6) -- Relative to a base currency, e.g., USD. This field is optional and depends on whether you intend to handle conversion rates within this table.
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
);