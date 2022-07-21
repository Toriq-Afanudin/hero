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

type Dokter struct {
	Id_user        int    `json:"id_user"`
	Sip            string `json:"sip"`
	Nama_dokter    string `json:"nama_dokter"`
	Poli           string `json:"poli"`
	Jenis_kelamin  string `json:"jenis_kelamin"`
	Jadwal_praktek string `json:"jadwal_praktek"`
	Nomor_str      string `json:"nomor_str"`
	Nomor_telfon   string `json:"nomor_telfon"`
}

type Rekam_medis struct {
	Id               int `json:"id"`
	Tanggal          time.Time
	Keluhan          string `json:"keluhan"`
	Poli             string `json:"poli"`
	Pemeriksaan      string `json:"pemeriksaan"`
	Jenis_penanganan string `json:"jenis_penanganan"`
	Id_pasien        int    `json:"id_pasien"`
}

type Perawat struct {
	Id_user       int    `json:"id_user"`
	Sip           string `json:"sip"`
	Nama_perawat  string `json:"nama_perawat"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Poli          string `json:"poli"`
	Jadwal_kerja  string `json:"jadwal_kerja"`
	Jabatan       string `json:"jabatan"`
	Nomor_telfon  string `json:"nomor_telfon"`
	Nomor_str     string `json:"nomor_str"`
}

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Level    string `json:"level"`
}

type Rawat_jalan struct {
	Id                 int       `json:"id"`
	Jadwal_rawat_jalan string    `json:"jadwal_rawat_jalan"`
	Nomer_antrian      int       `json:"nomer_antrian"`
	Tanggal            time.Time `json:"tanggal"`
	Keterangan         string    `json:"keterangan"`
	Poli               string    `json:"poli"`
	Bool               bool      `json:"bool"`
	Id_pasien          int       `json:"id_pasien"`
}
