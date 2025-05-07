-- customers 表
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_active_at TIMESTAMP
);

-- purchases 表
CREATE TABLE IF NOT EXISTS purchases (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id) ON DELETE CASCADE,
    amount DECIMAL NOT NULL,
    purchased_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- customer_tags 表
CREATE TABLE IF NOT EXISTS customer_tags (
    customer_id INT REFERENCES customers(id) ON DELETE CASCADE,
    tag VARCHAR(50),
    PRIMARY KEY (customer_id, tag)
);

-- coupons 表
CREATE TYPE coupon_type AS ENUM ('discount', 'cash');
CREATE TABLE IF NOT EXISTS coupons (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    type coupon_type NOT NULL,
    value DECIMAL NOT NULL,
    quantity INT NOT NULL,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- user_coupons 表
CREATE TYPE coupon_status AS ENUM ('unused', 'used', 'expired');
CREATE TABLE IF NOT EXISTS user_coupons (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    coupon_id INT REFERENCES coupons(id) ON DELETE CASCADE,
    status coupon_status NOT NULL,
    claimed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    used_at TIMESTAMP
);