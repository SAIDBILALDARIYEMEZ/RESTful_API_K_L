package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Kayıt olma işlemi için "/signup" yolunu dinle

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		// Form verilerini al

		uName := r.FormValue("username")
		email := r.FormValue("email")
		pwd := r.FormValue("password")
		pwdConfirm := r.FormValue("Confirm")

		// Verilerin boş olup olmadığını kontrol et

		if IsEmpty(uName) || IsEmpty(email) || IsEmpty(pwd) || IsEmpty(pwdConfirm) {
			fmt.Fprintf(w, "There is empty data.")
			return
		}

		// Şifre ve şifre onayını kontrol et

		if pwd == pwdConfirm {
			// Veritabanına kullanıcıyı kaydet (bu işlevselliği ekleyin)
			fmt.Fprintf(w, "Registration successful.")
		} else {
			fmt.Fprintln(w, "Password information must be the same.")
		}
	})

	// HTTP sunucusunu 8080 portunda başlat

	http.ListenAndServe(":8080", mux)
}

func IsEmpty(data string) bool { // kendisine gelen datayı boşmu dolumu diye kontrol eder
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}
