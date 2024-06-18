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
CREATE PROCEDURE NameProc()
BEGIN
    SELECT * FROM position;
END //
DELIMITER ;
