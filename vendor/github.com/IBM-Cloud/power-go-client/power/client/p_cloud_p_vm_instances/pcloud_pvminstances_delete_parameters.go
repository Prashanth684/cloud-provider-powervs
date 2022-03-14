// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewPcloudPvminstancesDeleteParams creates a new PcloudPvminstancesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPvminstancesDeleteParams() *PcloudPvminstancesDeleteParams {
	return &PcloudPvminstancesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesDeleteParamsWithTimeout creates a new PcloudPvminstancesDeleteParams object
// with the ability to set a timeout on a request.
func NewPcloudPvminstancesDeleteParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesDeleteParams {
	return &PcloudPvminstancesDeleteParams{
		timeout: timeout,
	}
}

// NewPcloudPvminstancesDeleteParamsWithContext creates a new PcloudPvminstancesDeleteParams object
// with the ability to set a context for a request.
func NewPcloudPvminstancesDeleteParamsWithContext(ctx context.Context) *PcloudPvminstancesDeleteParams {
	return &PcloudPvminstancesDeleteParams{
		Context: ctx,
	}
}

// NewPcloudPvminstancesDeleteParamsWithHTTPClient creates a new PcloudPvminstancesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPvminstancesDeleteParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesDeleteParams {
	return &PcloudPvminstancesDeleteParams{
		HTTPClient: client,
	}
}

/* PcloudPvminstancesDeleteParams contains all the parameters to send to the API endpoint
   for the pcloud pvminstances delete operation.

   Typically these are written to a http.Request.
*/
type PcloudPvminstancesDeleteParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* DeleteDataVolumes.

	   Indicates if all data volumes attached to the PVMInstance should be deleted when deleting the PVMInstance. Shared data volumes will be deleted if there are no other PVMInstances attached.
	*/
	DeleteDataVolumes *bool

	/* PvmInstanceID.

	   PCloud PVM Instance ID
	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud pvminstances delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesDeleteParams) WithDefaults() *PcloudPvminstancesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud pvminstances delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) WithContext(ctx context.Context) *PcloudPvminstancesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesDeleteParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithDeleteDataVolumes adds the deleteDataVolumes to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) WithDeleteDataVolumes(deleteDataVolumes *bool) *PcloudPvminstancesDeleteParams {
	o.SetDeleteDataVolumes(deleteDataVolumes)
	return o
}

// SetDeleteDataVolumes adds the deleteDataVolumes to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) SetDeleteDataVolumes(deleteDataVolumes *bool) {
	o.DeleteDataVolumes = deleteDataVolumes
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesDeleteParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if o.DeleteDataVolumes != nil {

		// query param delete_data_volumes
		var qrDeleteDataVolumes bool

		if o.DeleteDataVolumes != nil {
			qrDeleteDataVolumes = *o.DeleteDataVolumes
		}
		qDeleteDataVolumes := swag.FormatBool(qrDeleteDataVolumes)
		if qDeleteDataVolumes != "" {

			if err := r.SetQueryParam("delete_data_volumes", qDeleteDataVolumes); err != nil {
				return err
			}
		}
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
