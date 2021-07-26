create database hw02_k8s_db;
create user admin with encrypted password 'pass';
grant all privileges on database hw02_k8s_db to admin;