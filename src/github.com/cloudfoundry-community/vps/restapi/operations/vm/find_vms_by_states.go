package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/cloudfoundry-community/vps/models"
	middleware "github.com/go-openapi/runtime/middleware"
)

// FindVmsByStatesHandlerFunc turns a function with the right signature into a find vms by states handler
type FindVmsByStatesHandlerFunc func(FindVmsByStatesParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn FindVmsByStatesHandlerFunc) Handle(params FindVmsByStatesParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// FindVmsByStatesHandler interface for that can handle valid find vms by states params
type FindVmsByStatesHandler interface {
	Handle(FindVmsByStatesParams, *models.User) middleware.Responder
}

// NewFindVmsByStates creates a new http.Handler for the find vms by states operation
func NewFindVmsByStates(ctx *middleware.Context, handler FindVmsByStatesHandler) *FindVmsByStates {
	return &FindVmsByStates{Context: ctx, Handler: handler}
}

/*FindVmsByStates swagger:route GET /vms/findByState vm findVmsByStates

Finds Vms by states

*/
type FindVmsByStates struct {
	Context *middleware.Context
	Handler FindVmsByStatesHandler
}

func (o *FindVmsByStates) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewFindVmsByStatesParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}