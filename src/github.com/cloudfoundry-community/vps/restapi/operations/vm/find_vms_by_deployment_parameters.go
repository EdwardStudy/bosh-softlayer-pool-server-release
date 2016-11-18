package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindVmsByDeploymentParams creates a new FindVmsByDeploymentParams object
// with the default values initialized.
func NewFindVmsByDeploymentParams() FindVmsByDeploymentParams {
	var ()
	return FindVmsByDeploymentParams{}
}

// FindVmsByDeploymentParams contains all the bound params for the find vms by deployment operation
// typically these are obtained from a http.Request
//
// swagger:parameters findVmsByDeployment
type FindVmsByDeploymentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*deployment values that need to be considered for filter
	  In: query
	  Collection Format: multi
	*/
	Deployment []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *FindVmsByDeploymentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDeployment, qhkDeployment, _ := qs.GetOK("deployment")
	if err := o.bindDeployment(qDeployment, qhkDeployment, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindVmsByDeploymentParams) bindDeployment(rawData []string, hasKey bool, formats strfmt.Registry) error {

	deploymentIC := rawData

	if len(deploymentIC) == 0 {
		return nil
	}

	var deploymentIR []string
	for _, deploymentIV := range deploymentIC {
		deploymentI := deploymentIV

		deploymentIR = append(deploymentIR, deploymentI)
	}

	o.Deployment = deploymentIR

	return nil
}
