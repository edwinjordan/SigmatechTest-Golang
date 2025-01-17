package tenor_repository

import (
	"time"

	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/helpers"
	"gorm.io/gorm"
)

type Tenor struct {
	TenorId       string    `gorm:"column:tenor_id"`
	UserId        string    `gorm:"column:user_id"`
	Tenor         int       `gorm:"column:tenor"`
	TenorMaxLimit int       `gorm:"column:tenor_max_limit"`
	TenorInterest int       `gorm:"column:tenor_interest"`
	TenorCreateAt time.Time `gorm:"column:tenor_create_at"`
	TenorUpdateAt time.Time `gorm:"column:tenor_update_at"`
}

func (Tenor) TableName() string {
	return "tb_tenor"
}

func (model *Tenor) BeforeCreate(tx *gorm.DB) (err error) {
	model.TenorId = helpers.GenUUID()

	model.TenorCreateAt = helpers.CreateDateTime()
	model.TenorUpdateAt = helpers.CreateDateTime()
	return
}

func (model *Tenor) BeforeUpdate(tx *gorm.DB) (err error) {

	model.TenorUpdateAt = helpers.CreateDateTime()
	return
}

func (Tenor) FromEntity(e *entity.Tenor) *Tenor {
	return &Tenor{
		TenorId:       e.TenorId,
		UserId:        e.UserId,
		Tenor:         e.Tenor,
		TenorMaxLimit: e.TenorMaxLimit,
		TenorInterest: e.TenorInterest,
		TenorCreateAt: e.TenorCreateAt,
		TenorUpdateAt: e.TenorUpdateAt,
	}
}

func (model *Tenor) ToEntity() *entity.TenorResponse {
	modelData := &entity.TenorResponse{
		TenorId:       model.TenorId,
		UserId:        model.UserId,
		Tenor:         model.Tenor,
		TenorMaxLimit: model.TenorMaxLimit,
		TenorInterest: model.TenorInterest,
		TenorCreateAt: model.TenorCreateAt,
		TenorUpdateAt: model.TenorUpdateAt,
	}
	return modelData
}
