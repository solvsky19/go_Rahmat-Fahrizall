CREATE TABLE users (
    id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    status SMALLINT,
    dob DATE,
    gender char(1) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE payment_methods (
    id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE operators (
    id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE product_types (
    id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE products (
    id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    product_type_id int(11),
    operator_id int(11),
    code varchar(50),
    name varchar(255),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(product_type_id) REFERENCES product_types(id),
    FOREIGN KEY(operator_id) REFERENCES operators(id)
);

CREATE TABLE product_descriptions (
    id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    product_id int(11),
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(product_id) REFERENCES products(id)
);

CREATE TABLE transactions (
    id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    user_id int(11),
    payment_method_id int(11),
    status varchar(255),
    total_qty int(11),
    total_price NUMERIC(25,2),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(payment_method_id) REFERENCES payment_methods(id)
);

CREATE TABLE transaction_details (
    id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    transaction_id int(11),
    product_id int(11),
    status varchar(255),
    qty int(11),
    price NUMERIC(25,2),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(transaction_id) REFERENCES transactions(id),
    FOREIGN KEY(product_id) REFERENCES products(id)
);

CREATE TABLE kurir (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(255),
    ongkos_dasar int(50),
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);

-- Rename tabel kurir menjadi shipping --
Alter table kurir rename shipping;

-- Delete table shipping --
Drop table shipping;

CREATE TABLE payment_method_descriptions (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    payment_method_id int(11),
	description TEXT,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW(),
	FOREIGN KEY(payment_method_id) REFERENCES payment_methods(id)
);

CREATE TABLE alamat (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    user_id int(11),
	alamat TEXT,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE user_payment_method_details (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
	user_id int(11),
    payment_method_id int(11),
    nama varchar(255),
    created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(payment_method_id) REFERENCES payment_methods(id)
);

Alter table products add column deleted_at timestamp;
Alter table users add column name varchar(255);