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

-- Создание второй таблицы details с автоинкрементом для id и внешним ключом position_id
CREATE TABLE details (
    id INT AUTO_INCREMENT PRIMARY KEY,
    position_id INT,
    details_id INT,
    description VARCHAR(255),
    FOREIGN KEY (position_id) REFERENCES position(id)
);

-- Вставка данных в таблицу details с указанием внешнего ключа position_id
INSERT INTO details (position_id, details_id description) VALUES
(1, 'Description1'),
(2, 'Description2'),
(3, 'Description3'),
(4, 'Description4'),
(5, 'Description5');

-- Изменение разделителя перед созданием процедуры
DELIMITER //

-- Создание хранимой процедуры JoinTables с объединением таблиц
CREATE PROCEDURE JoinTables()
BEGIN
    SELECT p.id, p.earthpos, p.sunPosition, p.moonPosition, d.position_id d.details_id
    FROM position p
    JOIN details d ON p.id = d.position_id;
END //

-- Возвращение разделителя к точке с запятой
DELIMITER ;

