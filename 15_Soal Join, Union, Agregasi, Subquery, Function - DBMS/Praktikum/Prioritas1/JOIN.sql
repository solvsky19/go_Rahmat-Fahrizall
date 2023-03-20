-- Gabungkan data transaksi dari user id 1 dan user id 2 --
select * from transactions where user_id in (1, 2);

-- Tampilkan jumlah harga transaksi user id 1 --
select user_id, Sum(total_price) as total_price from transactions where user_id = 1;

-- Tampilkan total transaksi dengan product type 2 --
select count(distinct transaction_id) as total_transactions from transaction_details where product_id in (
    select id from products where product_type_id = 2
);

-- Tampilkan semua field table product dan field name table product type yang saling berhubungan --
select count(*) from transactions t join transaction_details td on t.id = td.transaction_id join products p on p.id = td.product_id where p.product_type_id = 2;

-- Tampilkan semua field table transaction, field name table product dan field name table user --
select t.*, p.name as products, u.name as users from transaction_details td
join transactions t on t.id = td.transaction_id
join users u on u.id = t.user_id
join products p on p.id = td.product_id
where p.deleted_at is null

-- Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud --
delimiter $$
 create trigger delete_transaction_details before delete on transactions for each row
 begin
 declare trans_id int;
 set trans_id = old.id;
 delete from transaction_details where transaction_id = trans_id;
 end $$
 
 delete from transactions where id = 4;
 
 -- Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus --
 delimiter $$
 create trigger update_transaction_id after delete on transaction_details for each row
 begin
 declare trans_id int;
 set trans_id = old.transaction_id;
 update transactions set total_qty = (select sum(qty) from transaction_details where transaction_id = trans_id)
 where id = trans_id;
 end $$
 
 select * from transactions;
 select * from transaction_details where transaction_id = 15;
 delete from transaction_details where transaction_id = 15 AND product_id = 4;
 
 -- Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query --
 select * from products where id not in(
 select product_id from transaction_details
 );
 
select * from products where id in (
select p.id from products p
left outer join transaction_details td on p.id = td.product_id
where td.transaction_id is null
);