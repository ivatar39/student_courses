package restapi

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var todo []TODO

func Routes() {
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON)) // принимаем только json формат

	r.Route("/todo", func(r chi.Router) {
		r.Post("/", CreateTodo)
		r.Put("/{id}", UpdateTodo)
		r.Delete("/{id}", DeleteTodo)
	})

	r.Get("/todos", GetTodos)

	http.ListenAndServe(":3333", r)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	data := &TODORequest{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't unmarshal body"))
		return
	}

	todo = append(todo, TODO{
		ID:        uint(len(todo) + 1),
		Name:      data.Name,
		Desc:      data.Desc,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	w.WriteHeader(http.StatusOK)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	data := &TODORequest{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't unmarshal body"))
		return
	}

	idForm := path.Base(r.URL.String())

	id, err := strconv.Atoi(idForm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't parse id"))
		return
	}

	if id > len(todo) || id <= 0 || len(todo) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for i := range todo {
		if todo[i].ID == uint(id) {
			if data.Name != "" {
				todo[id-1].Name = data.Name
			}

			if data.Desc != "" {
				todo[id-1].Desc = data.Desc
			}

			todo[id-1].UpdatedAt = time.Now()
		}
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idForm := path.Base(r.URL.String())

	id, err := strconv.Atoi(idForm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't parse id"))
		return
	}

	if id > len(todo) || id <= 0 || len(todo) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for i, v := range todo {
		if v.ID == uint(id) {
			todo = append(todo[:i], todo[i+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusOK)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("can't marshal todo"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}
