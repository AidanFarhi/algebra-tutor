package service

import "algtutor/repo"

type CurriculumService struct {
	cr *repo.CurriculumRepo
}

func NewCurriculumService(cr *repo.CurriculumRepo) *CurriculumService {
	return &CurriculumService{cr}
}
