// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2022 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package k_m_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// KMSDescribeSelfIdentityOKCode is the HTTP code returned for type KMSDescribeSelfIdentityOK
const KMSDescribeSelfIdentityOKCode int = 200

/*
KMSDescribeSelfIdentityOK A successful response.

swagger:response kMSDescribeSelfIdentityOK
*/
type KMSDescribeSelfIdentityOK struct {

	/*
	  In: Body
	*/
	Payload *models.KmsDescribeSelfIdentityResponse `json:"body,omitempty"`
}

// NewKMSDescribeSelfIdentityOK creates KMSDescribeSelfIdentityOK with default headers values
func NewKMSDescribeSelfIdentityOK() *KMSDescribeSelfIdentityOK {

	return &KMSDescribeSelfIdentityOK{}
}

// WithPayload adds the payload to the k m s describe self identity o k response
func (o *KMSDescribeSelfIdentityOK) WithPayload(payload *models.KmsDescribeSelfIdentityResponse) *KMSDescribeSelfIdentityOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the k m s describe self identity o k response
func (o *KMSDescribeSelfIdentityOK) SetPayload(payload *models.KmsDescribeSelfIdentityResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KMSDescribeSelfIdentityOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
KMSDescribeSelfIdentityDefault Generic error response.

swagger:response kMSDescribeSelfIdentityDefault
*/
type KMSDescribeSelfIdentityDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewKMSDescribeSelfIdentityDefault creates KMSDescribeSelfIdentityDefault with default headers values
func NewKMSDescribeSelfIdentityDefault(code int) *KMSDescribeSelfIdentityDefault {
	if code <= 0 {
		code = 500
	}

	return &KMSDescribeSelfIdentityDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the k m s describe self identity default response
func (o *KMSDescribeSelfIdentityDefault) WithStatusCode(code int) *KMSDescribeSelfIdentityDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the k m s describe self identity default response
func (o *KMSDescribeSelfIdentityDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the k m s describe self identity default response
func (o *KMSDescribeSelfIdentityDefault) WithPayload(payload *models.Error) *KMSDescribeSelfIdentityDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the k m s describe self identity default response
func (o *KMSDescribeSelfIdentityDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KMSDescribeSelfIdentityDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
