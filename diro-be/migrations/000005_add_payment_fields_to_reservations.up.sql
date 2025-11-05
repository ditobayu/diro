-- Migration: add_payment_fields_to_reservations
-- Created at:

-- Write your up migration here
ALTER TABLE reservations
ADD COLUMN payment_id VARCHAR(255) DEFAULT '',
ADD COLUMN invoice_url VARCHAR(500) DEFAULT '',
ADD COLUMN payment_status VARCHAR(50) DEFAULT '';
