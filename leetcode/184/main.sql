# Write your MySQL query statement below

SELECT Department, Employee.name as Employee, topSalary as Salary
FROM Employee,
    (SELECT Department.name as Department, Employee.departmentId, MAX(Employee.salary) as topSalary
    FROM Department LEFT JOIN Employee
    ON Department.id = Employee.departmentId
    GROUP BY Department.id) as T1
WHERE T1.topSalary = Employee.salary AND T1.departmentId = Employee.departmentId
