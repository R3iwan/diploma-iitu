package models

type BaseUser struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

type Admin struct {
	AdminID int `json:"admin_id"`
	BaseUser
}

type Manager struct {
	ManagerID int `json:"manager_id"`
	BaseUser
}

type Employee struct {
	EmployeeID int `json:"employee_id"`
	BaseUser
}
