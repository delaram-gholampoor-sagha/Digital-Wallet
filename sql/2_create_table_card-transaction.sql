CREATE TABLE public.card_transaction (
    transaction_id SERIAL PRIMARY KEY,
    transaction_group_id INT NOT NULL, -- To group the pair of transactions (sender & receiver)
    financial_card_id INT NOT NULL REFERENCES public.financial_card,
    amount DECIMAL(15, 2) NOT NULL, -- Positive for credits, negative for debits (sender would be negative, receiver positive)
    balance DECIMAL(15, 2) NOT NULL, -- Balance of the card's linked account after the transaction
    description TEXT,
    status ENUM(
        'pending',
        'completed',
        'failed',
        'reversed',
        'on_hold',
        'cancelled'
    ),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
);