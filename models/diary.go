package models

type Diary struct {
	ID                int    `json:"id"`
	UserID            int    `json:"user_id"`
	EmotionalState    int    `json:"emotional_state"`
	PhysicalCondition int    `json:"physical_condition"`
	Text              string `json:"text"`
	DiaryDate         string `json:"diary_date"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}
