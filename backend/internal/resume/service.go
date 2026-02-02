package resume

type Service interface {
	ScreenResume(resume Resume) (Resume, error)
}

type resumeService struct {
	repo Repository
}

func NewService(r Repository) Service { return &resumeService{repo: r} }

func (s *resumeService) ScreenResume(r Resume) (Resume, error) {
	// Business Logic: เช่น ถ้ามีคำว่า "Go" ให้คะแนน 100
	r.Score = 100

	err := s.repo.Save(r)
	return r, err
}
