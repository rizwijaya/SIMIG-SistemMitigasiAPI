# Data Pelapor (GET)

Daftar warga yang pernah melaporkan bencana yang terjadi

**URL** : `/pelapor`

**Method** : `GET`

**Otorisasi diperlukan** : Token

**Izin diperlukan** : Tidak ada

**Contoh Request (Curl)** :
```bash
curl -X GET "http://localhost:8080/pelapor" -H "accept: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjE2MDA4NzksInJvbGVzIjoiYXBwbGljYXRpb24iLCJ1c2VybmFtZSI6ImFkbWluIn0.0WASe37iCvxq_AOy9l-8QDHjMt6BeH1vnMKuoNKc4yw"
```

**Parameters** : `{}`

## Success Response

**Kondisi** : Jika data berhasil untuk ditampilkan.

**Kode** : `200 OK`

**Contoh Konten**

```json
{
    "id": 1,
    "username": "rizqi",
    "pelapor": "Rizqi Wijaya"
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

Jika request berhasil, maka akan menampilkan data pelapor.