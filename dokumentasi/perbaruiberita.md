# Perbarui Data Bencana (PUT)

Perbarui Berita Bencana, pengguna dapat melakukan pembaruan atau mengedit berita bencana yang sebelumnya ditulis dengan menggunakan kode berita.

**URL** : `/berita`

**Method** : `PUT`

**Otorisasi diperlukan** : Token

**Izin diperlukan** : Tidak ada

**Contoh Request (Curl)** :
```bash
curl -X PUT "http://localhost:8080/berita?id=2&judul=Diguyur%20Hujan%20Semalaman%2C%2022%20Desa%20di%20Kabupaten%20Madiun%20Dilanda%20Banjir&isi_berita=Sebanyak%2022%20desa%20dari%20tujuh%20kecamatan%20di%20Kabupaten%20Madiun%2C%20Jawa%20Timur%2C%20dilanda%20banjir.%20Banjir%20menggenangani%20rumah%20warga%20setelah%20hujan%20melanda%20Kabupaten%20Madiun%20dari%20Rabu%20(14%2F4%2F2021)%20malam%20hingga%20Kamis%20(15%2F4%2F2021).&penulis=Muhlis%20Al%20Alaw&tanggal_ditulis=2021-05-04%2009%3A18%3A18" -H "accept: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjE2MDA4NzksInJvbGVzIjoiYXBwbGljYXRpb24iLCJ1c2VybmFtZSI6ImFkbWluIn0.0WASe37iCvxq_AOy9l-8QDHjMt6BeH1vnMKuoNKc4yw"
```

**Parameters** : 
```json
{
  "id": 0,
  "judul": "string",
  "isi_berita": "string",
  "penulis": "string",
  "tanggal_ditulis": "string"
}
```
## Success Response

**Kondisi** : Jika data berhasil untuk diperbarui.

**Kode** : `200 OK`

**Contoh Konten**

```json
{
  "rows_Affected": 1
}
```

## Failed Response

**Kondisi** : Jika data gagal untuk diperbarui.

**Contoh Konten**

```json
{
  "rows_Affected": 0
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

Saat memperbarui berita menggunakan kode berita. Jika data berhasil diperbarui maka akan mengembalikan row affected 1 jika gagal row affected 0