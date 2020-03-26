// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the inventory client
type API interface {
	/*
	   DeregisterCluster deletes an open shift bare metal cluster definition*/
	DeregisterCluster(ctx context.Context, params *DeregisterClusterParams) (*DeregisterClusterNoContent, error)
	/*
	   DeregisterHost deregisters open shift bare metal host*/
	DeregisterHost(ctx context.Context, params *DeregisterHostParams) (*DeregisterHostNoContent, error)
	/*
	   DisableHost disables a host for use*/
	DisableHost(ctx context.Context, params *DisableHostParams) (*DisableHostNoContent, error)
	/*
	   DownloadClusterISO downloads open shift per cluster discovery i s o*/
	DownloadClusterISO(ctx context.Context, params *DownloadClusterISOParams, writer io.Writer) (*DownloadClusterISOOK, error)
	/*
	   EnableHost enables a host for use*/
	EnableHost(ctx context.Context, params *EnableHostParams) (*EnableHostNoContent, error)
	/*
	   GetCluster retrieves open shift bare metal cluster information*/
	GetCluster(ctx context.Context, params *GetClusterParams) (*GetClusterOK, error)
	/*
	   GetHost retrieves open shift bare metal host information*/
	GetHost(ctx context.Context, params *GetHostParams) (*GetHostOK, error)
	/*
	   GetNextSteps retrieves the next operations that the agent need to perform*/
	GetNextSteps(ctx context.Context, params *GetNextStepsParams) (*GetNextStepsOK, error)
	/*
	   InstallCluster installs a new open shift bare metal cluster*/
	InstallCluster(ctx context.Context, params *InstallClusterParams) (*InstallClusterOK, error)
	/*
	   ListClusters lists open shift bare metal clusters*/
	ListClusters(ctx context.Context, params *ListClustersParams) (*ListClustersOK, error)
	/*
	   ListHosts lists open shift bare metal hosts*/
	ListHosts(ctx context.Context, params *ListHostsParams) (*ListHostsOK, error)
	/*
	   PostStepReply posts the result of the required operations from the server*/
	PostStepReply(ctx context.Context, params *PostStepReplyParams) (*PostStepReplyNoContent, error)
	/*
	   RegisterCluster creates a new open shift bare metal cluster definition*/
	RegisterCluster(ctx context.Context, params *RegisterClusterParams) (*RegisterClusterCreated, error)
	/*
	   RegisterHost registers a new open shift bare metal host*/
	RegisterHost(ctx context.Context, params *RegisterHostParams) (*RegisterHostCreated, error)
	/*
	   SetDebugStep sets a single shot debug step that will be sent next time the agent will ask for a command*/
	SetDebugStep(ctx context.Context, params *SetDebugStepParams) (*SetDebugStepOK, error)
	/*
	   UpdateCluster updates an open shift bare metal cluster definition*/
	UpdateCluster(ctx context.Context, params *UpdateClusterParams) (*UpdateClusterCreated, error)
}

// New creates a new inventory API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for inventory API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeregisterCluster deletes an open shift bare metal cluster definition
*/
func (a *Client) DeregisterCluster(ctx context.Context, params *DeregisterClusterParams) (*DeregisterClusterNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeregisterCluster",
		Method:             "DELETE",
		PathPattern:        "/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeregisterClusterReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeregisterClusterNoContent), nil

}

/*
DeregisterHost deregisters open shift bare metal host
*/
func (a *Client) DeregisterHost(ctx context.Context, params *DeregisterHostParams) (*DeregisterHostNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeregisterHost",
		Method:             "DELETE",
		PathPattern:        "/hosts/{host_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeregisterHostReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeregisterHostNoContent), nil

}

/*
DisableHost disables a host for use
*/
func (a *Client) DisableHost(ctx context.Context, params *DisableHostParams) (*DisableHostNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DisableHost",
		Method:             "DELETE",
		PathPattern:        "/hosts/{host_id}/actions/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DisableHostReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DisableHostNoContent), nil

}

/*
DownloadClusterISO downloads open shift per cluster discovery i s o
*/
func (a *Client) DownloadClusterISO(ctx context.Context, params *DownloadClusterISOParams, writer io.Writer) (*DownloadClusterISOOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DownloadClusterISO",
		Method:             "GET",
		PathPattern:        "/clusters/{cluster_id}/actions/download",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DownloadClusterISOReader{formats: a.formats, writer: writer},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DownloadClusterISOOK), nil

}

/*
EnableHost enables a host for use
*/
func (a *Client) EnableHost(ctx context.Context, params *EnableHostParams) (*EnableHostNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EnableHost",
		Method:             "POST",
		PathPattern:        "/hosts/{host_id}/actions/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EnableHostReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EnableHostNoContent), nil

}

/*
GetCluster retrieves open shift bare metal cluster information
*/
func (a *Client) GetCluster(ctx context.Context, params *GetClusterParams) (*GetClusterOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCluster",
		Method:             "GET",
		PathPattern:        "/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterOK), nil

}

/*
GetHost retrieves open shift bare metal host information
*/
func (a *Client) GetHost(ctx context.Context, params *GetHostParams) (*GetHostOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHost",
		Method:             "GET",
		PathPattern:        "/hosts/{host_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHostReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetHostOK), nil

}

/*
GetNextSteps retrieves the next operations that the agent need to perform
*/
func (a *Client) GetNextSteps(ctx context.Context, params *GetNextStepsParams) (*GetNextStepsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNextSteps",
		Method:             "GET",
		PathPattern:        "/hosts/{host_id}/next-steps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNextStepsReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNextStepsOK), nil

}

/*
InstallCluster installs a new open shift bare metal cluster
*/
func (a *Client) InstallCluster(ctx context.Context, params *InstallClusterParams) (*InstallClusterOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InstallCluster",
		Method:             "POST",
		PathPattern:        "/clusters/{cluster_id}/actions/install",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InstallClusterReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InstallClusterOK), nil

}

/*
ListClusters lists open shift bare metal clusters
*/
func (a *Client) ListClusters(ctx context.Context, params *ListClustersParams) (*ListClustersOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListClusters",
		Method:             "GET",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListClustersReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListClustersOK), nil

}

/*
ListHosts lists open shift bare metal hosts
*/
func (a *Client) ListHosts(ctx context.Context, params *ListHostsParams) (*ListHostsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListHosts",
		Method:             "GET",
		PathPattern:        "/hosts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListHostsReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListHostsOK), nil

}

/*
PostStepReply posts the result of the required operations from the server
*/
func (a *Client) PostStepReply(ctx context.Context, params *PostStepReplyParams) (*PostStepReplyNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostStepReply",
		Method:             "POST",
		PathPattern:        "/hosts/{host_id}/next-steps/reply",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostStepReplyReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostStepReplyNoContent), nil

}

/*
RegisterCluster creates a new open shift bare metal cluster definition
*/
func (a *Client) RegisterCluster(ctx context.Context, params *RegisterClusterParams) (*RegisterClusterCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RegisterCluster",
		Method:             "POST",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegisterClusterReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegisterClusterCreated), nil

}

/*
RegisterHost registers a new open shift bare metal host
*/
func (a *Client) RegisterHost(ctx context.Context, params *RegisterHostParams) (*RegisterHostCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RegisterHost",
		Method:             "POST",
		PathPattern:        "/hosts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegisterHostReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegisterHostCreated), nil

}

/*
SetDebugStep sets a single shot debug step that will be sent next time the agent will ask for a command
*/
func (a *Client) SetDebugStep(ctx context.Context, params *SetDebugStepParams) (*SetDebugStepOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SetDebugStep",
		Method:             "POST",
		PathPattern:        "/hosts/{host_id}/actions/debug",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetDebugStepReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetDebugStepOK), nil

}

/*
UpdateCluster updates an open shift bare metal cluster definition
*/
func (a *Client) UpdateCluster(ctx context.Context, params *UpdateClusterParams) (*UpdateClusterCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateCluster",
		Method:             "PATCH",
		PathPattern:        "/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateClusterReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateClusterCreated), nil

}
