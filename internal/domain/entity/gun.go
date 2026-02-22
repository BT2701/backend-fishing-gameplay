package entity

import apperr "github.com/BT2701/backend-fishing-gameplay/pkg/error"

type (
	Gun struct {
		GunID      int `json:"gun_id" bson:"gun_id"`
		BulletCost int `json:"bullet_cost" bson:"bullet_cost"`
		Damage     int `json:"damage" bson:"damage"`
		FireRateMs int `json:"fire_rate_ms" bson:"fire_rate_ms"`
	}
)

func (g *Gun) IsValid() (ok bool, err error) {
	if g.GunID <= 0 {
		return false, apperr.New(apperr.Code("INVALID_GUN_ID"), "gun id must be > 0")
	}
	if g.BulletCost < 0 {
		return false, apperr.ErrInvalidBalance
	}
	if g.Damage <= 0 {
		return false, apperr.New(apperr.Code("INVALID_DAMAGE"), "damage must be > 0")
	}
	return true, nil
}

func (g *Gun) CanFireWith(playerBalance int64) bool {
	return playerBalance >= int64(g.BulletCost)
}
