DROP DATABASE IF EXISTS `random_message_prod`;

CREATE DATABASE random_message_prod;

DROP USER IF EXISTS 'random_message_prod' @'%';

CREATE USER 'random_message_prod' IDENTIFIED BY 'strongpassword';

GRANT ALL ON random_message_prod.* TO 'random_message_prod' @'%';

USE `random_message_prod`;

CREATE TABLE messages (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
message VARCHAR(500) NOT NULL
);

INSERT INTO
    messages (message)
VALUES
    ("Try not. Do or do not. There is no try. Yoda"),
    (
        "Who\'s the more foolish: the fool or the fool who follows him? Obi-Wan"
    ),
    (
        "Your focus determines your reality. Qui-Gon Jinn"
),
(
    "We'll always be with you. No one's ever really gone. A thousand generations live in you now. —Luke Skywalker"
    );
