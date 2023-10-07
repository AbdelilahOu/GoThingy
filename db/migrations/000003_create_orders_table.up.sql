
-- Create the "orders" table
CREATE TABLE IF NOT EXISTS orders (
  id TEXT NOT NULL PRIMARY KEY,
  status TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  client_id TEXT NOT NULL,
  CONSTRAINT orders_client_id_fkey FOREIGN KEY (client_id) REFERENCES clients (id) ON DELETE CASCADE ON UPDATE CASCADE
);