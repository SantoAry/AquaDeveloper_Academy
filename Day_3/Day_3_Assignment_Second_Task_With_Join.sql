create database Day_3_Assignment;

create table if not exists customers (
id int not null primary KEY,
customer_name char(50) not null
);

create table if not exists products (
id int not null primary KEY,
name char(50) not null
);

create table if not exists Orders (
order_id int not null primary KEY,
customer_id int not null,
product_id int not null,
order_date date not null,
total float not null,
constraint fk_customer
	foreign key (customer_id)
		references customers(id),
constraint fk_product
	foreign key (product_id)
		references products(id)
);

INSERT INTO public.customers (id, customer_name) VALUES (1, 'Dante');
INSERT INTO public.products (id, name) VALUES (1, 'eFeeder');
INSERT INTO public.orders (order_id, customer_id, product_id, order_date, total) VALUES (1, 1, 1, '2022-10-25', 3);

UPDATE public.customers SET customer_name = 'Vergil' WHERE id = 1;
UPDATE public.products SET name = 'Fish Pellet' WHERE id = 1;
UPDATE public.orders SET order_date = '2022-10-26' WHERE order_id = 1;

delete from public.orders where order_id = 1;
delete from public.products where id = 1;
delete from public.customers where id = 1;

select 
	o.order_id, 
	o.customer_id, 
	o.product_id, 
	o.order_date, 
	o.total 
from 
	orders o
inner join customers c  
	on c.id = o.customer_id
inner join products p 
	on p.id = o.customer_id 
order by order_date; 


select * from orders cross join customers;
select * from orders cross join products;
select * from orders cross join customers cross join products;
