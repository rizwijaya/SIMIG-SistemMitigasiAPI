# History Bencana (GET)

Pada history bencana ini akan menampilkan bencana yang pernah terjadi sebelumnya diberbagai wilayah di indonesia.

**URL** : `/historibencana`

**Method** : `GET`

**Otorisasi diperlukan** : Token

**Izin diperlukan** : Tidak ada

**Contoh Request (Curl)** :
```bash
curl -X GET "http://localhost:8080/historibencana" -H "accept: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjE2MDA4NzksInJvbGVzIjoiYXBwbGljYXRpb24iLCJ1c2VybmFtZSI6ImFkbWluIn0.0WASe37iCvxq_AOy9l-8QDHjMt6BeH1vnMKuoNKc4yw"
```

**Parameters** : `{}`

## Success Response

**Kondisi** : Jika data berhasil untuk ditampilkan.

**Kode** : `200 OK`

**Contoh Konten**

```json
{
    "kode_laporan": 3,
    "pelapor": "rizqi 2",
    "bencana": "Banjir",
    "lokasi": "Desa Purworejo, Kecamatan Pilangkenceng",
    "tanggal_laporan": "2021-04-28 18:52:35",
    "tanggal_terjadi": "2021-05-01 17:51:02",
    "tanggal_selesai": "2021-05-01 19:10:00"
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

Jika request berhasil, maka akan menampilkan data histori bencana.