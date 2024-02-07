package delivery

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/suyono3484/r3d3"
	"io"
	"net/http"
	"strconv"
)

const (
	listEndpoint      = "/list"
	getEndpoint       = "/spacecraft/:id"
	createEndpoint    = "/create"
	updateEndpoint    = "/spacecraft/:id/update"
	deleteEndpoint    = "/spacecraft/:id/delete"
	name              = "name"
	class             = "class"
	status            = "status"
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json"
)

//go:generate mockgen -package delivery -destination spacecraft_mock.go . UseCase,FilterHelper

type FilterHelper interface {
	NewNameFilter(string) r3d3.ListFilter
	NewClassFilter(string) r3d3.ListFilter
	NewStatusFilter(string) r3d3.ListFilter
}

type UseCase interface {
	List(filter ...r3d3.ListFilter) ([]r3d3.SpaceCraftInList, error)
	Get(id int64) (r3d3.SpaceCraft, error)
	Create(spacecraft r3d3.SpaceCraftCreate, fields []string) error
	Update(id int64, sc r3d3.SpaceCraftCreate, fields []string) error
	Delete(id int64) error
}

type errorContainer struct {
	Data struct {
		Message string `json:"message"`
	} `json:"data"`
}

type spacecraftArmamentOutContainer struct {
	TitleArmament string `json:"title"`
	QtyArmament   int    `json:"qty"`
}

type spacecraftOutContainer struct {
	IdSpaceCraftOut int64 `json:"id"`
	spacecraftAttributes
	ArmamentSpaceCraftOut []spacecraftArmamentOutContainer `json:"armament"`
}

type spacecraftArmamentInContainer struct {
	IdArmament  int64 `json:"id"`
	QtyArmament int   `json:"qty"`
}

type spacecraftInContainer struct {
	spacecraftAttributes
	ArmamentSpaceCraftIn []spacecraftArmamentInContainer `json:"armament"`
}

type spacecraftAttributes struct {
	NameSpaceCraft   string  `json:"name"`
	ClassSpaceCraft  string  `json:"class"`
	CrewSpaceCraft   uint64  `json:"crew"`
	ImageSpaceCraft  string  `json:"image"`
	ValueSpaceCraft  float64 `json:"value"`
	StatusSpaceCraft string  `json:"status"`
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

type actionResponse struct {
	Success bool `json:"success"`
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
	w.Header().Add(contentTypeHeader, jsonContentType)
	w.Write(b)
}

func (d *Delivery) Get(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	idStr := param.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var result r3d3.SpaceCraft
	result, err = d.u.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	armament := make([]spacecraftArmamentOutContainer, 0)
	for _, v := range result.Armament() {
		armament = append(armament, spacecraftArmamentOutContainer{
			TitleArmament: v.Title(),
			QtyArmament:   v.Qty(),
		})
	}

	sc := spacecraftOutContainer{
		IdSpaceCraftOut: result.ID(),
		spacecraftAttributes: spacecraftAttributes{
			NameSpaceCraft:   result.Name(),
			ClassSpaceCraft:  result.Class(),
			CrewSpaceCraft:   result.Crew(),
			ImageSpaceCraft:  result.ImageURL(),
			ValueSpaceCraft:  result.Value(),
			StatusSpaceCraft: result.Status(),
		},
		ArmamentSpaceCraftOut: armament,
	}

	var b []byte
	if b, err = json.Marshal(&sc); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add(contentTypeHeader, jsonContentType)
	w.Write(b)
}

func (d *Delivery) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	aResp := actionResponse{}
	w.Header().Add(contentTypeHeader, jsonContentType)
	jsonRespEnc := json.NewEncoder(w)

	var (
		err  error
		body []byte
	)
	sc := &spacecraftInContainer{}
	fin := make(map[string]any)
	fields := make([]string, 0)

	contentType := r.Header.Get(contentTypeHeader)
	if contentType != "" && contentType != jsonContentType {
		goto createBadRequest
	}

	if body, err = io.ReadAll(r.Body); err != nil {
		goto createBadRequest
	}

	if err = json.Unmarshal(body, sc); err != nil {
		goto createBadRequest
	}

	if err = json.Unmarshal(body, &fin); err != nil {
		goto createBadRequest
	}

	for k, _ := range fin {
		fields = append(fields, k)
	}

	if err = d.u.Create(sc, fields); err != nil {
		goto createBadRequest
	}

	w.WriteHeader(http.StatusOK)
	aResp.Success = true
	jsonRespEnc.Encode(&aResp)
	return

createBadRequest:
	w.WriteHeader(http.StatusBadRequest)
	aResp.Success = false
	jsonRespEnc.Encode(&aResp)
	return
}

func (d *Delivery) Update(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	aResp := actionResponse{}
	w.Header().Add(contentTypeHeader, jsonContentType)
	jsonRespEnc := json.NewEncoder(w)

	var body []byte
	scIn := &spacecraftInContainer{}
	fin := make(map[string]any)
	fields := make([]string, 0)

	idStr := param.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		goto updateBadRequest
	}

	if body, err = io.ReadAll(r.Body); err != nil {
		goto updateBadRequest
	}

	if err = json.Unmarshal(body, scIn); err != nil {
		goto updateBadRequest
	}

	if err = json.Unmarshal(body, &fin); err != nil {
		goto updateBadRequest
	}

	for k, _ := range fin {
		fields = append(fields, k)
	}

	if err = d.u.Update(id, scIn, fields); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		aResp.Success = false
		jsonRespEnc.Encode(&aResp)
		return
	}

	w.WriteHeader(http.StatusOK)
	aResp.Success = true
	jsonRespEnc.Encode(&aResp)
	return

updateBadRequest:
	w.WriteHeader(http.StatusBadRequest)
	aResp.Success = false
	jsonRespEnc.Encode(&aResp)
	return
}

func (d *Delivery) Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	aResp := actionResponse{}
	w.Header().Add(contentTypeHeader, jsonContentType)
	jsonRespEnc := json.NewEncoder(w)

	idStr := param.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		aResp.Success = false
		jsonRespEnc.Encode(&aResp)
		return
	}

	if err = d.u.Delete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		aResp.Success = false
		jsonRespEnc.Encode(&aResp)
		return
	}

	w.WriteHeader(http.StatusOK)
	aResp.Success = true
	jsonRespEnc.Encode(&aResp)
	return
}

func (d *Delivery) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.router.ServeHTTP(w, r)
}

func (si *spacecraftInContainer) Name() string {
	return si.NameSpaceCraft
}

func (si *spacecraftInContainer) Class() string {
	return si.ClassSpaceCraft
}

func (si *spacecraftInContainer) Crew() uint64 {
	return si.CrewSpaceCraft
}

func (si *spacecraftInContainer) ImageURL() string {
	return si.ImageSpaceCraft
}

func (si *spacecraftInContainer) Value() float64 {
	return si.ValueSpaceCraft
}

func (si *spacecraftInContainer) Status() string {
	return si.StatusSpaceCraft
}

func (si *spacecraftInContainer) Armament() []r3d3.ArmamentSpaceCraftCreate {
	result := make([]r3d3.ArmamentSpaceCraftCreate, 0)
	for i := range si.ArmamentSpaceCraftIn {
		result = append(result, &si.ArmamentSpaceCraftIn[i])
	}
	return result
}

func (ai *spacecraftArmamentInContainer) ID() int64 {
	return ai.IdArmament
}

func (ai *spacecraftArmamentInContainer) Qty() int {
	return ai.QtyArmament
}
