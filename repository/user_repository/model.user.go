package user_repository

import (
	"time"

	"github.com/edwinjordan/SigmatechTest-Golang/entity"
	"github.com/edwinjordan/SigmatechTest-Golang/pkg/helpers"
	"gorm.io/gorm"
)

type User struct {
	UserId       string `gorm:"primaryKey;column:user_id;type:varchar(32)"`
	NIK          string `gorm:"column:nik;type:varchar(200)"`
	FullName     string `gorm:"column:full_name;type:varchar(200)"`
	LegalName    string `gorm:"column:legal_name;type:varchar(200)"`
	TempatLahir  string `gorm:"column:tempat_lahir;type:varchar(200)"`
	TanggalLahir string `gorm:"column:tanggal_lahir;type:date"`
	Gaji         int    `gorm:"column:gaji;type:int"`
	FotoKtp      string `gorm:"column:foto_ktp;type:varchar(200)"`
	FotoKtpPath  string `gorm:"column:foto_ktp_path;type:varchar(200)"`

	FotoSelfie     string `gorm:"column:foto_selfie;type:varchar(200)"`
	FotoSelfiePath string `gorm:"column:foto_selfie_path;type:varchar(200)"`

	UserCreateAt time.Time `gorm:"column:user_create_at;type:datetime"`
	UserUpdateAt time.Time `gorm:"column:user_update_at;type:datetime"`
}

func (User) TableName() string {
	return "tb_user"
}

func (model *User) BeforeCreate(tx *gorm.DB) (err error) {
	model.UserId = helpers.GenUUID()

	model.UserCreateAt = helpers.CreateDateTime()
	model.UserUpdateAt = helpers.CreateDateTime()
	return
}

func (model *User) BeforeUpdate(tx *gorm.DB) (err error) {

	model.UserUpdateAt = helpers.CreateDateTime()
	return
}

func (User) FromEntity(e *entity.User) *User {
	return &User{
		UserId:       e.UserId,
		NIK:          e.NIK,
		FullName:     e.FullName,
		LegalName:    e.LegalName,
		TempatLahir:  e.TempatLahir,
		TanggalLahir: e.TanggalLahir,
		Gaji:         e.Gaji,
		FotoKtp:      e.FotoKtp,
		FotoKtpPath:  e.FotoKtpPath,
		FotoSelfie:   e.FotoSelfie,
		UserCreateAt: e.UserCreateAt,
		UserUpdateAt: e.UserUpdateAt,
	}
}

func (model *User) ToEntity() *entity.UserResponse {
	modelData := &entity.UserResponse{
		UserId:       model.UserId,
		NIK:          model.NIK,
		FullName:     model.FullName,
		LegalName:    model.LegalName,
		TempatLahir:  model.TempatLahir,
		TanggalLahir: model.TanggalLahir,
		Gaji:         model.Gaji,
		FotoKtp:      model.FotoKtp,
		FotoKtpPath:  model.FotoKtpPath,
		FotoSelfie:   model.FotoSelfie,
		UserCreateAt: model.UserCreateAt,
		UserUpdateAt: model.UserUpdateAt,
	}
	return modelData
}
