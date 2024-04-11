package routes

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gosimple/slug"

	"github.com/lakrizz/sleepi/pkg/util"
)

func (r *Routes) AddLibraryRoutes() error {
	if r == nil {
		return errors.New("routes is null")
	}
	prefix := "/library"
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/":            r.LibraryIndex,
		"/upload":      r.LibraryUpload,
		"/remove/{id}": r.LibraryRemove,
	}
	for url, fn := range routes {
		u := fmt.Sprintf("%v%v", prefix, url)
		r.m.HandleFunc(u, fn)
		log.Println(u)
	}
	return nil
}

func (routes *Routes) LibraryIndex(w http.ResponseWriter, r *http.Request) {
	vars := map[string]any{
		"Files": routes.rt.Library.GetFiles(),
	}
	routes.ren.HTML(routes.withoutFrontendCache(w), http.StatusOK, "library/main", vars)
}

func (routes *Routes) LibraryRemove(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		routes.ren.Text(w, 404, err.Error())
		return
	}
	err = routes.rt.Library.RemoveByID(id)
	if err != nil {
		routes.ren.Text(w, 404, err.Error())
		return
	}

	playlists, err := routes.rt.Playlists.GetAllPlaylists()
	if err != nil {
		return
	}

	// remove from all playlists
	for _, v := range playlists {
		if v.ContainsFile(id) {
			v.RemoveFile(id)
		}
	}

	routes.ren.HTML(routes.withoutFrontendCache(w), http.StatusPermanentRedirect, "util/redirect", "/library/")
}

func (routes *Routes) LibraryUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		routes.ren.Text(w, 405, "method not allowed")
		return
	}

	files, err := util.ParseHTTPMultiFormFiles(r)
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
		sluggedFileName := slug.Make(strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename)))
		err = routes.rt.Library.AddFile(dat, fmt.Sprintf("%v%v", sluggedFileName, filepath.Ext(file.Filename)))
		if err != nil {
			log.Println(err.Error())
			routes.ren.Text(w, 404, err.Error())
			return
		}
	}

	// this is where we want to parse the files :D
	routes.ren.HTML(routes.withoutFrontendCache(w), http.StatusPermanentRedirect, "util/redirect", "/library/")
}
