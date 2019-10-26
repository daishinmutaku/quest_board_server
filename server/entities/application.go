package entities

type Application struct {
	Id          int64
	Status      bool
	QuestId     int64
	Quest       Quest
	ApplicantId int64
	Applicant   User
}
