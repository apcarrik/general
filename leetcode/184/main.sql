# Write your MySQL query statement below

SELECT Department.name as Department, E3.name as Employee, E3.salary as Salary
FROM Department,
	(SELECT E2.name, E2.salary, E2.id, E2.departmentId
	FROM Employee as E2,
		(SELECT MAX(salary) as topSalary, departmentId
		FROM Employee
		GROUP BY departmentId) as E1
	WHERE E2.salary = E1.topSalary AND E2.departmentId = E1.departmentId) as E3
WHERE Department.id = E3.departmentId
