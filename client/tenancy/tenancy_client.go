// Code generated by go-swagger; DO NOT EDIT.

// SPDX-License-Identifier: Apache-2.0
//

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new tenancy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tenancy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
TenancyChoicesList tenancy choices list API
*/
func (a *Client) TenancyChoicesList(params *TenancyChoicesListParams, authInfo runtime.ClientAuthInfoWriter) (*TenancyChoicesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyChoicesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tenancy__choices_list",
		Method:             "GET",
		PathPattern:        "/tenancy/_choices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenancyChoicesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TenancyChoicesListOK), nil

}

/*
TenancyChoicesRead tenancy choices read API
*/
func (a *Client) TenancyChoicesRead(params *TenancyChoicesReadParams, authInfo runtime.ClientAuthInfoWriter) (*TenancyChoicesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyChoicesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tenancy__choices_read",
		Method:             "GET",
		PathPattern:        "/tenancy/_choices/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenancyChoicesReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TenancyChoicesReadOK), nil

}

/*
TenancyTenantGroupsCreate tenancy tenant groups create API
*/
func (a *Client) TenancyTenantGroupsCreate(params *TenancyTenantGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*TenancyTenantGroupsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantGroupsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tenancy_tenant-groups_create",
		Method:             "POST",
		PathPattern:        "/tenancy/tenant-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenancyTenantGroupsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TenancyTenantGroupsCreateCreated), nil

}

/*
TenancyTenantGroupsDelete tenancy tenant groups delete API
*/
func (a *Client) TenancyTenantGroupsDelete(params *TenancyTenantGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*TenancyTenantGroupsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantGroupsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tenancy_tenant-groups_delete",
		Method:             "DELETE",
		PathPattern:        "/tenancy/tenant-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenancyTenantGroupsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TenancyTenantGroupsDeleteNoContent), nil

}

/*
TenancyTenantGroupsList Call to super to allow for caching
*/
func (a *Client) TenancyTenantGroupsList(params *TenancyTenantGroupsListParams, authInfo runtime.ClientAuthInfoWriter) (*TenancyTenantGroupsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantGroupsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tenancy_tenant-groups_list",
		Method:             "GET",
		PathPattern:        "/tenancy/tenant-groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenancyTenantGroupsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TenancyTenantGroupsListOK), nil

}

/*
TenancyTenantGroupsPartialUpdate tenancy tenant groups partial update API
*/
func (a *Client) TenancyTenantGroupsPartialUpdate(params *TenancyTenantGroupsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*TenancyTenantGroupsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantGroupsPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tenancy_tenant-groups_partial_update",
		Method:             "PATCH",
		PathPattern:        "/tenancy/tenant-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenancyTenantGroupsPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TenancyTenantGroupsPartialUpdateOK), nil

}

/*
TenancyTenantGroupsRead Call to super to allow for caching
*/
func (a *Client) TenancyTenantGroupsRead(params *TenancyTenantGroupsReadParams, authInfo runtime.ClientAuthInfoWriter) (*TenancyTenantGroupsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantGroupsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tenancy_tenant-groups_read",
		Method:             "GET",
		PathPattern:        "/tenancy/tenant-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenancyTenantGroupsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TenancyTenantGroupsReadOK), nil

}

/*
TenancyTenantGroupsUpdate tenancy tenant groups update API
*/
func (a *Client) TenancyTenantGroupsUpdate(params *TenancyTenantGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*TenancyTenantGroupsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantGroupsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tenancy_tenant-groups_update",
		Method:             "PUT",
		PathPattern:        "/tenancy/tenant-groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenancyTenantGroupsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TenancyTenantGroupsUpdateOK), nil

}

/*
TenancyTenantsCreate tenancy tenants create API
*/
func (a *Client) TenancyTenantsCreate(params *TenancyTenantsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*TenancyTenantsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tenancy_tenants_create",
		Method:             "POST",
		PathPattern:        "/tenancy/tenants/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenancyTenantsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TenancyTenantsCreateCreated), nil

}

/*
TenancyTenantsDelete tenancy tenants delete API
*/
func (a *Client) TenancyTenantsDelete(params *TenancyTenantsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*TenancyTenantsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tenancy_tenants_delete",
		Method:             "DELETE",
		PathPattern:        "/tenancy/tenants/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenancyTenantsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TenancyTenantsDeleteNoContent), nil

}

/*
TenancyTenantsList Call to super to allow for caching
*/
func (a *Client) TenancyTenantsList(params *TenancyTenantsListParams, authInfo runtime.ClientAuthInfoWriter) (*TenancyTenantsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tenancy_tenants_list",
		Method:             "GET",
		PathPattern:        "/tenancy/tenants/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenancyTenantsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TenancyTenantsListOK), nil

}

/*
TenancyTenantsPartialUpdate tenancy tenants partial update API
*/
func (a *Client) TenancyTenantsPartialUpdate(params *TenancyTenantsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*TenancyTenantsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantsPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tenancy_tenants_partial_update",
		Method:             "PATCH",
		PathPattern:        "/tenancy/tenants/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenancyTenantsPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TenancyTenantsPartialUpdateOK), nil

}

/*
TenancyTenantsRead Call to super to allow for caching
*/
func (a *Client) TenancyTenantsRead(params *TenancyTenantsReadParams, authInfo runtime.ClientAuthInfoWriter) (*TenancyTenantsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tenancy_tenants_read",
		Method:             "GET",
		PathPattern:        "/tenancy/tenants/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenancyTenantsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TenancyTenantsReadOK), nil

}

/*
TenancyTenantsUpdate tenancy tenants update API
*/
func (a *Client) TenancyTenantsUpdate(params *TenancyTenantsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*TenancyTenantsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTenancyTenantsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tenancy_tenants_update",
		Method:             "PUT",
		PathPattern:        "/tenancy/tenants/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TenancyTenantsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TenancyTenantsUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
