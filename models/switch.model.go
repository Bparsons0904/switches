package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Switch struct {
	ID               uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v7();primaryKey"                                         json:"id"`
	Name             string         `gorm:"type:varchar(50);not null;index:idx_name;uniqueIndex:idx_name_manufacturer_type"         json:"name"`
	ShortDescription string         `gorm:"type:varchar(255);not null"                                                              json:"shortDescription"`
	LongDescription  string         `gorm:"type:text;not null"                                                                      json:"longDescription"`
	ManufacturerID   *int           `gorm:"type:int;index;uniqueIndex:idx_name_manufacturer_type"                                   json:"manufacturerId,omitempty"`
	Manufacturer     *Producer      `gorm:"foreignKey:ManufacturerID"                                                               json:"manufacturer,omitempty"`
	BrandID          *int           `gorm:"type:int;index"                                                                          json:"brandId,omitempty"`
	Brand            *Producer      `gorm:"foreignKey:BrandID"                                                                      json:"brand,omitempty"`
	SwitchTypeID     int            `gorm:"type:int;not null;index;uniqueIndex:idx_name_manufacturer_type"                          json:"switchTypeId"`
	SwitchType       *Type          `gorm:"foreignKey:SwitchTypeID"                                                                 json:"switchType,omitempty"`
	ReleaseDate      *time.Time     `gorm:"type:date"                                                                               json:"releaseDate,omitempty"`
	Available        bool           `gorm:"type:boolean;default:true"                                                               json:"available"`
	PricePoint       int            `gorm:"type:int;not null"                                                                       json:"pricePoint"`
	SiteURL          string         `gorm:"type:varchar(255)"                                                                       json:"siteURL,omitempty"`
	CreatedAt        time.Time      `gorm:"autoCreateTime"                                                                          json:"createdAt"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"                                                                          json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index"                                                                                   json:"deletedAt"`
	CreatedByID      uuid.UUID      `gorm:"type:uuid"                                                                               json:"createdById"`
	CreatedBy        User           `gorm:"foreignKey:CreatedByID;references:ID"                                                    json:"createdBy,omitempty"`
	UpdatedByID      uuid.UUID      `gorm:"type:uuid"                                                                               json:"updatedById"`
	UpdatedBy        User           `gorm:"foreignKey:UpdatedByID;references:ID"                                                    json:"updatedBy,omitempty"`
	ImageLinks       []ImageLink    `gorm:"polymorphic:Owner;polymorphicValue:switch;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"imageLinks,omitempty"`
	Ratings          []Rating       `gorm:"foreignKey:SwitchID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"         json:"ratings,omitempty"`
	AverageRating    float64        `gorm:"-"                                                                                       json:"averageRating,omitempty"`
	RatingsCount     int            `gorm:"-"                                                                                       json:"ratingsCount,omitempty"`
	UserRating       *Rating        `gorm:"-"                                                                                       json:"userRating,omitempty"`
}

func (s *Switch) AfterFind(tx *gorm.DB) error {
	type RatingQuery struct {
		Rating float64
		Count  int
	}

	query := `
		SELECT
			ROUND(COALESCE(AVG(ratings.rating), 0), 1) AS rating,
			COUNT(ratings.id) AS count
		FROM ratings
		WHERE switch_id = ?
		 AND admin_review_required = false
	`

	var rating RatingQuery
	if err := tx.Raw(query, s.ID).Scan(&rating).Error; err != nil {
		log.Error().Err(err).Msg("Error getting the ratings from the after find")
		return err
	}

	s.AverageRating = rating.Rating
	s.RatingsCount = rating.Count

	return nil
}

func (s *Switch) GetUserRating(userID uuid.UUID) {
	for _, rating := range s.Ratings {
		if userID == uuid.Nil {
			s.UserRating = &Rating{
				Rating: -1,
			}
		} else {
			if rating.UserID == userID {
				s.UserRating = &rating
				return
			}
		}
	}

	s.UserRating = &Rating{
		Rating: -1,
	}
}
