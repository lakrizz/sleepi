package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
	"krizz.org/sleepi/pkg/services"
)

func (r *Routes) addLibraryRoutes() error {
	if r == nil {
		return errors.New("routes is null")
	}
	prefix := "/library"
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/":       r.LibraryIndex,
		"/add":    r.LibraryAdd,
		"/upload": r.LibraryUpload,
	}
	for url, fn := range routes {
		u := fmt.Sprintf("%v%v", prefix, url)
		r.m.HandleFunc(u, fn)
	}
	return nil
}

func (routes *Routes) LibraryIndex(w http.ResponseWriter, r *http.Request) {
	vars := make(map[string]interface{})
	vars["files"] = routes.api.library.Files
	// pp.Println(vars)
	routes.ren.HTML(routes.withoutFrontendCache(w), http.StatusOK, "library/index", vars)
}

func (routes *Routes) LibraryAdd(w http.ResponseWriter, r *http.Request) {
	routes.ren.HTML(w, http.StatusOK, "library/add", nil)
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
		log.Println(file.Filename)
		fp, err := file.Open()
		if err != nil {
			routes.ren.Text(w, 404, err.Error())
			return
		}

		dat, err := ioutil.ReadAll(fp)
		if err != nil {
			routes.ren.Text(w, 404, err.Error())
			return
		}
		pp.Println(dat)

		err = routes.api.library.AddFile(dat, file.Filename)
		if err != nil {
			routes.ren.Text(w, 404, err.Error())
			return
		}
	}

	pp.Println(routes.api.library.Files)

	// this is where we want to parse the files :D
	routes.ren.HTML(w, 200, "library/upload", nil)
}
