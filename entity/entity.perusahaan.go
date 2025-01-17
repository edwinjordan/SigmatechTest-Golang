package entity

import "time"

type Perusahaan struct {
	PerusahaanId       string    `json:"perusahaan_id"`
	PerusahaanNama     string    `json:"perusahaan_nama"`
	PerusahaanFee      int       `json:"perusahaan_fee"`
	PerusahaanAlamat   string    `json:"perusahaan_alamat"`
	PerusahaanCreateAt time.Time `json:"-"`
	PerusahaanUpdateAt time.Time `json:"-"`
}

type PerusahaanResponse struct {
	PerusahaanId       string    `json:"perusahaan_id"`
	PerusahaanNama     string    `json:"perusahaan_nama"`
	PerusahaanFee      int       `json:"perusahaan_fee"`
	PerusahaanAlamat   string    `json:"perusahaan_alamat"`
	PerusahaanCreateAt time.Time `json:"-"`
	PerusahaanUpdateAt time.Time `json:"-"`
}
