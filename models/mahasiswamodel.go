package models

type Mahasiswa struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Nim    string `json:"nim"`
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	Prodi  string `json:"prodi"`
	Gambar string `json:"gambar"`
}