package model

import "time"

type Pasien struct {
	Id            int    `json:"id"`
	Nik           string `json:"nik"`
	Nama          string `json:"nama"`
	Alamat        string `json:"alamat"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	No_hp         string `json:"no_hp"`
	Tempat_lahir  string `json:"tempat_lahir"`
	Tanggal_lahir string `json:"tanggal_lahir"`
}

type Jadwal struct {
	Id             int    `json:"id"`
	Agenda         string `json:"agenda"`
	Nik            string `json:"nik"`
	User_type      string `json:"user_type"`
	Id_ruangan     int    `json:"id_ruangan"`
	Tanggal_masuk  string `json:"tanggal_masuk"`
	Tanggal_keluar string `json:"tanggal_keluar"`
}

type Dokter struct {
	Id_user       int    `json:"id_user"`
	Sip           string `json:"sip"`
	Nama_dokter   string `json:"nama_dokter"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Nik           string `json:"nik"`
	Alamat        string `json:"alamat"`
	Spesialis     string `json:"spesialis"`
}

type Rekam_medis struct {
	Id               int `json:"id"`
	Tanggal          time.Time
	Keluhan          string `json:"keluhan"`
	Pemeriksaan      string `json:"pemeriksaan"`
	Jenis_penanganan string `json:"jenis_penanganan"`
	Kode_obat        string `json:"kode_obat"`
	Id_pasien        int    `json:"id_pasien"`
}

type Obat struct {
	Kode_obat      int    `json:"kode_obat"`
	Nama_obat      string `json:"nama_obat"`
	Jenis_obat     string `json:"jenis_obat"`
	Tahun_produksi string `json:"tahun_produksi"`
	Masa_berlaku   string `json:"masa_berlaku"`
	Total_obat     int    `json:"total_obat"`
}

type Ruangan struct {
	Id            int    `json:"id"`
	Nama_ruangan  string `json:"nama_ruangan"`
	Jenis_ruangan string `json:"jenis_ruangan"`
	Kapasitas     string `json:"kapasitas"`
}

type Perawat struct {
	Id_user       int    `json:"id_user"`
	Sip           string `json:"sip"`
	Nama_perawat  string `json:"nama_perawat"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Nik           string `json:"nik"`
	No_telepon    string `json:"no_telepon"`
}

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Level    string `json:"level"`
}

type Rawat_jalan struct {
	Id                 int    `json:"id"`
	Jadwal_rawat_jalan string `json:"jadwal_rawat_jalan"`
	Nomer_antrian      string `json:"nomer_antrian"`
}
