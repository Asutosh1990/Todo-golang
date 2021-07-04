package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"

	"weekend_project/Todo-golang/src/models"

	"github.com/gorilla/mux"
)

//toDoItems := make(map[string]models.ToDoItem)
type ToDoService struct {
	ToDoItems map[string]models.ToDoItem
}

func NewToDoService() *ToDoService {

	Items := make(map[string]models.ToDoItem)
	//initialize with composite literal
	return &ToDoService{ToDoItems: Items}
}

func (s ToDoService) DeleteToDoItems(id string) bool {
	f := s.ToDoItems[id]
	deleted := false
	if f.Id == id {
		delete(s.ToDoItems, id)
		deleted = true
	}
	return deleted
}

func (s ToDoService) HandleDeleteToDoItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	deleted := s.DeleteToDoItems(id)

	if deleted {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (s ToDoService) PostToDoItems(description string, dueDate time.Time) *models.ToDoItem {
	tdi := models.NewToDoItem(fmt.Sprintf("%s %d", description, len(s.ToDoItems)+1))
	if !dueDate.IsZero() {
		tdi.DueDate = dueDate
	}
	s.ToDoItems[tdi.Id] = *tdi
	return tdi
}

func (s ToDoService) HandlePostToDoItems(w http.ResponseWriter, r *http.Request) {

	var item models.ToDoItem
	description := "[No description]"
	item.Description = description
	json.NewDecoder(r.Body).Decode(&item)

	tdi := s.PostToDoItems(item.Description, item.DueDate)
	b, _ := marshelFormat(tdi)
	handleResponse(w, b, http.StatusCreated)
}

func (s ToDoService) ListToDoitem() []models.ToDoItem {
	temp := []models.ToDoItem{}
	for _, value := range s.ToDoItems {
		temp = append(temp, value)
	}

	//Apply sorting
	sort.Slice(temp, func(first, second int) bool {
		return temp[first].DueDate.Before(temp[second].DueDate)
	})
	return temp
}

func (s ToDoService) HandleListToDoItems(w http.ResponseWriter, r *http.Request) {

	temp := s.ListToDoitem()

	if len(temp) > 0 {
		b, _ := marshelFormat(temp)
		handleResponse(w, b, http.StatusOK)
	} else {
		handleResponse(w, []byte{}, http.StatusNotFound)
	}

}

func (s ToDoService) GetToDoItem(id string) models.ToDoItem {
	item := s.ToDoItems[id]
	return item
}

func (s ToDoService) HandleGetToDoItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	item := s.GetToDoItem(id)

	if len(item.Id) > 0 {
		b, _ := marshelFormat(item)
		handleResponse(w, b, http.StatusOK)
	} else {
		handleResponse(w, []byte{}, http.StatusNotFound)
	}

}

func (s ToDoService) PatchToDoItems(id string, description string, dueDate time.Time) (*models.ToDoItem, bool) {

	tdi := s.GetToDoItem(id)
	isFound := false

	if len(tdi.Id) > 0 {
		isFound = true
		if !dueDate.IsZero() {
			tdi.DueDate = dueDate
		}
		if len(description) > 0 {
			tdi.Description = description
		}
		s.ToDoItems[tdi.Id] = tdi
	}
	return &tdi, isFound
}

func (s ToDoService) HandlePatchToDoItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	var item models.ToDoItem
	json.NewDecoder(r.Body).Decode(&item)

	tdi, found := s.PatchToDoItems(id, item.Description, item.DueDate)

	if !found {
		handleResponse(w, []byte{}, http.StatusNotFound)
	} else {
		b, _ := marshelFormat(tdi)
		handleResponse(w, b, http.StatusCreated)
	}

}
