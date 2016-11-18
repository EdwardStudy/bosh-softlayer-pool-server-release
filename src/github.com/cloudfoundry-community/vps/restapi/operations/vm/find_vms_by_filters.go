package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FindVmsByFiltersHandlerFunc turns a function with the right signature into a find vms by filters handler
type FindVmsByFiltersHandlerFunc func(FindVmsByFiltersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindVmsByFiltersHandlerFunc) Handle(params FindVmsByFiltersParams) middleware.Responder {
	return fn(params)
}

// FindVmsByFiltersHandler interface for that can handle valid find vms by filters params
type FindVmsByFiltersHandler interface {
	Handle(FindVmsByFiltersParams) middleware.Responder
}

// NewFindVmsByFilters creates a new http.Handler for the find vms by filters operation
func NewFindVmsByFilters(ctx *middleware.Context, handler FindVmsByFiltersHandler) *FindVmsByFilters {
	return &FindVmsByFilters{Context: ctx, Handler: handler}
}

/*FindVmsByFilters swagger:route POST /vms/findByFilters vm findVmsByFilters

Finds Vms by filters (cpu, memory, private_vlan, public_vlan, state)

*/
type FindVmsByFilters struct {
	Context *middleware.Context
	Handler FindVmsByFiltersHandler
}

func (o *FindVmsByFilters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewFindVmsByFiltersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
