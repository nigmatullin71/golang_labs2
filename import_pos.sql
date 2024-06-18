-- Создание таблицы position с автоинкрементом для id
CREATE TABLE position (
    id INT AUTO_INCREMENT PRIMARY KEY,
    earthpos VARCHAR(100),
    sunPosition VARCHAR(100),
    moonPosition VARCHAR(100)
);

-- Вставка данных в таблицу position без указания id
INSERT INTO position (earthpos, sunPosition, moonPosition) VALUES
('EarthPos1', 'SunPos1', 'MoonPos1'),
('EarthPos2', 'SunPos2', 'MoonPos2'),
('EarthPos3', 'SunPos3', 'MoonPos3'),
('EarthPos4', 'SunPos4', 'MoonPos4'),
('EarthPos5', 'SunPos5', 'MoonPos5');

DELIMITER //
CREATE PROCEDURE JoinTables(IN table1 VARCHAR(50), IN table2 VARCHAR(50))
BEGIN
    SET @query = CONCAT('SELECT * FROM ', table1, ' JOIN ', table2, ' ON ', table1, '.id = ', table2, '.', table1, '_id');
    PREPARE stmt FROM @query;
    EXECUTE stmt;
    DEALLOCATE PREPARE stmt;
DELIMITER ;
