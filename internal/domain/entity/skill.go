package entity

import (
	"time"

	apperr "github.com/BT2701/backend-fishing-gameplay/pkg/error"
)

type (
	Skill struct {
		SkillType  string `json:"skill_type" bson:"skill_type"`
		Cost       int    `json:"cost" bson:"cost"`
		CooldownMs int    `json:"cooldown_ms" bson:"cooldown_ms"`
	}

	SkillCooldown struct {
		SkillType  string `json:"skill_type" bson:"skill_type"`
		LastUsedAt int64  `json:"last_used_at" bson:"last_used_at"`
	}
)

func (s *Skill) IsValid() (ok bool, err error) {
	if s.SkillType == "" {
		return false, apperr.New(apperr.Code("INVALID_SKILL_TYPE"), "skill type is required")
	}
	if s.Cost < 0 {
		return false, apperr.ErrInvalidBalance
	}
	if s.CooldownMs < 0 {
		return false, apperr.New(apperr.Code("INVALID_COOLDOWN"), "cooldown must be >= 0")
	}
	return true, nil
}

func (sc *SkillCooldown) IsReady(now time.Time, skill *Skill) bool {
	if sc.LastUsedAt == 0 {
		return true
	}
	elapsed := now.Unix()*1000 - sc.LastUsedAt
	return elapsed >= int64(skill.CooldownMs)
}
