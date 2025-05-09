INSERT INTO customer (name, email, area_code, phone, customer_level, last_active_at) VALUES
('Alice Chen', 'alice@example.com', '886','0912345678', 'NEW', NOW() - INTERVAL '5 days'),
('Bob Wang', 'bob@example.com', '886','0922333444', 'SLIVER', NOW() - INTERVAL '10 days'),
('Charlie Lin', 'charlie@example.com', '886','0933444555', 'VIP', NOW() - INTERVAL '2 days'),
('Doris Wu', 'doris@example.com', '886','0966778899', 'NEW', NOW() - INTERVAL '1 day'),
('Eric Tsai', 'eric@example.com', '886', '0955667788', 'SLIVER', NOW());

-- purchase
INSERT INTO purchase (customer_id, amount, purchased_at) VALUES
(1, 440.00, NOW() - INTERVAL '2 days'),
(1, 150.00, NOW() - INTERVAL '29 day'),
(2, 300.00, NOW() - INTERVAL '25 days'),
(3, 600.00, NOW() - INTERVAL '10 days'),
(4, 50.00, NOW() - INTERVAL '1 day'),
(5, 399.00, NOW() - INTERVAL '33 days'),
(5, 299.00, NOW() - INTERVAL '3 days');


-- coupon
INSERT INTO coupon (name, type, value, quantity, coupon_level, start_time, end_time) VALUES
('100元折扣券', 'discount', 100, 10, 'NEW',NOW() - INTERVAL '1 day', NOW() + INTERVAL '7 days'),
('9折優惠券', 'discount', 0.9, 5, 'SLIVER', NOW(), NOW() + INTERVAL '10 days'),
('5折優惠券', 'discount', 0.5, 2, 'VIP', NOW(), NOW() + INTERVAL '3 days'),
('買2送1', 'fill', 200, 1, 'COPPER', NOW() - INTERVAL '1 days', NOW() - INTERVAL '6 day'),
('1000元限時折扣', 'discount', 1000, 3, 'SLIVER',NOW() - INTERVAL '3 days', NOW() - INTERVAL '5 day');




-- user_coupon
INSERT INTO user_coupon (user_id, coupon_id, status, claimed_at, used_at) VALUES
(1, 1, 'unused', NOW() - INTERVAL '1 day', NULL),
(2, 2, 'used', NOW() - INTERVAL '2 days', NOW() - INTERVAL '1 day'),
(3, 3, 'expired', NOW() - INTERVAL '4 days', NULL),
(4, 1, 'unused', NOW() - INTERVAL '2 hours', NULL),
(5, 2, 'unused', NOW() - INTERVAL '1 hour', NULL);