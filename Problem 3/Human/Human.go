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

// this code block is an HTTP request handler for the "POST" method. tt first checks if the HTTP request method is "POST" by
// checking r.Method. if the method is "POST", it processes the request body, which is expected to contain a form value with the key "Nama".
// it then compares the "Nama" value with the Nama values of each Human in the humans slice using a loop.
func HumanData(w http.ResponseWriter, r *http.Request) {
	//If no matching Human is found, the variable count remains 0,
	count := 0

	if r.Method == "POST" {
		GetNameHuman := Human{
			Nama: r.FormValue("Nama"),
		}
		//if a matching Human is found, it constructs a JSON response containing the Nama, NIM, and Alamat fields of the matching Human using the json.Marshal() function.
		//if there are no errors during the marshaling process, it sets the response content type header to "application/json" using w.Header().Set("Content-Type", "application/json"),
		//and writes the response to the http.ResponseWriter using w.Write(response).
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
		//if no matching Human is found, the variable count remains 0, and the code block after the loop checks if the count is still 0. If it is, it means that there was no match found, and it writes a response to the http.ResponseWriter
		//with an HTTP status code of "405 Method Not Allowed" and an error message of "Name doesnt exist".
		if count == 0 {
			http.Error(w, "Name doesnt exist", http.StatusMethodNotAllowed)
		}

	} else {
		http.Error(w, "Method doesnt exist", http.StatusMethodNotAllowed)
	}

}

// the AllHumanData function handles requests to retrieve all the human data in the humans slice. if the request method is GET,
// the AllHumanData function returns all the human data in the humans slice as JSON. If not, an error is returned
func AllHumanData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		humanJson, _ := json.Marshal(humans)

		w.Header().Set("Content-Type", "application/json")
		w.Write(humanJson)
	} else {
		http.Error(w, "Method doesnt exist", http.StatusMethodNotAllowed)
	}
}
