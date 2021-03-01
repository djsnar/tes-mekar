package utils

const (
	// FOR USER
	SELECT_ALL_USER   = "SELECT * FROM penduduk"
	SELECT_USER_BY_ID = "SELECT * FROM penduduk WHERE no_ktp=?"
	INSERT_USER       = "INSERT INTO penduduk VALUES (?, ?, ?, ?, ?)"
	UPDATE_USER       = "UPDATE penduduk SET nama_user=?, tanggal_lahir=?, no_ktp=?, pekerjaan=?, pendidikan=? WHERE no_ktp=?"
	DELETE_USER       = "DELETE FROM penduduk WHERE no_ktp=?"
)
