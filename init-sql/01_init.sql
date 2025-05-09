-- 建立 ENUM
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'coupon_type') THEN
        CREATE TYPE coupon_type AS ENUM ('discount', 'fill');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'coupon_status') THEN
        CREATE TYPE coupon_status AS ENUM ('unused', 'used', 'expired');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'customer_level') THEN
        CREATE TYPE level_type AS ENUM ('NEW', 'VIP', 'SLIVER', 'COPPER');
    END IF;
END$$;

-- customer 表
CREATE TABLE IF NOT EXISTS customer (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    area_code VARCHAR(6),
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(20),
    customer_level level_type NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_active_at TIMESTAMP
);

-- purchase 表
CREATE TABLE IF NOT EXISTS purchase (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customer(id) ON DELETE CASCADE,
    amount DECIMAL NOT NULL,
    purchased_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- coupon 表
CREATE TABLE IF NOT EXISTS coupon (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    type coupon_type NOT NULL,
    value DECIMAL NOT NULL,
    quantity INT NOT NULL,
    coupon_level level_type NOT NULL,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- user_coupon 表
CREATE TABLE IF NOT EXISTS user_coupon (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    coupon_id INT REFERENCES coupon(id) ON DELETE CASCADE,
    status coupon_status NOT NULL,
    claimed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    used_at TIMESTAMP
);