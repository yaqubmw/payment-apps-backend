CREATE TABLE role
(
    id character varying(5) NOT NULL,
    role character varying(100) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE customers
(
    id character varying(100) NOT NULL,
    name character varying(100) NOT NULL,
    username character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    role_id character varying(5) NOT NULL,
    PRIMARY KEY (id)
    ADD FOREIGN KEY (roleid) REFERENCES role (id)
);

CREATE TABLE merchants
(
    id character varying(100) NOT NULL,
    name character varying(100) NOT NULL,
    address character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    phone character varying(15) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE public.tx_transaction
(
    id character varying(100) NOT NULL,
    customer_id character varying(100) NOT NULL,
    merchant_id character varying(100) NOT NULL,
    amount bigint NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (customerid) REFERENCES customers (id),
    FOREIGN KEY (merchantid) REFERENCES merchants (id)
);