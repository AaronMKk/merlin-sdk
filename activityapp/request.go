package activityapp

type ReqToCreateActivity struct {
	Id           string `json:"id"`
	Owner        string `json:"owner"`
	Type         string `json:"type"`
	Time         string `json:"time"`
	ResourceId   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}

type ReqToDeleteActivity struct {
	ResourceId string `json:"resource_id"`
}
