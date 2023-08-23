DELIMITER //
CREATE PROCEDURE InsertRandomData()
BEGIN
    DECLARE i INT DEFAULT 0;

    WHILE i < 40000000 DO
    INSERT INTO users (name, date_of_birth) VALUES (CONCAT('User', i), DATE_ADD('1980-01-01', INTERVAL FLOOR(RAND() * (365 * 30)) DAY));
    SET i = i + 1;
    END WHILE;
    END;
    //
DELIMITER ;


CALL InsertRandomData();
