package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"krizz.org/sleepi/pkg/services"
)

func (r *Routes) addLibraryRoutes() error {
	if r == nil {
		return errors.New("routes is null")
	}
	prefix := "/library"
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/":       r.LibraryIndex,
		"/upload": r.LibraryUpload,
	}
	for url, fn := range routes {
		u := fmt.Sprintf("%v%v", prefix, url)
		r.m.HandleFunc(u, fn)
		log.Println(u)
	}
	return nil
}

func (routes *Routes) LibraryIndex(w http.ResponseWriter, r *http.Request) {
	vars := make(map[string]interface{})
	vars["files"] = routes.api.library.Files
	routes.ren.HTML(routes.withoutFrontendCache(w), http.StatusOK, "library/index", vars)
}

func (routes *Routes) LibraryUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		routes.ren.Text(w, 405, "method not allowed")
		return
	}

	files, err := services.ParseFiles(r)
	if err != nil {
		routes.ren.Text(w, 404, err.Error())
		return
	}

	for _, file := range files {
		fp, err := file.Open()
		if err != nil {
			log.Println(err.Error())
			routes.ren.Text(w, 404, err.Error())
			return
		}

		dat, err := ioutil.ReadAll(fp)
		if err != nil {
			log.Println(err.Error())
			routes.ren.Text(w, 404, err.Error())
			return
		}

		err = routes.api.library.AddFile(dat, file.Filename)
		if err != nil {
			log.Println(err.Error())
			routes.ren.Text(w, 404, err.Error())
			return
		}
	}

	// this is where we want to parse the files :D
	log.Println("redirecting")
	http.Redirect(routes.withoutFrontendCache(w), r, "/library", http.StatusPermanentRedirect)
}
