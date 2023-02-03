package request

type Login struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}
