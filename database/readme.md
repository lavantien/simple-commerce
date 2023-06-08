# Simple Commerce Database

## Diagram

![Sample](schemas/sample.png)

## Models

```sql
CREATE TABLE user (
  id integer PRIMARY KEY,
  name text NOT NULL,
  email text NOT NULL,
  password text NOT NULL,
  role text NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE TABLE review (
  id integer PRIMARY KEY,
  user_id integer NOT NULL REFERENCES user (id),
  product_id integer NOT NULL REFERENCES product (id),
  rating integer NOT NULL,
  comment text NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE TABLE product (
  id integer PRIMARY KEY,
  name text NOT NULL,
  price numeric NOT NULL,
  description text NOT NULL,
  picture text NOT NULL
  category_id integer NOT NULL REFERENCES category (id),
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE TABLE inventory (
  id serial PRIMARY KEY,
  product_id integer NOT NULL REFERENCES product (id),
  quantity integer NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE TABLE category (
  id serial PRIMARY KEY,
  name text NOT NULL,
  description text NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE TABLE cart (
  id integer PRIMARY KEY,
  user_id integer NOT NULL REFERENCES user (id),
  product_id integer NOT NULL REFERENCES product (id),
  quantity integer NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE TABLE order (
  id integer PRIMARY KEY,
  user_id integer NOT NULL REFERENCES user (id),
  product_id integer NOT NULL REFERENCES product (id),
  quantity integer NOT NULL,
  total_price numeric NOT NULL,
  payment_id integer NOT NULL REFERENCES payment (id),
  shipping_id integer NOT NULL REFERENCES shipping (id),
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE TABLE payment (
  id serial PRIMARY KEY,
  order_id integer NOT NULL REFERENCES order (id),
  payment_method text NOT NULL,
  amount numeric NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE TABLE shipping (
  id serial PRIMARY KEY,
  order_id integer NOT NULL REFERENCES order (id),
  shipping_method text NOT NULL,
  tracking_number text,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);
```
