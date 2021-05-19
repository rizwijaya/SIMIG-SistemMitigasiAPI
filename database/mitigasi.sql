-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 19 Bulan Mei 2021 pada 04.08
-- Versi server: 10.4.14-MariaDB
-- Versi PHP: 7.2.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `mitigasi`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `bencana`
--

CREATE TABLE `bencana` (
  `id_bencana` int(11) NOT NULL,
  `id_pelapor` int(11) NOT NULL,
  `bencana` varchar(200) NOT NULL,
  `lokasi` varchar(200) NOT NULL,
  `tgl_terjadi` datetime NOT NULL DEFAULT current_timestamp(),
  `tgl_selesai` datetime DEFAULT NULL,
  `status` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `bencana`
--

INSERT INTO `bencana` (`id_bencana`, `id_pelapor`, `bencana`, `lokasi`, `tgl_terjadi`, `tgl_selesai`, `status`) VALUES
(1, 1, 'Banjir', 'Desa Kedungrejo, Kecamatan Pilangkenceng', '2021-05-01 17:49:56', NULL, 1),
(2, 1, 'Tanah Longsor', 'Desa Sebayi, Kecamatan Gemarang', '2021-05-01 17:50:15', '2021-05-01 18:31:00', 0),
(3, 2, 'Banjir', 'Desa Purworejo, Kecamatan Balerejo', '2021-05-01 17:51:02', '2021-05-01 19:10:00', 1),
(4, 4, 'Banjir', 'Desa Telagan, Kecamatan Pilangkenceng', '2021-05-01 18:31:00', '2021-05-03 17:56:37', 1),
(5, 5, 'Tanah Longsor', 'Desa Kebunagung, Kecamatan Mejayan', '2021-05-01 18:33:00', NULL, 0),
(6, 6, 'Tanah Longsor', 'Desa kare, Kecamatan Kare', '2021-05-02 18:33:00', NULL, 1),
(7, 7, 'Banjir', 'Desa Balerejo, Kecamatan Balerejo', '2021-05-06 18:33:00', NULL, 1),
(8, 8, 'Tanah Longsor', 'Desa Kare, Kecamatan Kare', '2021-05-06 18:33:00', '2021-05-07 16:36:44', 1),
(9, 9, 'Tanah Longsor', 'Desa Kare, Kecamatan Kare', '2021-05-06 18:33:00', NULL, 0),
(10, 10, 'Banjir', 'Desa Ngampel, Kecamatan Mejayan', '2021-05-16 16:36:12', '2021-05-17 16:36:44', 1),
(13, 13, 'Tanah Longsor', 'Desa Pandean, Kecamatan Mejayan ', '2021-05-15 16:36:12', NULL, 1),
(14, 14, 'Banjir', 'Desa Telagan, Kecamatan Pilangkenceng', '2021-05-04 09:18:18', NULL, 0),
(15, 15, 'Banjir', 'Desa Telagan, Kecamatan Pilangkenceng', '2021-05-04 09:18:18', NULL, 0);

-- --------------------------------------------------------

--
-- Struktur dari tabel `berita`
--

CREATE TABLE `berita` (
  `id_berita` int(11) NOT NULL,
  `judul` varchar(200) NOT NULL,
  `isi_berita` varchar(5000) NOT NULL,
  `penulis` varchar(200) NOT NULL,
  `tanggal_ditulis` datetime NOT NULL DEFAULT current_timestamp(),
  `status` int(11) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `berita`
--

INSERT INTO `berita` (`id_berita`, `judul`, `isi_berita`, `penulis`, `tanggal_ditulis`, `status`) VALUES
(1, 'Bencana Longsor Juga Terjadi di Madiun', 'Selain banjir menerjang 21 desa di 6 kecamatan Kabupaten Madiun, bencana longsor juga dialami warga. Tanah longsor menimpa empat rumah warga di Desa Segulung, Kecamatan Dagangan.\r\n\"Jadi selain banjir Madiun juga ada tanah longsor. Ada empat rumah yang terdampak longsor di Madiun, Desa Segulung Kecamatan Dagangan,\" kata Kepala Pelaksanaan BPBD Kabupaten Madiun, Zahrowi, saat meninjau lokasi banjir di Desa Kedung Banteng Kecamatan Pilangkenceng, Kamis (15/4/2021).\r\n\r\nLongsor yang menimpa empat rumah warga, kata Zahrowi, terjadi pada Rabu (14/4) sekitar pukul 18.30 WIB. Saat itu hujan deras mengguyur mulai sore pukul 17.00 WIB hingga pukul 01.00 WIB.\r\n\r\n\"Hujan merata mulai sore hingga dini hari dan longsor sekitar pukul 18.30 WIB,\" jelas Zahrowi.\r\n\r\nZahrowi menambah, keempat rumah yang menjadi terdampak longsor Madiun telah dilakukan pembersihan dan tidak ada korban jiwa. \"Sudah dibersihkan dan tidak ada korban jiwa,\" tandasnya.\r\n\r\nIni empat nama pemilik rumah yang terdampak longsor Desa Segulung, Kecamatan Dagangan Madiun:\r\n\r\n1. Rumah Untung di RT 04/01\r\n2. Rumah Slamet di RT 07/03\r\n3. Rumah Kuat di RT 10/04\r\n4. Rumah Sarno di RT 27/10', 'Sugeng Harianto', '2021-04-15 08:16:38', 0),
(3, 'Diguyur Hujan Semalaman, 22 Desa di Kabupaten Madiun Dilanda Banjir', 'Sebanyak 22 desa dari tujuh kecamatan di Kabupaten Madiun, Jawa Timur, dilanda banjir. Banjir menggenangani rumah warga setelah hujan melanda Kabupaten Madiun dari Rabu (14/4/2021) malam hingga Kamis (15/4/2021). Kepala Pelaksana Badan Penanggulangan Bencana Daerah (BPBD) Kabupaten Madiun, Muhammad Zahrowi mengatakan, tujuh kecamatan yang terdampak banjir yakni Saradan, Mejayan, Pilangkenceng, Wonoasri, Balerejo, dan Wungu “Hujan dengan intensitas sedang hingga tinggi terjadi mulai pukul 18.00 tadi malam hingga Kamis (15/4/2021) dini hari. Akibat hujan itu banjir menggenang 22 desa di enam kecamatan di Kabupaten Madiun,” kata Zahrowi, Kamis. Zahrowi memerinci, 22 desa terdampak banjir yakni enam di Kecamatan Saradan, delapan di Kecamatan Pilangkenceng, satu desa di Kecamatan Wonoasri, lima desa di Kecamatan Balerejo, dan satu desa di Kecamatan Wungu. Menurut Zahrowi, banjir menggenangi permukiman warga, ruas jalan, hingga lahan pertanian. Air yang masuk ke permukiman dan ruas jalan berasal dari luapan sungai terdekat. Air menggenangi rumah warga dan ruas jalan desa hingga kedalaman 80 centimeter. Tak hanya itu, hujan deras mengakibatkan longsor di dua desa di Kecamatan Dagangan. “Di Desa Sewulan hujan deras menjadikan tanggul sungai longsor. Sementara di Desa Segulung empat rumah warga tertimpa tanah longsor,” kata Zahrowi. Total kerugian akibat banjir dan longsor masih didata. Sampai saat ini, Zahrowi belum menerima laporan korban jiwa akibat bencana itu. Untuk keamanan warga, beberapa warga yang bermukim di dekat sungai diungsikan ke rumah tetangga terdekat.', ' Muhlis Al Alaw', '2021-05-04 09:18:18', 1),
(6, 'Bencana Longsor Juga Terjadi di Madiun', 'Selain banjir menerjang 21 desa di 6 kecamatan Kabupaten Madiun, bencana longsor juga dialami warga. Tanah longsor menimpa empat rumah warga di Desa Segulung, Kecamatan Dagangan. \"Jadi selain banjir Madiun juga ada tanah longsor. Ada empat rumah yang terdampak longsor di Madiun, Desa Segulung Kecamatan Dagangan,\" kata Kepala Pelaksanaan BPBD Kabupaten Madiun, Zahrowi, saat meninjau lokasi banjir di Desa Kedung Banteng Kecamatan Pilangkenceng, Kamis (15/4/2021).  Longsor yang menimpa empat rumah warga, kata Zahrowi, terjadi pada Rabu (14/4) sekitar pukul 18.30 WIB. Saat itu hujan deras mengguyur mulai sore pukul 17.00 WIB hingga pukul 01.00 WIB.  \"Hujan merata mulai sore hingga dini hari dan longsor sekitar pukul 18.30 WIB,\" jelas Zahrowi.  Zahrowi menambah, keempat rumah yang menjadi terdampak longsor Madiun telah dilakukan pembersihan dan tidak ada korban jiwa. \"Sudah dibersihkan dan tidak ada korban jiwa,\" tandasnya.  Ini empat nama pemilik rumah yang terdampak longsor Desa Segulung, Kecamatan Dagangan Madiun:  1. Rumah Untung di RT 04/01 2. Rumah Slamet di RT 07/03 3. Rumah Kuat di RT 10/04 4. Rumah Sarno di RT 27/10', 'Sugeng Harianto', '2021-04-15 08:16:38', 0),
(7, 'Bencana Longsor Juga Terjadi di Madiun', 'Selain banjir menerjang 21 desa di 6 kecamatan Kabupaten Madiun, bencana longsor juga dialami warga. Tanah longsor menimpa empat rumah warga di Desa Segulung, Kecamatan Dagangan. \"Jadi selain banjir Madiun juga ada tanah longsor. Ada empat rumah yang terdampak longsor di Madiun, Desa Segulung Kecamatan Dagangan,\" kata Kepala Pelaksanaan BPBD Kabupaten Madiun, Zahrowi, saat meninjau lokasi banjir di Desa Kedung Banteng Kecamatan Pilangkenceng, Kamis (15/4/2021).  Longsor yang menimpa empat rumah warga, kata Zahrowi, terjadi pada Rabu (14/4) sekitar pukul 18.30 WIB. Saat itu hujan deras mengguyur mulai sore pukul 17.00 WIB hingga pukul 01.00 WIB.  \"Hujan merata mulai sore hingga dini hari dan longsor sekitar pukul 18.30 WIB,\" jelas Zahrowi.  Zahrowi menambah, keempat rumah yang menjadi terdampak longsor Madiun telah dilakukan pembersihan dan tidak ada korban jiwa. \"Sudah dibersihkan dan tidak ada korban jiwa,\" tandasnya.  Ini empat nama pemilik rumah yang terdampak longsor Desa Segulung, Kecamatan Dagangan Madiun:  1. Rumah Untung di RT 04/01 2. Rumah Slamet di RT 07/03 3. Rumah Kuat di RT 10/04 4. Rumah Sarno di RT 27/10', 'Sugeng Harianto', '2021-04-15 08:16:38', 0);

-- --------------------------------------------------------

--
-- Struktur dari tabel `pelapor`
--

CREATE TABLE `pelapor` (
  `id_pelapor` int(20) NOT NULL,
  `username` varchar(200) NOT NULL,
  `nama_pelapor` varchar(200) NOT NULL,
  `no_telp` varchar(20) NOT NULL,
  `alamat_email` varchar(200) NOT NULL,
  `tgl_laporan` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `pelapor`
--

INSERT INTO `pelapor` (`id_pelapor`, `username`, `nama_pelapor`, `no_telp`, `alamat_email`, `tgl_laporan`) VALUES
(1, 'rizqi', 'Rizqi Wijaya', '085647374832', 'rizwijaya12@gmail.com', '2021-04-27 20:55:03'),
(2, 'rizqi 2', 'Rizwijaya', '085232313213', 'testing@gmail.com', '2021-04-28 18:52:35'),
(3, 'Ujang', 'ujang selamet', '0856784392832', 'ujang@gmail.com', '2021-05-01 17:17:41'),
(4, 'Antoni', 'Antonius Selamet', '0856784356722', 'antonius@gmail.com', '2021-05-01 18:33:07'),
(5, 'Sriyadi', 'Sriyadi Ainun', '0854324354354', 'sriyadi@gmail.com', '2021-05-01 18:34:35'),
(6, 'Susilo', 'Susilowati', '085647328432', 'susilo@testing.com', '2021-05-02 18:41:43'),
(7, 'Risma', 'Risma Pratama', '0854324354321', 'test@gmail.com', '2021-05-07 16:27:30'),
(8, 'Test 2', 'testing', '0854324354321', 'test2@gmail.com', '2021-05-07 16:35:41'),
(9, 'Test 2', 'testing', '0854324354321', 'test2@gmail.com', '2021-05-07 16:36:12'),
(10, 'ayu', 'ayu sulis', '0867342382934', 'ayu@gmail.com', '2021-05-17 21:15:02'),
(13, 'alfina', 'Alfina Adi', '086737428392', 'alfin@gmail.com', '2021-05-17 21:19:16'),
(14, 'riz', 'rizwijaya', '086758348232', 'riz@gmail.com', '2021-05-18 21:04:27'),
(15, 'riz', 'rizwijaya', '086758348232', 'riz@gmail.com', '2021-05-18 21:04:28');

-- --------------------------------------------------------

--
-- Struktur dari tabel `status_bencana`
--

CREATE TABLE `status_bencana` (
  `id_status` int(11) NOT NULL,
  `nama_status` varchar(200) NOT NULL,
  `date_created` datetime NOT NULL DEFAULT current_timestamp(),
  `date_updated` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `status_bencana`
--

INSERT INTO `status_bencana` (`id_status`, `nama_status`, `date_created`, `date_updated`) VALUES
(0, 'Proses Pengajuan', '2021-05-01 18:48:58', '2021-05-01 18:48:58'),
(1, 'Terverifikasi', '2021-05-01 18:48:58', '2021-05-01 18:48:58'),
(2, 'Ditolak', '2021-05-01 18:48:58', '2021-05-01 18:48:58');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id_user` int(11) NOT NULL,
  `nama_user` varchar(200) NOT NULL,
  `email` varchar(200) NOT NULL,
  `username` varchar(256) NOT NULL,
  `password` varchar(256) NOT NULL,
  `tanggal_dibuat` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id_user`, `nama_user`, `email`, `username`, `password`, `tanggal_dibuat`) VALUES
(1, 'Administration', 'admin@admin.com', 'admin', '$2a$10$QwEDKN2Lpkyg/MSCJKhG9uTmqaKdhXzJGxvvtGJ4LrZ1exXw2holi', '2021-05-05 09:15:56'),
(2, 'User ID', 'user@user.com', 'user', '$2a$10$QwEDKN2Lpkyg/MSCJKhG9uTmqaKdhXzJGxvvtGJ4LrZ1exXw2holi', '2021-05-07 16:37:42'),
(3, 'testing user', 'testing@gmail.com', 'testing', '$2a$10$Mf3HpZekIl0781WKBfJnbO9lVKtnmJaGDEGNQuZc3gyXoDwHTjDXS', '2021-05-16 19:22:30');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `bencana`
--
ALTER TABLE `bencana`
  ADD PRIMARY KEY (`id_bencana`);

--
-- Indeks untuk tabel `berita`
--
ALTER TABLE `berita`
  ADD PRIMARY KEY (`id_berita`);

--
-- Indeks untuk tabel `pelapor`
--
ALTER TABLE `pelapor`
  ADD PRIMARY KEY (`id_pelapor`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id_user`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `bencana`
--
ALTER TABLE `bencana`
  MODIFY `id_bencana` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT untuk tabel `berita`
--
ALTER TABLE `berita`
  MODIFY `id_berita` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT untuk tabel `pelapor`
--
ALTER TABLE `pelapor`
  MODIFY `id_pelapor` int(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id_user` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
