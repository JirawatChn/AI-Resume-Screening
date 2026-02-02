package resume

type Resume struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Score int    `json:"score"` // คะแนนที่ AI ประเมิน
}
