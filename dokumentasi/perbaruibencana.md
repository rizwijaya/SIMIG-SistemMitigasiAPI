# Perbarui Data Bencana (PUT)

Melalui perbarui data bencana, pengguna dapat memperbarui tanggal bencana jika sudah berhasil diatasi. Dalam memperbarui data tersebut menggunakan kode laporan bencana yang sebelumnya dilaporkan.

**URL** : `/bencana`

**Method** : `PUT`

**Otorisasi diperlukan** : Token

**Izin diperlukan** : Tidak ada

**Contoh Request (Curl)** :
```bash
curl -X PUT "http://localhost:8080/bencana?id=2&tgl_selesai=2021-05-01%2018%3A31%3A00" -H "accept: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjE2MDA4NzksInJvbGVzIjoiYXBwbGljYXRpb24iLCJ1c2VybmFtZSI6ImFkbWluIn0.0WASe37iCvxq_AOy9l-8QDHjMt6BeH1vnMKuoNKc4yw"
```

**Parameters** : 
```json
{
  "id": 0,
  "tgl_selesai": "string"
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

Dalam memperbarui bencana menggunakan id bencana, Jika data berhasil diperbarui maka akan mengembalikan row affected 1 jika gagal row affected 0.