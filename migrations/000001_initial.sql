-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS orders (
                                      order_uid          TEXT PRIMARY KEY,
                                      track_number       TEXT NOT NULL,
                                      entry              TEXT,
                                      locale             TEXT,
                                      internal_signature TEXT,
                                      customer_id        TEXT,
                                      delivery_service   TEXT,
                                      shardkey           TEXT,
                                      sm_id              INT,
                                      date_created       TIMESTAMP,
                                      oof_shard          TEXT
);

CREATE TABLE IF NOT EXISTS delivery (
                                        order_uid   TEXT PRIMARY KEY REFERENCES orders(order_uid) ON DELETE CASCADE,
    name        TEXT,
    phone       TEXT,
    zip         TEXT,
    city        TEXT,
    address     TEXT,
    region      TEXT,
    email       TEXT
    );

CREATE TABLE IF NOT EXISTS payments (
                                        transaction   TEXT PRIMARY KEY REFERENCES orders(order_uid) ON DELETE CASCADE,
    request_id    TEXT,
    currency      TEXT,
    provider      TEXT,
    amount        NUMERIC,
    payment_dt    BIGINT,
    bank          TEXT,
    delivery_cost INT,
    goods_total   INT,
    custom_fee    INT
    );

CREATE TABLE IF NOT EXISTS items (
                                     id           SERIAL PRIMARY KEY,
                                     order_uid    TEXT REFERENCES orders(order_uid) ON DELETE CASCADE,
    chrt_id      BIGINT,
    track_number TEXT,
    price        NUMERIC,
    rid          TEXT,
    name         TEXT,
    sale         INT,
    size         TEXT,
    total_price  NUMERIC,
    nm_id        BIGINT,
    brand        TEXT,
    status       INT
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS delivery;
DROP TABLE IF EXISTS orders;

-- +goose StatementEnd
