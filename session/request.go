package session

type RequestToCheckAndRefresh struct {
	IP        string `json:"ip"`
	LoginId   string `json:"login_id"`
	CSRFToken string `json:"csrf_token"`
	UserAgent string `json:"user_agent"`
}

type ResponseToCheckAndRefresh struct {
	User      string `json:"user"`
	CSRFToken string `json:"csrf_token"`
}

type RequestToClear struct {
	SessionId string `json:"session_id"`
}
