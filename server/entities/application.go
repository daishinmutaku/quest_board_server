package entities

type Application struct {
	Id        int64 `db:"id"`
	Status    int64 `db:"status"`
	Quest     Quest `db:"quest"`
	Applicant User  `db:"applicant"`
}
