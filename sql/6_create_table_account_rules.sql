CREATE TABLE public.account_rules (
    rule_id SERIAL PRIMARY KEY,
    financial_account_id INT NOT NULL REFERENCES public.financial_account,
    
    -- Transaction Limits
    daily_limit INT DEFAULT 10,
    weekly_limit INT DEFAULT 50,
    monthly_limit INT DEFAULT 200,

    -- for new accounts ??
    min_amount DECIMAL(15,2) DEFAULT 1.00,   -- Defaulting to a minimum of $1
    max_amount DECIMAL(15,2) DEFAULT 1000.00, -- Defaulting to a maximum of $1000

    -- Velocity Checks
    max_failed_transactions_daily INT DEFAULT 5,

    -- Account Age Restrictions 
    new_account_max_amount DECIMAL(15,2) DEFAULT 500.00, 

    -- Account Balance Restrictions
    min_balance_required DECIMAL(15,2) DEFAULT 0.00,


    -- Location-Based Checks
    allowed_countries TEXT[],  -- List of ISO country codes (e.g., ['US', 'CA', 'GB'])
    disallowed_countries TEXT[],
    notify_on_new_location BOOLEAN DEFAULT TRUE,
    block_on_unusual_location BOOLEAN DEFAULT TRUE,

    -- For more granular location control
    allowed_regions TEXT[],    -- List of specific regions/states
    allowed_cities TEXT[],     -- List of specific cities


    -- Two-Factor Authentication
    require_2fa_for_amount DECIMAL(15,2) DEFAULT 500.00,  -- Require 2FA for amounts greater than this

    -- Tiered Verification
    verification_level ENUM('minimal', 'intermediate', 'full') DEFAULT 'minimal',
    
    -- Cooling Periods (measured in hours)
    password_change_cooling_period INT DEFAULT 24,
    new_transaction_method_cooling_period INT DEFAULT 48,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
);
