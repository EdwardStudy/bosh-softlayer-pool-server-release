package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	"strings"

	"github.com/go-openapi/swag"
)

// UpdateVMWithStateURL generates an URL for the update Vm with state operation
type UpdateVMWithStateURL struct {
	Cid int32

	// avoid unkeyed usage
	_ struct{}
}

// Build a url path and query string
func (o *UpdateVMWithStateURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/vms/{cid}"

	cid := swag.FormatInt32(o.Cid)
	if cid != "" {
		_path = strings.Replace(_path, "{cid}", cid, -1)
	} else {
		return nil, errors.New("Cid is required on UpdateVMWithStateURL")
	}
	result.Path = _path

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *UpdateVMWithStateURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *UpdateVMWithStateURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *UpdateVMWithStateURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on UpdateVMWithStateURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on UpdateVMWithStateURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *UpdateVMWithStateURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
