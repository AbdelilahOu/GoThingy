CREATE TABLE accounts (
    id uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    owner varchar NOT NULL,
    balance bigint NOT NULL,
    currency varchar NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE entries (
    id uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    account_id uuid NOT NULL,
    amount bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE transfers (
    id uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    from_account_id uuid NOT NULL,
    to_account_id uuid NOT NULL,
    amount bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE entries ADD CONSTRAINT entries_account_id_fk FOREIGN KEY (account_id) REFERENCES accounts (id);

ALTER TABLE transfers ADD CONSTRAINT transfers_from_account_id_fk FOREIGN KEY (from_account_id) REFERENCES accounts (id);

ALTER TABLE transfers ADD CONSTRAINT transfers_to_account_id_fk FOREIGN KEY (to_account_id) REFERENCES accounts (id);

CREATE INDEX ON accounts (owner);

CREATE INDEX ON entries (account_id);

CREATE INDEX ON transfers (from_account_id);

CREATE INDEX ON transfers (to_account_id);

CREATE INDEX ON transfers (from_account_id, to_account_id);

COMMENT ON COLUMN entries.amount IS 'can be negative or positive';

COMMENT ON COLUMN transfers.amount IS 'must be positive';
