# Data Bencana yang sedang terjadi (GET)

Menampilkan daftar bencana yang sedang terjadi diberbagai wilayah di indonesia.

**URL** : `/bencana`

**Method** : `GET`

**Otorisasi diperlukan** : Tidak ada

**Izin diperlukan** : Tidak ada

**Contoh Request (Curl)** :
```bash
curl -X GET "http://localhost:8080/bencana" -H "accept: application/json"
```

**Parameters** : `{}`

## Success Response

**Kondisi** : Jika data berhasil untuk ditampilkan.

**Kode** : `200 OK`

**Contoh Konten**

```json
{
      "kode_laporan": 1,
      "pelapor": "rizqi",
      "bencana": "Banjir",
      "lokasi": "Desa Purworejo, Kecamatan Pilangkenceng",
      "tanggal_laporan": "2021-04-27 20:55:03",
      "tanggal_terjadi": "2021-05-01 17:49:56",
      "tanggal_selesai": ""
},
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

Jika request berhasil, maka akan menampilkan data daftar bencana yang sedang terjadi.