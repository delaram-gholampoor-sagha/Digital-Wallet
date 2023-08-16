CREATE TABLE public.bank (
  bank_id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  bank_code NOT NULL VARCHAR(10), -- Unique code representing the bank
  status ENUM('active', 'inactive') DEFAULT 'active',
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
);