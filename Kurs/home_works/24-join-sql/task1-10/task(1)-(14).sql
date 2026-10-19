-- Задача 3
SELECT s.first_name, s.last_name, c.course_name
FROM
    students AS s
JOIN courses AS c
    ON s.id = c.student_id;

-- Задача 4
SELECT s.first_name, s.last_name, c.course_name, c.price
FROM students AS s
JOIN courses AS c
    ON s.id = c.student_id;

-- Задача 5
SELECT s.first_name, c.course_name
FROM students AS s
JOIN courses AS c
    ON s.id = c.student_id
WHERE c.course_name = 'Go';

-- Задача 6
SELECT 
