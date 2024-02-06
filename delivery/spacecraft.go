package delivery

import (
	"github.com/julienschmidt/httprouter"
	"github.com/suyono3484/r3d3"
	"net/http"
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

type UseCase interface {
	List(filter ...r3d3.ListFilter) ([]r3d3.SpaceCraftInList, error)
	Get(id int64) (r3d3.SpaceCraft, error)
	Create()
	Update()
	Delete()
}

type Delivery struct {
	router *httprouter.Router
	u      UseCase
}

func NewDelivery(u UseCase) *Delivery {
	d := &Delivery{
		router: httprouter.New(),
		u:      u,
	}

	d.router.GET(listEndpoint, d.List)
	d.router.GET(getEndpoint, d.Get)
	d.router.POST(createEndpoint, d.Create)
	d.router.POST(updateEndpoint, d.Update)
	d.router.POST(deleteEndpoint, d.Delete)

	return d
}

func (d *Delivery) List(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	//qParams := r.URL.Query()
	//var ok bool

}

func (d *Delivery) Get(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
}

func (d *Delivery) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
}

func (d *Delivery) Update(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
}

func (d *Delivery) Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
}
