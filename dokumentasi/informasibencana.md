# Informasi Bencana (GET)

Pada informasi bencana ini akan menampilkan informasi terkait bencana yang sedang terjadi seperti korban jiwa, total kerugian hingga rentang waktu bencana terjadi dan bagaimana penanggulangannya.

**URL** : `/berita`

**Method** : `GET`

**Otorisasi diperlukan** : Tidak ada

**Izin diperlukan** : Tidak ada

**Contoh Request (Curl)** :
```bash
curl -X GET "http://localhost:8080/berita" -H "accept: application/json"
```

**Parameters** : `{}`

## Success Response

**Kondisi** : Jika data berhasil untuk ditampilkan.

**Kode** : `200 OK`

**Contoh Konten**

```json
{
    "kode_berita": 3,
    "judul": "Diguyur Hujan Semalaman, 22 Desa di Kabupaten Madiun Dilanda Banjir",
    "isi_berita": "Sebanyak 22 desa dari tujuh kecamatan di Kabupaten Madiun, Jawa Timur, dilanda banjir. Banjir menggenangani rumah warga setelah hujan melanda Kabupaten Madiun dari Rabu (14/4/2021) malam hingga Kamis (15/4/2021). Kepala Pelaksana Badan Penanggulangan Bencana Daerah (BPBD) Kabupaten Madiun, Muhammad Zahrowi mengatakan, tujuh kecamatan yang terdampak banjir yakni Saradan, Mejayan, Pilangkenceng, Wonoasri, Balerejo, dan Wungu “Hujan dengan intensitas sedang hingga tinggi terjadi mulai pukul 18.00 tadi malam hingga Kamis (15/4/2021) dini hari.",
    "penulis": " Muhlis Al Alaw",
    "tanggal_ditulis": "2021-05-04 09:18:18"
}
```

## Error Responses

**Kondisi** : Jika request dari pengguna gagal.

**Kode** : `400 Bad Request`

**Konten** : `{}`

### Atau

**Kondisi** : Jika terdapat kesalahan dari input pengguna, sehingga input tidak bisa diproses oleh server.

**Kode** : `422 Unprocessable Entity`

**Konten** : `{}`

## Catatan

Jika request berhasil, maka akan menampilkan informasi bencana.