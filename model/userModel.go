package model

type User struct {
	NamaUser     string `json:"nama_user"`
	TanggalLahir string `json:"tanggal_lahir"`
	NoKTP        string `json:"no_ktp"`
	Pekerjaan    string `json:"pekerjaan"`
	Pendidikan   string `json:"pendidikan"`
}

// nama_user VARCHAR(255) NOT NULL,
// tanggal_lahir DATE NOT NULL,
// no_ktp VARCHAR(16) NOT NULL PRIMARY KEY,
// pekerjaan VARCHAR(50) NOT NULL,
// pendidikan VARCHAR(50) NOT NULL
