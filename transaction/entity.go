package transaction

import "time"

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Stauts     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
