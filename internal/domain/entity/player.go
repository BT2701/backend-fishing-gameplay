package entity

import apperr "github.com/BT2701/backend-fishing-gameplay/pkg/error"

type (
	Player struct {
		PlayerID     string `json:"player_id" bson:"player_id"`
		Balance      int64  `json:"balance" bson:"balance"`
		SeatID       int    `json:"seat_id" bson:"seat_id"`
		GunID        int    `json:"gun_id" bson:"gun_id"`
		RoomID       string `json:"room_id" bson:"room_id"`
		SessionID    string `json:"session_id" bson:"session_id"`
		IsOnline     bool   `json:"is_online" bson:"is_online"`
		LastActionAt int64  `json:"last_action_at" bson:"last_action_at"`
	}
)

func (p *Player) CanSpend(amount int64) bool {
	return p.Balance >= amount && p.Balance > 0
}

func (p *Player) Spend(amount int64) error {
	if !p.CanSpend(amount) {
		return apperr.ErrInsufficientBalance
	}
	p.Balance -= amount
	return nil
}

func (p *Player) AddReward(amount int64) error {
	if amount < 0 {
		return apperr.ErrInvalidBalance
	}
	p.Balance += amount
	return nil
}

func (p *Player) IsValid() (ok bool, err error) {
	if p.PlayerID == "" {
		return false, apperr.New(apperr.CodeInvalidPlayerID, "player id is required")
	}
	if p.Balance < 0 {
		return false, apperr.ErrInvalidBalance
	}
	return true, nil
}
