package vm

import (
	"github.com/cloudfoundry-community/vps/generated/models"
)

func (o *FindVmsByDeploymentOK) GetPayload() *models.VmsResponse{
	return o.Payload
}

func (o *FindVmsByDeploymentNotFound) GetStatusCode() int {
	return 404
}

func (o *FindVmsByDeploymentDefault) GetStatusCode() int {
	return o._statusCode
}

func (o *FindVmsByDeploymentDefault) GetPayload() *models.Error {
	return o.Payload
}
