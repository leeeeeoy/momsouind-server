package models

type Baby struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Birth     string `json:"birth"`
	Selected  bool   `json:"selected"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
