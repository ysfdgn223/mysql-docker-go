create table users(
	id INT AUTO_INCREMENT,
    name VARCHAR(20),
    weight INT,
    PRIMARY KEY (id)
);

INSERT INTO users (name, weight) VALUES ("joseph", 77), ("Alico", 45);