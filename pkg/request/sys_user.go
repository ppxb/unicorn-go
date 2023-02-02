package request

type CreateUser struct {
	Mobile   string `json:"mobile" example:"18111111111"`
	Password string `json:"password"`
	Name     string `json:"name" example:"张三"`
	RoleId   uint   `json:"roleId"`
}
