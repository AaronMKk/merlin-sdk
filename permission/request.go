package permission

type RequestToCheckPermission struct {
	User  string `json:"user"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
