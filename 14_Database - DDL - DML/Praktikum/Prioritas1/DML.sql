-- Bagian Insert --
-- Insert 5 operators pada table operators --
insert into operators (name) values('LAPTOP ASUS');
insert into operators (name) values('LAPTOP HP');
insert into operators (name) values('LAPTOP SAMSUNG');
insert into operators (name) values('LAPTOP ACER');
insert into operators (name) values('LAPTOP LENOVO');

-- Insert 3 product type --
insert into product_types (name) values ('LOW');
insert into product_types (name) values ('HIGT');
insert into product_types (name) values ('MEDIUM');

-- Insert 2 product dengan product type id = 1, dan operators id = 3 --
insert into products (product_type_id, operator_id, code, name, status) values ('1','3','1','LAPTOP SAMSUNG','1');
insert into products (product_type_id, operator_id, code, name, status) values ('1','3','2','LAPTOP SAMSUNG','1');

-- Insert 3 product dengan product type id = 2, dan operators id = 1 --
insert into products (product_type_id, operator_id, code, name, status) values ('2','1','3','LAPTOP ASUS','1');
insert into products (product_type_id, operator_id, code, name, status) values ('2','1','4','LAPTOP ASUS','1');
insert into products (product_type_id, operator_id, code, name, status) values ('2','1','5','LAPTOP ASUS','1');

-- Insert 3 product dengan product type id = 3, dan operators id = 4 --
insert into products (product_type_id, operator_id, code, name, status) values ('3','4','6','LAPTOP ACER','1');
insert into products (product_type_id, operator_id, code, name, status) values ('3','4','7','LAPTOP ACER','1');
insert into products (product_type_id, operator_id, code, name, status) values ('3','4','8','LAPTOP ACER','1');

-- Insert product description pada setiap product --
insert into product_descriptions (description, product_id) values ('Murah','1');
insert into product_descriptions (description, product_id) values ('Cukup Murah','2');
insert into product_descriptions (description, product_id) values ('Sangat Murah','3');
insert into product_descriptions (description, product_id) values ('Lumayan Murah','4');
insert into product_descriptions (description, product_id) values ('Lumayan Mahal','5');
insert into product_descriptions (description, product_id) values ('Mahal','6');
insert into product_descriptions (description, product_id) values ('Cukup Mahal','7');
insert into product_descriptions (description, product_id) values ('Sangat Mahal','8');

-- Insert 3 payment methods --
insert into payment_methods (name, status) values ('MANDIRI','5');
insert into payment_methods (name, status) values ('BCA','5');
insert into payment_methods (name, status) values ('BRI','5');

-- Insert 5 user pada tabel user --
insert into users (name, status, dob, gender) values('Deo','1','2001-02-08','M');
insert into users (name, status, dob, gender) values('Delfero','1','2001-05-17','M');
insert into users (name, status, dob, gender) values('Arip','1','2001-01-06','M');
insert into users (name, status, dob, gender) values('Aqsa','1','2002-06-20','W');
insert into users (name, status, dob, gender) values('Vonzy','1','2002-11-13','W');

-- Insert 3 transaksi di masing-masing user --
-- User Arif --  
insert into transactions (user_id, payment_method_id, status, total_qty, total_price) values ('1','1','READY','1','25000000');
insert into transactions (user_id, payment_method_id, status, total_qty, total_price) values ('1','2','READY','2','18000000');
insert into transactions (user_id, payment_method_id, status, total_qty, total_price) values ('1','3','READY','3','4000000');
-- User Amek --  
insert into transactions (user_id, payment_method_id, status, total_qty, total_price) values ('2','1','READY','1','25000000');
insert into transactions (user_id, payment_method_id, status, total_qty, total_price) values ('2','2','READY','2','18000000');
insert into transactions (user_id, payment_method_id, status, total_qty, total_price) values ('2','3','READY','3','4000000');
-- User Rafli --  
insert into transactions (user_id, payment_method_id, status, total_qty, total_price) values ('3','1','READY','1','25000000');
insert into transactions (user_id, payment_method_id, status, total_qty, total_price) values ('3','2','READY','2','18000000');
insert into transactions (user_id, payment_method_id, status, total_qty, total_price) values ('3','3','READY','3','4000000');
-- User YolMay --  
insert into transactions (user_id, payment_method_id, status, total_qty, total_price) values ('4','1','READY','1','25000000');
insert into transactions (user_id, payment_method_id, status, total_qty, total_price) values ('4','2','READY','2','18000000');
insert into transactions (user_id, payment_method_id, status, total_qty, total_price) values ('4','3','READY','3','4000000');
-- User Tatia --  
insert into transactions (user_id, payment_method_id, status, total_qty, total_price) values ('5','1','READY','1','25000000');
insert into transactions (user_id, payment_method_id, status, total_qty, total_price) values ('5','2','READY','2','18000000');
insert into transactions (user_id, payment_method_id, status, total_qty, total_price) values ('5','3','READY','3','4000000');

-- Insert 3 product di masing-masing transaksi --
-- Transaksi 1 --
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','1','OK','1','25000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','2','OK','2','18000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','3','OK','3','4000000');
-- Transaksi 2 --
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','4','OK','1','25000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','5','OK','2','18000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','6','OK','3','4000000');
-- Transaksi 3 --
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','7','OK','1','25000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','8','OK','2','18000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','1','OK','3','4000000');
-- Transaksi 4 --
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','2','OK','1','25000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','3','OK','2','18000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','4','OK','3','4000000');
-- Transaksi 5 --
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','5','OK','1','25000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','6','OK','2','18000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','7','OK','3','4000000');
-- Transaksi 6 --
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','8','OK','1','25000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','1','OK','2','18000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','2','OK','3','4000000');
-- Transaksi 7 --
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','3','OK','1','25000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','4','OK','2','18000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','5','OK','3','4000000');
-- Transaksi 8 --
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','6','OK','1','25000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','7','OK','2','18000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','8','OK','3','4000000');
-- Transaksi 9 --
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','1','OK','1','25000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','2','OK','2','18000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','3','OK','3','4000000');
-- Transaksi 10 --
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','4','OK','1','25000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','5','OK','2','18000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','6','OK','3','4000000');
-- Transaksi 11 --
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','7','OK','1','25000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','8','OK','2','18000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','1','OK','3','4000000');
-- Transaksi 12 --
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','2','OK','1','25000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','3','OK','2','18000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','4','OK','3','4000000');
-- Transaksi 13 --
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','5','OK','1','25000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','6','OK','2','18000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','7','OK','3','4000000');
-- Transaksi 14 --
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','8','OK','1','25000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','1','OK','2','18000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','2','OK','3','4000000');
-- Transaksi 15 --
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','3','OK','1','25000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','4','OK','2','18000000');
insert into transaction_details (transaction_id, product_id, status, qty, price) values ('1','5','OK','3','4000000');


-- Bagian Select --
-- Tampilkan nama user / pelanggan dengan gender Laki-laki / M --
select * from users where gender='M';

-- Tampilkan product dengan id = 3 --
select * from products where id='3';

-- Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’ --
select * from users where created_at > date_sub(now(), interval 1 week) AND name LIKE '%a%';

-- Hitung jumlah user / pelanggan dengan status gender Perempuan --
select sum(gender='W') from users;

-- Tampilkan data pelanggan dengan urutan sesuai nama abjad --
select * from users order by name ASC;

-- Tampilkan 5 data pada data product --
select * from products limit 5;


-- Bagian Update --
-- Ubah data product id 1 dengan nama ‘product dummy’ --
update products set name = 'product dummy' where id='1'; 

-- Update qty = 3 pada transaction detail dengan product id = 1 --
update transaction_details set qty = '3' where product_id = '1'; 


-- Bagian Delete --
-- Delete data pada tabel product dengan id = 1 --
update products set deleted_at = now() where id = 1;

-- Delete pada pada tabel product dengan product type id 1 --
update products set deleted_at = now() where product_type_id = 1;
