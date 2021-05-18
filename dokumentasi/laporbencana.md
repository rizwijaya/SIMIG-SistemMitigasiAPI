# Lapor Bencana (POST)

Lapor Bencana, pengguna dapat melaporkan bencana yang sedang terjadi dilokasinya.

**URL** : `/bencana`

**Method** : `POST`

**Otorisasi diperlukan** : Token

**Izin diperlukan** : Tidak ada

**Contoh Request (Curl)** :
```bash
curl -X POST "http://localhost:8080/bencana?username=riz&nama=rizwijaya&telp=086758348232&email=riz%40gmail.com&bencana=Banjir&lokasi=Desa%20Telagan%2C%20Kecamatan%20Pilangkenceng&tgl_terjadi=2021-05-04%2009%3A18%3A18" -H "accept: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjE2MDA4NzksInJvbGVzIjoiYXBwbGljYXRpb24iLCJ1c2VybmFtZSI6ImFkbWluIn0.0WASe37iCvxq_AOy9l-8QDHjMt6BeH1vnMKuoNKc4yw"
```

**Parameters** : 
```json
{
  "username": "string",
  "nama": "string",
  "telp": "string",
  "email": "string",
  "bencana": "string",
  "lokasi": "string",
  "tgl_terjadi": "string",
}
```

## Success Response

**Kondisi** : Jika data berhasil untuk ditambahkan.

**Kode** : `200 OK`

**Contoh Konten**

```json
{
  "Kode Laporan Bencana": 0
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

Apabila laporan bencana berhasil diterima maka akan mendapatkan kode laporan.