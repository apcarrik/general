# Write your MySQL query statement below
SELECT if(id % 2 = 0, id-1, if(id = (SELECT COUNT(id) FROM Seat), id, id+1)) as id, student
FROM Seat
ORDER BY id ASC
