package Human

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Human struct {
	Nama   string
	NIM    string
	Alamat string
}

var humans = []Human{
	{Nama: "Chantikka", NIM: "2502023650", Alamat: "Karanglo"},
	{Nama: "Caca", NIM: "1223456789", Alamat: "Sawojajar"},
}

func HumanData(w http.ResponseWriter, r *http.Request) {

	count := 0

	if r.Method == "POST" {
		GetNameHuman := Human{
			Nama: r.FormValue("Nama"),
		}

		for _, value := range humans {
			if strings.EqualFold(GetNameHuman.Nama, value.Nama) {
				response, err := json.Marshal(Human{
					Nama:   value.Nama,
					NIM:    value.NIM,
					Alamat: value.Alamat,
				})

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				count++

				w.Header().Set("Content-Type", "application/json")
				w.Write(response)
			}
		}

		if count == 0 {
			http.Error(w, "Name doesnt exist", http.StatusMethodNotAllowed)
		}

	} else {
		http.Error(w, "Method doesnt exist", http.StatusMethodNotAllowed)
	}

}

func AllHumanData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		humanJson, _ := json.Marshal(humans)

		w.Header().Set("Content-Type", "application/json")
		w.Write(humanJson)
	} else {
		http.Error(w, "Method doesnt exist", http.StatusMethodNotAllowed)
	}
}
