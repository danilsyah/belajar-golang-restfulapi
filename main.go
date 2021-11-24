package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ToDo struct {
	Kegiatan string `json:"kegiatan"`
	Waktu    string `json:"waktu"`
}

type JSONResponse struct {
	Code    int    `json:"code"`
	Succes  bool   `json:"success"`
	Message string `json:"message"`
	// Data    []ToDo `json:"data"`
	Data interface{} `json:"data"`
}

func main() {
	daftarKegiatan := []ToDo{}
	daftarKegiatan = append(daftarKegiatan, ToDo{"Belajar Golang Fundamental", "2021-11-10"})
	daftarKegiatan = append(daftarKegiatan, ToDo{"Belajar Golang Restful API", "2021-11-23"})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// GET http://localhost:8080/
		if r.Method == "GET" {
			rw.Header().Add("Content-Type", "application/json")
			// res := JSONResponse{
			// 	http.StatusOK,
			// 	true,
			// 	"Uji Coba GET Method pada Postman",
			// 	[]ToDo{},
			// }
			// resJSON, err := json.Marshal(res)
			// if err != nil {
			// 	http.Error(rw, "Terjadi Kesalahan", http.StatusInternalServerError)
			// }
			// rw.Write(resJSON)

			res := JSONResponse{
				http.StatusOK,
				true,
				"Berhasil mendapatkan daftar aktivitas",
				daftarKegiatan,
			}
			resJSON, err := json.Marshal(res)
			if err != nil {
				fmt.Println("Terjadi Kesalahan")
				http.Error(rw, "Terjadi Kesalahan", http.StatusInternalServerError)
				return
			}
			rw.Write(resJSON)
			return
		} else if r.Method == "POST" { //Menambahkan data baru
			jsonDecode := json.NewDecoder(r.Body)
			aktivitasBaru := ToDo{}
			res := JSONResponse{}

			if err := jsonDecode.Decode(&aktivitasBaru); err != nil {
				fmt.Println("Terjadi Kesalahan")
				http.Error(rw, "Terjadi Kesalahan saat baca input", http.StatusInternalServerError)
				return
			}

			res.Code = http.StatusCreated
			res.Succes = true
			res.Message = "Berhasil menambahkan data"
			res.Data = aktivitasBaru

			daftarKegiatan = append(daftarKegiatan, aktivitasBaru)

			resJSON, err := json.Marshal(res)
			if err != nil {
				fmt.Println("Terjadi Kesalahan")
				http.Error(rw, "Terjadi Kesalahan saat ubah JSON", http.StatusInternalServerError)
				return
			}
			rw.Header().Add("Content-Type", "application/json")
			rw.Write(resJSON)
		}
	})

	fmt.Println("Listening on: 8080 ....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
