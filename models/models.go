package models

//участник
type Participant struct {
	ID   int64
	Name string
}

//розыгрыш
type Riffle struct {
	Participants []Participant
	DrawTime     int64 //time Unix
	IsActive     bool
}

type Channel struct {
	OwnerID int64
	Raffle  *Riffle
}
