-- Шаг 2. Учимся выбирать данные (выполняем в psql, по одному запросу)
-- После каждого запроса - пауза: "что мы увидим?"

-- Только нужные столбцы и условие
SELECT name, age FROM users WHERE age > 20;

-- Сортировка (DESC - по убыванию)
SELECT * FROM users ORDER BY age DESC;

-- Два самых молодых
SELECT * FROM users ORDER BY age LIMIT 2;

-- Сколько всего строк
SELECT COUNT(*) FROM users;

-- Агрегаты: среднее, минимум, максимум
SELECT AVG(age), MIN(age), MAX(age) FROM users;

-- Поиск по шаблону: имена на букву "А"
SELECT * FROM users WHERE name LIKE 'А%';

-- Диапазон
SELECT * FROM users WHERE age BETWEEN 20 AND 28;

-- Несколько условий сразу
SELECT * FROM users WHERE email IS NOT NULL AND is_active = true;

-- Дайте студентам самим написать пару запросов, например:
--   всех, кто старше 25, отсортированных по имени
--   сколько пользователей младше 26
