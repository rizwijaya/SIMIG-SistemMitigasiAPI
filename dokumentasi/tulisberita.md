# Tulis Berita Bencana (POST)

Tulis berita bencana, pengguna dapat menulis informasi maupun berita terkait bencana yang sedang terjadi. Informasi yang ditulis pengguna, akan dilakukan validasi oleh sistem apakah berdasarkan fakta atau tidak.

**URL** : `/berita`

**Method** : `POST`

**Otorisasi diperlukan** : Token

**Izin diperlukan** : Tidak ada

**Contoh Request (Curl)** :
```bash
curl -X POST "http://localhost:8080/berita?judul=Bencana%20Longsor%20Juga%20Terjadi%20di%20Madiun&isi_berita=Selain%20banjir%20menerjang%2021%20desa%20di%206%20kecamatan%20Kabupaten%20Madiun%2C%20bencana%20longsor%20juga%20dialami%20warga.&penulis=Sugeng%20Harianto&tanggal_ditulis=2021-04-15%2008%3A16%3A38" -H "accept: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjE2MDA4NzksInJvbGVzIjoiYXBwbGljYXRpb24iLCJ1c2VybmFtZSI6ImFkbWluIn0.0WASe37iCvxq_AOy9l-8QDHjMt6BeH1vnMKuoNKc4yw"
```

**Parameters** : 
```json
{
  "judul": "string",
  "isi_berita": "string",
  "penulis": "string",
  "tanggal_ditulis": "string"
}
```

## Success Response

**Kondisi** : Jika data berhasil untuk ditambahkan.

**Kode** : `200 OK`

**Contoh Konten**

```json
{
  "Kode Berita": 0
}
```

## Error Responses

**Kondisi** : Jika request dari pengguna gagal atau token yang diinputkan salah.

**Kode** : `400 Bad Request`

**Konten** : 
```json
{
  "message": "missing or malformed jwt"
}
```

### Atau

**Kondisi** : Jika terdapat kesalahan dari input pengguna, sehingga input tidak bisa diproses oleh server.

**Kode** : `422 Unprocessable Entity`

**Konten** : `{}`

## Catatan

Apabila berita berhasil ditulis maka akan mendapatkan kode berita.
