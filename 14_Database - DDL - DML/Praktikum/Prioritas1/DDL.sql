CREATE TABLE users (
  id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  nama varchar(255) NOT NULL,
  alamat_id int(11) NOT NULL,
  status_user smallint(20) NOT NULL,
  gender varchar(20) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp(),
  update_at timestamp NOT NULL DEFAULT current_timestamp(),
  tanggal_lahir date NOT NULL,
  foreign key(alamat_id) references alamat(id)
  
);

CREATE TABLE product (
  id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  nama varchar(255) NOT NULL,
  product_type_id int(11) NOT NULL,
  operator_id int(11) NOT NULL,
  code varchar(255) NOT NULL,
  status smallint(6) NOT NULL,
  created_at timestamp NULL DEFAULT current_timestamp(),
  update_at timestamp NOT NULL DEFAULT current_timestamp()
);

CREATE TABLE product_type (
  id int(11) NOT NULL  AUTO_INCREMENT PRIMARY KEY,
  nama varchar(255) NOT NULL,
  created_at timestamp NULL DEFAULT current_timestamp(),
  update_at timestamp NOT NULL DEFAULT current_timestamp()
);

CREATE TABLE operators (
  id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  nama varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp(),
  update_at timestamp NOT NULL DEFAULT current_timestamp()
);
 
CREATE TABLE product_description (
  id int(11) NOT NULL  AUTO_INCREMENT PRIMARY KEY,
  description text NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp(),
  update_at timestamp NOT NULL DEFAULT current_timestamp()
);

CREATE TABLE payment_method (
  id int(11) NOT NULL  AUTO_INCREMENT PRIMARY KEY,
  nama varchar(255) NOT NULL,
  update_at timestamp NOT NULL DEFAULT current_timestamp(),
  created_at timestamp NOT NULL DEFAULT current_timestamp()
  
);

CREATE TABLE transactions (
  id int(11) NOT NULL  AUTO_INCREMENT PRIMARY KEY,
  user_id int(11) NOT NULL,
  payment_method_id int(11) NOT NULL,
  status varchar(255) NOT NULL,
  total_price int(11) NOT NULL,
  total_qnty int(11) NOT NULL,
  update_at timestamp NOT NULL DEFAULT current_timestamp(),
  created_at timestamp NOT NULL DEFAULT current_timestamp()
);

CREATE TABLE transaction_detail (
  id int(11) NOT NULL  AUTO_INCREMENT PRIMARY KEY,
  transaction_id int(11) NOT NULL,
  qnty int(11) NOT NULL,
  price int(11) NOT NULL,
  status varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT current_timestamp(),
  update_at timestamp NOT NULL DEFAULT current_timestamp()
);

CREATE TABLE kurir (
  id int(11) NOT NULL  AUTO_INCREMENT PRIMARY KEY,
  nama varchar(255) NOT NULL,
  update_at timestamp NOT NULL DEFAULT current_timestamp(),
  created_at timestamp NOT NULL DEFAULT current_timestamp()
);

ALTER TABLE kurir ADD COLUMN ongkos_dasar INT(11);
ALTER TABLE kurir RENAME shipping;
DROP TABLE shipping;

CREATE TABLE alamat (
	id int(11) not null AUTO_INCREMENT PRIMARY KEY,
    alamat varchar(255) not null,
    update_at timestamp not null default current_timestamp(),
	created_at timestamp NOT NULL DEFAULT current_timestamp()
);

CREATE TABLE payment_method_description (
	ID INT(11) NOT NULL  AUTO_INCREMENT PRIMARY KEY,
    PAYMENT_METHOD_ID INT(11) NOT NULL,
    update_at timestamp not null default current_timestamp(),
    created_at timestamp NOT NULL DEFAULT current_timestamp(),
	foreign key(payment_method_id) references payment_method(id)
);	

CREATE TABLE user_payment_method_detail (
	ID INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    payment_method_id int(11) not null,
    user_id INT(11) NOT NULL,
	update_at timestamp not null default current_timestamp(),
    created_at timestamp NOT NULL DEFAULT current_timestamp(),
    foreign key(payment_method_id) references payment_method(id),
    foreign key(user_id) references users(id)
    
);
-- 1 insert (A)
insert into operator values('1','OPPO', now(), now());
insert into operator values('2','IPONE', now(), now());
insert into operator values('3','ASUS', now(), now());
insert into operator values('4','SAMSUNG', now(), now());
insert into operator values('5','MITO', now(), now());

-- insert (B)
INSERT INTO product_type(id,nama) values (1,'MURAH'),(2,'MAHAL'),(3,'SEDANG');

-- insert (C)
insert into product(id,nama,product_type_id,operator_id,code,status) 
 values 
	(1,'IPHONE 11',1,3,'IP11',1),
	(2,'SAMSUNG S8',1,3,'S8',2);

-- insert (D)
 insert into product(id,nama,product_type_id,operator_id,code,status) 
 values 
	(3,'IPHONE 12',2,1,'IP12',3),
	(4,'OPPO RENO',2,1,'RENO',4);

-- insert (E)
INSERT INTO product (nama, product_type_id, operator_id, code, status) VALUES
('SAMSUNG J2', 3, 4, 'J2', 1),
('SAMSUNG S3', 3, 4, 'S3', 1),
('SAMSUNG J1', 3, 4, 'J1', 1);

-- insert (F)
insert into product_description (id,description)
values 
(1,'Memiliki RAM 2GB'),
(2,'Memiliki RAM 3GB'),
(3,'Memiliki RAM 4GB'),
(4,'Memiliki RAM 6GB'),
(5,'Memiliki RAM 8GB'),
(6,'Memiliki RAM 12GB'),
(7,'Memiliki RAM 14GB');

-- insert (G)
insert into payment_method (id,nama) 
values
(1,'GOPAY'),
(2,'SHOPPEPAY'),
(3,'OVO');

-- insert (H)
 INSERT INTO users (nama, alamat_id, status_user, gender, tanggal_lahir)
VALUES 
('Rini', 1, 1, 'M', '1995-05-21'),
('Budi', 2, 1, 'W', '1990-10-10'),
('Dewi', 3, 1, 'M', '1985-02-14'),
('Rizki', 4, 1, 'M', '2000-12-31'),
('Sari', 5, 1, 'W', '1998-07-07');

-- insert (G)

-- transaksi 1
INSERT INTO transactions(user_id, payment_method_id, status, total_price, total_qnty) VALUES (16, 15, 'completed', 50000, 2);
SET @transaction_id1 = LAST_INSERT_ID();
-- produk 1 pada transaksi 1
INSERT INTO transaction_detail(transaction_id, qnty, price, status) VALUES (@transaction_id1, 1, 25000, 'completed');
-- produk 2 pada transaksi 1
INSERT INTO transaction_detail(transaction_id, qnty, price, status) VALUES (@transaction_id1, 1, 25000, 'completed');

-- transaksi 2
INSERT INTO transactions(user_id, payment_method_id, status, total_price, total_qnty) VALUES (16, 15, 'completed', 75000, 3);
SET @transaction_id2 = LAST_INSERT_ID();
-- produk 1 pada transaksi 2
INSERT INTO transaction_detail(transaction_id, qnty, price, status) VALUES (@transaction_id2, 2, 25000, 'completed');
-- produk 2 pada transaksi 2
INSERT INTO transaction_detail(transaction_id, qnty, price, status) VALUES (@transaction_id2, 1, 25000, 'completed');
-- produk 3 pada transaksi 2
INSERT INTO transaction_detail(transaction_id, qnty, price, status) VALUES (@transaction_id2, 1, 25000, 'completed');

-- transaksi 3
INSERT INTO transactions(user_id, payment_method_id, status, total_price, total_qnty) VALUES (16, 15, 'completed', 100000, 4);
SET @transaction_id3 = LAST_INSERT_ID();
-- produk 2 pada transaksi 3
INSERT INTO transaction_detail(transaction_id, qnty, price, status) VALUES (@transaction_id3, 2, 25000, 'completed');
-- produk 3 pada transaksi 3
INSERT INTO transaction_detail(transaction_id, qnty, price, status) VALUES (@transaction_id3, 1, 25000, 'completed');
-- produk 4 pada transaksi 3
INSERT INTO transaction_detail(transaction_id, qnty, price, status) VALUES (@transaction_id3, 1, 25000, 'completed');

-- Tampilkan nama user / pelanggan dengan gender Laki-laki / M. -- 

select * from users where gender='M';



-- Tampilkan product dengan id = 3. --

select * from product where id='3';



-- Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’. --

select * from users where created_at > date_sub(now(), interval 1 week) AND nama LIKE '%a%';



-- Hitung jumlah user / pelanggan dengan status gender Perempuan. --

 select sum(gender='') from users;
 
 
 
 -- Tampilkan data pelanggan dengan urutan sesuai nama abjad --
 
 select * from users order by nama ASC;


-- Tampilkan 5 data pada data product -- 

select * from product limit 5;



-- Ubah data product id 1 dengan nama ‘product dummy’. --

update product set nama='product dummy' where id='1'; 




-- Update qty = 3 pada transaction detail dengan product id = 1. --

update transaction_details set quanty='3' where product_id='1'; 




-- Delete data pada tabel product dengan id = 1. --
-- Delete pada pada tabel product dengan product type id 1. -- 

update product set deleted_at = now() where id = 1 OR product_type_id =1;






-- Gabungkan data transaksi dari user id 1 dan user id 2. --

 select * from transaction where users_id in(1,2);
 
 
 
 -- Tampilkan jumlah harga transaksi user id 1. -- 
 
 select users_id, Sum(total_price) as total_price from transaction where users_id = 1;


-- Tampilkan semua field table product dan field name table product type yang saling berhubungan. -- 

select count(*) from transaction t
join transaction_details td on t.id = td.transaction_id
join product p on p.id = td.product_id
where p.product_type_id = 2;


-- Tampilkan semua field table transaction, field name table product dan field name table user. --

select t.*, p.nama as product, u.nama as users from transaction_details td
join transaction t on t.id = td.transaction_id
join users u on u.id = t.users_id
join product p on p.id = td.product_id
where p.deleted_at is null
 
 
 -- Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud. --
 -- delete row dorign key -- 
 delimiter $$
 create trigger delete_transaction_details
 before delete on transaction for each row
 begin
 declare trans_id int;
 set trans_id = old.id;
 delete from transaction_details where transaction_id = trans_id;
 end $$
 
 delete from transaction where id = 4;
 
 
 
 -- Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus. --
 
 delimiter $$
 create trigger update_transaction_id
 after delete on transaction_details for each row
 begin
 declare trans_id int;
 set trans_id = old.transaction_id;
 update transaction set total_quanty = (select sum(quanty) from transaction_details where transaction_id = trans_id)
 where id = trans_id;
 end $$
 
 select * from transaction;
 
 select * from transaction_details where transaction_id = 15;
 
 delete from transaction_details where transaction_id = 15 AND product_id = 4;
 
 
 
 -- Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query. -- 
 
 select * from product where id not in(
 select product_id from transaction_details
 );
 
select * from product where id in (
select p.id from product p
left outer join transaction_details td on p.id = td.product_id
where td.transaction_id is null
);









