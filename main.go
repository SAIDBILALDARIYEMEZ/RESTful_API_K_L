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

	// Giriş yapma işlemi için "/login" yolunu dinle (giriş işlevselliğini ekleyin)

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email := r.FormValue("email")
		pwd := r.FormValue("password")

		if IsEmpty(email) || IsEmpty(pwd) {
			fmt.Fprintf(w, "There is empty data.")
			return
		}

		dbPwd := "1234!*"
		dbEmail := "saidbilal@gmail.com"

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprintf(w, "Login succesful.")
		} else {
			fmt.Fprintf(w, "Login failed.")
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

/*

	for key, value := range r.Form {
		fmt.Printf("%s = %s\n", key, value)
	}
*/
