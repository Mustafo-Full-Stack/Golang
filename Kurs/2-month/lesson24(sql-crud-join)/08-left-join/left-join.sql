-- Шаг 8. INNER JOIN и LEFT JOIN - в чём разница (показываем в psql)

-- INNER JOIN: только те пользователи, у кого ЕСТЬ заказы
SELECT u.name, COUNT(o.id) AS orders_count
FROM users u
INNER JOIN orders o ON u.id = o.user_id
GROUP BY u.id, u.name
ORDER BY u.id;

-- LEFT JOIN: ВСЕ пользователи, у тех, у кого заказов нет, - ноль (а сумма - NULL)
SELECT u.name, COUNT(o.id) AS orders_count, SUM(o.amount) AS total
FROM users u
LEFT JOIN orders o ON u.id = o.user_id
GROUP BY u.id, u.name
ORDER BY total DESC NULLS LAST;
