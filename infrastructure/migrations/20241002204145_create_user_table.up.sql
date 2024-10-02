
CREATE TABLE users (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),  -- Menghasilkan UUID secara otomatis
    name VARCHAR(50) NOT NULL,                 -- Nama user, wajib diisi
    email VARCHAR(100) NOT NULL UNIQUE,        -- Email unik, wajib diisi
    telp VARCHAR(15) NOT NULL,                 -- Nomor telepon, wajib diisi
    password VARCHAR(255) NOT NULL,            -- Password user, wajib diisi
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Tanggal pembuatan record
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Diperbarui secara otomatis
    deleted_at TIMESTAMP NULL DEFAULT NULL,    -- Untuk soft delete
    is_verified BOOLEAN DEFAULT FALSE          -- Verifikasi status user, default FALSE
);