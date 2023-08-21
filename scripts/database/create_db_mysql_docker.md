# Crear y arrancar servidor MySQL con Docker
```
docker run -d -p 3306:3306 --name='mysql-server' --env="MYSQL_ROOT_PASSWORD=password" mysql --default-authentication-plugin=mysql_native_password
```

# Ingresar al servidor MySQL Docker
```
docker exec -ti mysql-server bash
```

# Primera Configuracion DB
```
mysql -u root -p
create database bia_db;
create user 'bia_user'@'%%' identified with mysql_native_password BY 'password';
grant all privileges on bia_db.* TO 'bia_user'@'%%';
flush privileges;
quit;
```