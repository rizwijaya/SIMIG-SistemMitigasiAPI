# Hapus Informasi Bencana (DELETE)

Hapus Informasi Bencana, pengguna dapat melakukan penghapusan berita yang ditulisnya dengan menggunakan kode berita.

**URL** : `/berita`

**Method** : `DELETE`

**Otorisasi diperlukan** : Token

**Izin diperlukan** : Tidak Ada

**Contoh Request (Curl)** :
```bash
curl -X DELETE "http://localhost:8080/berita?id=2" -H "accept: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjE2MDA4NzksInJvbGVzIjoiYXBwbGljYXRpb24iLCJ1c2VybmFtZSI6ImFkbWluIn0.0WASe37iCvxq_AOy9l-8QDHjMt6BeH1vnMKuoNKc4yw"
```

**Parameters** : 
```json
{
  "id": 0
}
```

## Success Response

**Kondisi** : Jika data berhasil untuk dihapus.

**Kode** : `200 OK`

**Contoh Konten**

```json
{
  "rows_Affected": 0
}
```

## Failed Response

**Kondisi** : Jika data gagal untuk dihapus.

**Contoh Konten**

```json
{
  "rows_Affected": 1
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

Apabila berita berhasil untuk dihapus maka rows affected adalah 0.