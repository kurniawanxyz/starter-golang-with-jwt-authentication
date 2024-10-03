ALTER TABLE token_verifications
ADD COLUMN IF NOT EXISTS type ENUM('email_verification', 'forgot_password') NOT NULL DEFAULT 'email_verification';