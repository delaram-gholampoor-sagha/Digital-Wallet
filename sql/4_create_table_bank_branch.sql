CREATE TABLE public.bank_branche (
  branch_id SERIAL PRIMARY KEY,
  bank_id INTEGER NOT NULL REFERENCES public.bank,
  branch_name VARCHAR(100) NOT NULL,
  branch_code VARCHAR(10) unique,
  address VARCHAR(255),
  city VARCHAR(50),
  province VARCHAR(50),
  postal_code VARCHAR(10),
  phone_number VARCHAR(20),
  status ENUM('active', 'inactive') DEFAULT 'active',
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
);