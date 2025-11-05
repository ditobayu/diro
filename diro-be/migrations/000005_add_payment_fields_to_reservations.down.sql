-- Migration: add_payment_fields_to_reservations
-- Created at:

-- Write your down migration here
ALTER TABLE reservations DROP COLUMN payment_status;
ALTER TABLE reservations DROP COLUMN invoice_url;
ALTER TABLE reservations DROP COLUMN payment_id;
