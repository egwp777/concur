# Concur

Simple Go script that adds dummy data to a users table. This is a test script to test out concurrency with goroutines, 
as well as how long would it take to go through records and update the values based on a simple condition.


Command:

Options: -- 
-- update
-- created

```
go run main.go update
```

The update portion of this logic adds 70k user entries with dummy data


```
username    string
status      enum('active', 'inactive', 'deleted')
created_on  datetime 
updated_on  datetime
```

MySql table schema:

```
CREATE TABLE users (
id INT(12) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
username varchar(50),
status enum('active', 'inactive', 'deleted') DEFAULT 'inactive',
created_on DATETIME DEFAULT NOW(),
updated_on DATETIME
)
ENGINE=INNODB, charset=utf8;
```
