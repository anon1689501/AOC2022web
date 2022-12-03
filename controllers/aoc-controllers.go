package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Data struct {
	AocDays []AocDay
}

type AocDay struct {
	Day     int
	CookieA string
	CookieB string
	New     bool
}

func MakeIndex(w http.ResponseWriter, r *http.Request) {

	aocInput := ""
	aocDayNum := ""

	data := Data{
		AocDays: []AocDay{
			{Day: 1, CookieA: ReadCookie(r, "day1a"), CookieB: ReadCookie(r, "day1b"), New: true},
			{Day: 2, CookieA: ReadCookie(r, "day2a"), CookieB: ReadCookie(r, "day2b"), New: false},
		},
	}

	if r.Method == "POST" {

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		aocInput = r.FormValue("aocInput")
		aocDayNum = r.FormValue("aocDay")

		aocDayClear := r.FormValue("aocDayClear")

		if aocDayClear != "" {
			//clear cookie
			switch aocDayClear {
			case "1":
				DeleteCookie(w, "day1a")
				DeleteCookie(w, "day1b")
			case "2":
				DeleteCookie(w, "day2a")
				DeleteCookie(w, "day2b")
			default:
				break

			}
			http.Redirect(w, r, "/", http.StatusFound)
		}

		switch aocDayNum {
		case "1":
			Day1(w, aocInput)

		case "2":
			Day2pA(w, aocInput)
			Day2pB(w, aocInput)
		default:
			break

		}
		http.Redirect(w, r, "/", http.StatusFound)
	}

	if r.Method == "DELETE" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		aocDayClear := r.FormValue("aocDayClear")
		log.Print(aocDayClear)
		switch aocDayClear {
		case "1":
			DeleteCookie(w, "day1a")
			DeleteCookie(w, "day1b")
		case "2":
			DeleteCookie(w, "day2a")
			DeleteCookie(w, "day2b")
		default:
			break

		}

		http.Redirect(w, r, "/", http.StatusFound)

	}

	indexTemplate := template.Must(template.ParseFiles("static/index.html"))

	indexTemplate.Execute(w, data)
}
func WriteCookie(w http.ResponseWriter, name string, value string) {

	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   2630000,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
}

func DeleteCookie(w http.ResponseWriter, name string) {
	cookie := http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	//log.Println("cookie deleted")
}
func ReadCookie(r *http.Request, name string) (value string) {
	cookie, _ := r.Cookie(name)
	if cookie == nil {
		value = ""
	} else {
		value = cookie.Value
	}
	return value
}
