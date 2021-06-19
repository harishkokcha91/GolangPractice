# GolangPractice



CREATE TABLE test.itemDetails (
	item_id INT auto_increment NOT NULL,
	item_name varchar(100) NULL,
	item_price varchar(100) NULL,
	PRIMARY KEY (item_id)
	
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;



CREATE TABLE test.itemsOffer (
	id INT auto_increment NOT NULL,
	item_name varchar(100) NULL,
	item_qty varchar(100) NULL,
	item_price varchar(100) NULL,
	item_total_price varchar(100) NULL,
	item_discount varchar(100) NULL,
	PRIMARY KEY (id)
	
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;


go mod init


go get github.com/go-sql-driver/mysql


