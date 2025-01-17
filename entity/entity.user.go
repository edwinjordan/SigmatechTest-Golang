package entity

import "time"

type User struct {
	UserId         string                 `json:"user_id"`
	NIK            string                 `json:"nik" validate:"required"`
	FullName       string                 `json:"full_name" validate:"required"`
	LegalName      string                 `json:"legal_name" validate:"required"`
	TempatLahir    string                 `json:"tempat_lahir"`
	TanggalLahir   string                 `json:"tanggal_lahir"`
	Gaji           int                    `json:"gaji"`
	FotoKtp        string                 `json:"foto_ktp"`
	FotoKtpData    map[string]interface{} `json:"foto_ktp_data"`
	FotoKtpPath    string                 `json:"foto_ktp_path"`
	FotoSelfie     string                 `json:"foto_selfie"`
	FotoSelfieData map[string]interface{} `json:"foto_selfie_data"`
	FotoSelfiePath string                 `json:"foto_selfie_path"`
	UserCreateAt   time.Time              `json:"-"`
	UserUpdateAt   time.Time              `json:"-"`
}

type UserResponse struct {
	UserId         string                 `json:"user_id"`
	NIK            string                 `json:"nik" validate:"required"`
	FullName       string                 `json:"full_name" validate:"required"`
	LegalName      string                 `json:"legal_name" validate:"required"`
	TempatLahir    string                 `json:"tempat_lahir"`
	TanggalLahir   string                 `json:"tanggal_lahir"`
	Gaji           int                    `json:"gaji"`
	FotoKtp        string                 `json:"foto_ktp"`
	FotoKtpData    map[string]interface{} `json:"foto_ktp_data"`
	FotoKtpPath    string                 `json:"foto_ktp_path"`
	FotoSelfie     string                 `json:"foto_selfie"`
	FotoSelfieData map[string]interface{} `json:"foto_selfie_data"`
	FotoSelfiePath string                 `json:"foto_selfie_path"`
	UserCreateAt   time.Time              `json:"-"`
	UserUpdateAt   time.Time              `json:"-"`
}
