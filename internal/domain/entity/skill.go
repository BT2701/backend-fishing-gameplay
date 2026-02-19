package entity

type (
	Skill struct {
		SkillType  string `json:"skill_type" bson:"skill_type"`
		Cost       int    `json:"cost" bson:"cost"`
		CooldownMs int    `json:"cooldown_ms" bson:"cooldown_ms"`
	}
)
