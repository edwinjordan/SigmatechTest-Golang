package perusahaan_repository

import (
	"time"

	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/helpers"
	"gorm.io/gorm"
)

type Perusahaan struct {
	PerusahaanId       string    `gorm:"column:perusahaan_id"`
	PerusahaanNama     string    `gorm:"column:perusahaan_nama"`
	PerusahaanFee      int       `gorm:"column:perusahaan_fee"`
	PerusahaanAlamat   string    `gorm:"column:perusahaan_alamat"`
	PerusahaanCreateAt time.Time `gorm:"column:perusahaan_create_at"`
	PerusahaanUpdateAt time.Time `gorm:"column:perusahaan_update_at"`
}

func (Perusahaan) TableName() string {
	return "tb_perusahaan"
}

func (model *Perusahaan) BeforeCreate(tx *gorm.DB) (err error) {
	model.PerusahaanId = helpers.GenUUID()

	model.PerusahaanCreateAt = helpers.CreateDateTime()
	model.PerusahaanUpdateAt = helpers.CreateDateTime()
	return
}

func (model *Perusahaan) BeforeUpdate(tx *gorm.DB) (err error) {

	model.PerusahaanUpdateAt = helpers.CreateDateTime()
	return
}

func (Perusahaan) FromEntity(e *entity.Perusahaan) *Perusahaan {
	return &Perusahaan{
		PerusahaanId:       e.PerusahaanId,
		PerusahaanNama:     e.PerusahaanNama,
		PerusahaanFee:      e.PerusahaanFee,
		PerusahaanAlamat:   e.PerusahaanAlamat,
		PerusahaanCreateAt: e.PerusahaanCreateAt,
		PerusahaanUpdateAt: e.PerusahaanUpdateAt,
	}
}

func (model *Perusahaan) ToEntity() *entity.PerusahaanResponse {
	modelData := &entity.PerusahaanResponse{
		PerusahaanId:       model.PerusahaanId,
		PerusahaanNama:     model.PerusahaanNama,
		PerusahaanFee:      model.PerusahaanFee,
		PerusahaanAlamat:   model.PerusahaanAlamat,
		PerusahaanCreateAt: model.PerusahaanCreateAt,
		PerusahaanUpdateAt: model.PerusahaanUpdateAt,
	}
	return modelData
}
