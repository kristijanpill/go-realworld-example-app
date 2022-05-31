package model

import "github.com/google/uuid"

type Follow struct {
	ProfileID      uuid.UUID `gorm:"primaryKey; type:uuid;"`
	Profile        Profile
	TargetID  uuid.UUID `gorm:"primaryKey; type:uuid;"`
	Target    Profile
}