package entity

type (
	RTPState struct {
		TotalBet int64 `json:"total_bet" bson:"total_bet"`
		TotalWin int64 `json:"total_win" bson:"total_win"`
	}
)
