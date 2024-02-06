package delivery

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/suyono3484/r3d3"
	"net/http"
	"strconv"
)

const (
	listEndpoint   = "/list"
	getEndpoint    = "/spacecraft/:id"
	createEndpoint = "/create"
	updateEndpoint = "/spacecraft/:id/update"
	deleteEndpoint = "/spacecraft/:id/delete"
	name           = "name"
	class          = "class"
	status         = "status"
)

type FilterHelper interface {
	NewNameFilter(string) r3d3.ListFilter
	NewClassFilter(string) r3d3.ListFilter
	NewStatusFilter(string) r3d3.ListFilter
}

type UseCase interface {
	List(filter ...r3d3.ListFilter) ([]r3d3.SpaceCraftInList, error)
	Get(id int64) (r3d3.SpaceCraft, error)
	Create()
	Update()
	Delete()
}

type errorContainer struct {
	Data struct {
		Message string `json:"message"`
	} `json:"data"`
}

type itemArmamentContainer struct {
	Title string `json:"title"`
	Qty   int    `json:"qty"`
}

type itemContainer struct {
	Id       int64                   `json:"id"`
	Name     string                  `json:"name"`
	Class    string                  `json:"class"`
	Crew     uint64                  `json:"crew"`
	Image    string                  `json:"image"`
	Value    float64                 `json:"value"`
	Status   string                  `json:"status"`
	Armament []itemArmamentContainer `json:"armament"`
}

type Delivery struct {
	router *httprouter.Router
	u      UseCase
	helper FilterHelper
}

type listItemResult struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type listResult struct {
	Data []listItemResult `json:"data"`
}

func NewDelivery(u UseCase, helper FilterHelper) *Delivery {
	d := &Delivery{
		router: httprouter.New(),
		u:      u,
		helper: helper,
	}

	d.router.GET(listEndpoint, d.List)
	d.router.GET(getEndpoint, d.Get)
	d.router.POST(createEndpoint, d.Create)
	d.router.POST(updateEndpoint, d.Update)
	d.router.POST(deleteEndpoint, d.Delete)

	return d
}

func (d *Delivery) List(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	qParams := r.URL.Query()
	listFil := make([]r3d3.ListFilter, 0)
	for _, v := range qParams[name] {
		listFil = append(listFil, d.helper.NewNameFilter(v))
	}

	for _, v := range qParams[class] {
		listFil = append(listFil, d.helper.NewClassFilter(v))
	}

	for _, v := range qParams[status] {
		listFil = append(listFil, d.helper.NewStatusFilter(v))
	}

	result, err := d.u.List(listFil...)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output := make([]listItemResult, 0, len(result))
	for _, item := range result {
		output = append(output, listItemResult{
			ID:     item.ID(),
			Name:   item.Name(),
			Status: item.Status(),
		})
	}

	container := listResult{
		Data: output,
	}

	var b []byte
	b, err = json.Marshal(&container)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func (d *Delivery) Get(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	idStr := param.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	var result r3d3.SpaceCraft
	result, err = d.u.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = result
}

func (d *Delivery) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
}

func (d *Delivery) Update(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
}

func (d *Delivery) Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
}
