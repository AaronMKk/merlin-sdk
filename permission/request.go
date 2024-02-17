package permission

type RequestToCheckPermission struct {
	User    string `json:"user"`
	Object  string `json:"object"`
	ObjType string `json:"obj_type"`
	Action  int    `json:"action"`
}
