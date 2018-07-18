package domain

// Author - user
type Author struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
	Github   string `json:"github"`
}
