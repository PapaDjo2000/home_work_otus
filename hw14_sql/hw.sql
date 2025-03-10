create schema hw;
create extension if not exists  "uuid-ossp";

create table if not exists hw.users(
id uuid default uuid_generate_v4() primary key,
name varchar(255)not null,
password varchar(255)not null,
email varchar(255)not nul						l
);

create table if not exists hw.orders(
id uuid default uuid_generate_v4()primary key,
user_id uuid references hw.users (id) on delete cascade,
order_date timestamptz not null default now(),
total_amount int not null
);

create table if not exists hw.products(
id uuid default uuid_generate_v4() primary key,
name varchar(255)not null,
price int not null default 0
);

create table hw.order_products(
order_id uuid references hw.orders(id) on delete cascade not null ,
products_id uuid references hw.products(id)on delete cascade  not null,
quantity int default 0
);

insert into hw.users (name,password,email) values ('dima','rth','dimaos@mail.ru');

insert into hw.products  (name,price) values ('honor',430);

update hw.users set name='oleg', password='oleg142', email='oleg@mail.ru' where id='9c9d5249-7f25-4d01-b722-457a7a88f22f';

delete from hw.users where name='oleg';

delete from hw.products  where id='c9cc528a-d212-4e0a-9bbd-dc57cbc632b4';

insert into hw.orders (user_id,total_amount) values ('c0c732ea-6a66-430e-8ac6-7a2994e6ad3a' ,300);

delete from hw.orders where user_id='dddce09b-a378-4e06-8501-9d0d47948fc6';


insert into hw.order_products (order_id ,products_id, quantity) values ('07474e09-fce0-462a-aa3f-e68a8b50a455','755305b0-b16f-4174-a7a8-967440470cab',2);



select * from hw.products ;

select * from hw.users ;

select * from hw.orders where user_id='dddce09b-a378-4e06-8501-9d0d47948fc6';

select u.id ,u.name, o.id,o.order_date, o.total_amount
from hw.orders  o
inner join 
hw.users u on u.id=o.user_id 


select u.id ,u.name, o.id,o.order_date, o.total_amount
from hw.orders  o
inner join 
hw.users u on u.id=o.user_id 
where u.id='dddce09b-a378-4e06-8501-9d0d47948fc6'


select u.id ,  u.name , SUM(op.quantity * p.price)
from  hw.users u
inner join 
    hw.orders o ON u.id = o.user_id
inner join 
    hw.order_products op ON o.id = op.order_id
inner join 
    hw.products p ON op.products_id = p.id
where
    u.id = 'dddce09b-a378-4e06-8501-9d0d47948fc6'
group by
    u.id, u.name;

select u.id ,  u.name , avg(p.price)
from  hw.users u
inner join 
    hw.orders o ON u.id = o.user_id
inner join 
    hw.order_products op ON o.id = op.order_id
inner join 
    hw.products p ON op.products_id = p.id
where
    u.id = 'dddce09b-a378-4e06-8501-9d0d47948fc6'
group by
    u.id, u.name;

create unique index if not exists idx_users_id on hw.users(id );

create unique index if not exists idx_orders_id on hw.orders (id);

create unique index if not exists idx_product_id on hw.products (id);