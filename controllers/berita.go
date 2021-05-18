package controllers

import (
	"net/http"
	"project-2-rizwijaya/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary Informasi Bencana
// @Description Pada informasi bencana ini akan menampilkan informasi terkait bencana yang sedang terjadi seperti korban jiwa, total kerugian hingga rentang waktu bencana terjadi dan bagaimana penanggulangannya.
// @Tags Informasi Bencana
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Berita
// @Failure 422 "Unprocessable Entity"
// @Failure 400 "Bad Request"
// @Router /berita [get]
func SemuaBerita(c echo.Context) error { //Menghandle response yang masuk
	result, err := models.SemuaBerita()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()}) //Menampilakan interface error
	}

	return c.JSON(http.StatusOK, result)
}

// @Summary Tulis Berita Bencana
// @Description Tulis berita bencana, pengguna dapat menulis informasi maupun berita terkait bencana yang sedang terjadi.
// @Description Informasi yang ditulis pengguna, akan dilakukan validasi oleh sistem apakah berdasarkan fakta atau tidak.
// @Tags Informasi Bencana
// @Produce  json
// @Security ApiKeyAuth
// @Param judul query string true "Judul"
// @Param isi_berita query string true "Isi_berita"
// @Param penulis query string true "Penulis"
// @Param tanggal_ditulis query string true "Tanggal_ditulis"
// @Success 200 {object} models.Tulis_Berita
// @Failure 422 "Unprocessable Entity"
// @Failure 400 "Bad Request/Token Salah"
// @Router /berita [post]
//Fungsi untuk menulis berita
func TulisBerita(c echo.Context) error {
	judul := c.FormValue("judul")
	isi_berita := c.FormValue("isi_berita")
	penulis := c.FormValue("penulis")
	tanggal_ditulis := c.FormValue("tanggal_ditulis")

	result, err := models.TulisBerita(judul, isi_berita, penulis, tanggal_ditulis)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// @Summary Perbarui Berita Bencana
// @Description Perbarui Berita Bencana, pengguna dapat melakukan pembaruan atau mengedit berita bencana yang sebelumnya ditulis dengan menggunakan kode berita.
// @Description Jika data berhasil diperbarui maka akan mengembalikan row affected 1 jika gagal row affected 0
// @Tags Informasi Bencana
// @Produce json
// @Security ApiKeyAuth
// @Param id query int true "Id"
// @Param judul query string true "Judul"
// @Param isi_berita query string true "Isi_berita"
// @Param penulis query string true "Penulis"
// @Param tanggal_ditulis query string true "Tanggal_ditulis"
// @Success 200 {object} models.Perbarui_Bencana
// @Failure 422 "Unprocessable Entity"
// @Failure 400 "Bad Request/Token Salah"
// @Router /berita [put]
//Fungsi untuk memperbarui berita
func UpdateBerita(c echo.Context) error {
	id := c.FormValue("id")
	judul := c.FormValue("judul")
	isi_berita := c.FormValue("isi_berita")
	penulis := c.FormValue("penulis")
	tanggal_ditulis := c.FormValue("tanggal_ditulis")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateBerita(conv_id, judul, isi_berita, penulis, tanggal_ditulis)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// @Summary Hapus Informasi Bencana
// @Description Hapus Informasi Bencana, pengguna dapat melakukan penghapusan berita yang ditulisnya dengan menggunakan kode berita.
// @Tags Informasi Bencana
// @Produce  json
// @Security ApiKeyAuth
// @Param id query string true "Id"
// @Success 200 {object} models.Perbarui_Bencana
// @Failure 422 "Unprocessable Entity"
// @Failure 400 "Bad Request/Token Salah"
// @Router /berita [delete]
//Fungsi hapus berita
func HapusBerita(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.HapusBerita(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
