package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"www.github.com/surajjyoti/finalproject/handler"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			handler.Getstatus(w)
		}

	case "POST":
		{
			handler.QParam = r.FormValue("name")
			if handler.QParam == "" {
				handler.Postall(w, r)
			} else {
				handler.Postquerry(w, handler.QParam)
			}
		}

	}
}

func PrintInMinute(mapData map[string]string) {
	for {
		if handler.Siteslist.SitesName != nil {
			handler.CheckStatus()
			for site, status := range mapData {
				fmt.Printf("%s : %s\n", site, status)

			}
			time.Sleep(5 * time.Second)
		}
	}
}

func main() {
	fmt.Println("Listening at port 3000....")
	go PrintInMinute(handler.Webstatus)
	http.HandleFunc("/websites", requestHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))

}
