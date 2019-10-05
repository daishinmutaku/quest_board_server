package models

type Application struct {
	Id        int64
	Status    int64
	Quest     Quest
	Applicant User
}
