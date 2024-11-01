package models

import (
	"time"

	"github.com/google/uuid"
)

type Details struct {
	ID                    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
	SwitchID              uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"                  json:"switchId"`
	Version               *int      `gorm:"type:int"                                        json:"version,omitempty"`
	SpringTypeID          *int      `gorm:"type:int"                                        json:"springTypeId,omitempty"`
	SpringForce           *float32  `gorm:"type:real"                                       json:"springForce,omitempty"`
	SpringMaterialTypeID  *int      `gorm:"type:int"                                        json:"springMaterialTypeId,omitempty"`
	SpringMaterialType    *Type     `gorm:"foreignKey:SpringMaterialTypeID"                 json:"springMaterialType,omitempty"`
	TopHousingMaterialID  *int      `gorm:"type:int"                                        json:"topHousingMaterialId,omitempty"`
	TopHousingMaterial    *Type     `gorm:"foreignKey:TopHousingMaterialID"                 json:"topHousingMaterial,omitempty"`
	BaseHousingMaterialID *int      `gorm:"type:int"                                        json:"baseHousingMaterialId,omitempty"`
	BaseHousingMaterial   *Type     `gorm:"foreignKey:BaseHousingMaterialID"                json:"baseHousingMaterial,omitempty"`
	StemMaterialID        *int      `gorm:"type:int"                                        json:"stemMaterialId,omitempty"`
	StemMaterial          *Type     `gorm:"foreignKey:StemMaterialID"                       json:"stemMaterial,omitempty"`
	HasShineThrough       *bool     `gorm:"type:boolean"                                    json:"hasShineThrough,omitempty"`
	PreTravel             *float32  `gorm:"type:real"                                       json:"preTravel,omitempty"`
	TotalTravel           *float32  `gorm:"type:real"                                       json:"totalTravel,omitempty"`
	InitialForce          *float32  `gorm:"type:real"                                       json:"initialForce,omitempty"`
	ActuationPoint        *float32  `gorm:"type:real"                                       json:"actuationPoint,omitempty"`
	ActuationForce        *float32  `gorm:"type:real"                                       json:"actuationForce,omitempty"`
	ResetPoint            *float32  `gorm:"type:real"                                       json:"resetPoint,omitempty"`
	BottomOutForcePoint   *float32  `gorm:"type:real"                                       json:"bottomOutForcePoint,omitempty"`
	BottomOutForce        *float32  `gorm:"type:real"                                       json:"bottomOutForce,omitempty"`
	SoundLevelID          *int      `gorm:"type:int"                                        json:"soundLevelId,omitempty"`
	SoundLevel            *Type     `gorm:"foreignKey:SoundLevelID"                         json:"soundLevel,omitempty"`
	SoundTypeID           *int      `gorm:"type:int"                                        json:"soundTypeId,omitempty"`
	SoundType             *Type     `gorm:"foreignKey:SoundTypeID"                          json:"soundType,omitempty"`
	TactilityTypeID       *int      `gorm:"type:int"                                        json:"tactilityTypeId,omitempty"`
	TactilityType         *Type     `gorm:"foreignKey:TactilityTypeID"                      json:"tactilityType,omitempty"`
	BumpPosition          *float32  `gorm:"type:real"                                       json:"bumpPosition,omitempty"`
	BumpForce             *float32  `gorm:"type:real"                                       json:"bumpForce,omitempty"`
	TactilityFeedbackID   *int      `gorm:"type:int"                                        json:"tactilityFeedbackId,omitempty"`
	TactilityFeedback     *Type     `gorm:"foreignKey:TactilityFeedbackID"                  json:"tactilityFeedback,omitempty"`
	FactoryLubed          *bool     `gorm:"type:boolean"                                    json:"factoryLubed"`
	StemColorID           *int      `gorm:"type:int"                                        json:"stemColorId,omitempty"`
	StemColor             *Type     `gorm:"foreignKey:StemColorID"                          json:"stemColor,omitempty"`
	TopHousingColorID     *int      `gorm:"type:int"                                        json:"topHousingColorId,omitempty"`
	TopHousingColor       *Type     `gorm:"foreignKey:TopHousingColorID"                    json:"topHousingColor,omitempty"`
	BottomHousingColorID  *int      `gorm:"type:int"                                        json:"bottomHousingColorId,omitempty"`
	BottomHousingColor    *Type     `gorm:"foreignKey:BottomHousingColorID"                 json:"bottomHousingColor,omitempty"`
	PinConfigurationID    *int      `gorm:"type:int"                                        json:"pinConfigurationId,omitempty"`
	PinConfiguration      *Type     `gorm:"foreignKey:PinConfigurationID"                   json:"pinConfiguration,omitempty"`
	CreatedAt             time.Time `gorm:"type:timestamp;autoCreateTime"                   json:"createdAt"`
	UpdatedAt             time.Time `gorm:"type:timestamp;autoUpdateTime"                   json:"updatedAt"`
}
