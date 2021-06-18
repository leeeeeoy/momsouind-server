package models

type RecordData struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	Filename   string `json:"file_name"`
	RecordDate string `json:"record_date"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
