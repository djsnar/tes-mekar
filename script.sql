CREATE DATABASE db_populasi;

CREATE TABLE penduduk (
    nama_user VARCHAR(255) NOT NULL,
    tanggal_lahir DATE NOT NULL,
    no_ktp VARCHAR(16) NOT NULL PRIMARY KEY,
    pekerjaan VARCHAR(50) NOT NULL,
    pendidikan VARCHAR(50) NOT NULL
);

INSERT INTO penduduk (nama_user, tanggal_lahir, no_ktp, pekerjaan, pendidikan) VALUES 
('Johan', '1995-09-10', '3520170832480003', 'PNS', 'Sarjana'),
('Fulan', '1998-11-23', '4453284832430001', 'Wirausaha', 'Diploma'),
('Rojul', '1994-02-16', '1221894929900002', 'Wiraswasta', 'SMA');

SELECT * FROM penduduk;