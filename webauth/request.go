package webauth

type RequestToAuth struct {
	IP        string `json:"ip"`
	LoginId   string `json:"login_id"`
	CSRFToken string `json:"csrf_token"`
	UserAgent string `json:"user_agent"`
}

type ResponseToAuth struct {
	User      string `json:"user"`
	CSRFToken string `json:"csrf_token"`
}
