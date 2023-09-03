CREATE TABLE public.financial_cards (
    card_id SERIAL PRIMARY KEY,
    account_id INTEGER NOT NULL REFERENCES public.financial_account,
    card_number VARCHAR(19) NOT NULL, -- Usually 16-19 digits for bank cards
    card_type ENUM('debit', 'credit', 'gift'),
    expiration_date DATE NOT NULL,
    card_holder_name VARCHAR(100),
    cvv VARCHAR(4), -- Card Verification Value ??  ensuring security ??
    status ENUM('active', 'inactive', 'lost', 'stolen') DEFAULT 'active',
    issued_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
);