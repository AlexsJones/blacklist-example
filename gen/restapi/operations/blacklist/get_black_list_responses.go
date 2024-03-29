// Code generated by go-swagger; DO NOT EDIT.

package blacklist

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/AlexsJones/blacklist-example/gen/models"
)

// GetBlackListOKCode is the HTTP code returned for type GetBlackListOK
const GetBlackListOKCode int = 200

/*GetBlackListOK successful operation

swagger:response getBlackListOK
*/
type GetBlackListOK struct {

	/*
	  In: Body
	*/
	Payload *models.Blacklist `json:"body,omitempty"`
}

// NewGetBlackListOK creates GetBlackListOK with default headers values
func NewGetBlackListOK() *GetBlackListOK {

	return &GetBlackListOK{}
}

// WithPayload adds the payload to the get black list o k response
func (o *GetBlackListOK) WithPayload(payload *models.Blacklist) *GetBlackListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get black list o k response
func (o *GetBlackListOK) SetPayload(payload *models.Blacklist) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBlackListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBlackListBadRequestCode is the HTTP code returned for type GetBlackListBadRequest
const GetBlackListBadRequestCode int = 400

/*GetBlackListBadRequest get black list bad request

swagger:response getBlackListBadRequest
*/
type GetBlackListBadRequest struct {
}

// NewGetBlackListBadRequest creates GetBlackListBadRequest with default headers values
func NewGetBlackListBadRequest() *GetBlackListBadRequest {

	return &GetBlackListBadRequest{}
}

// WriteResponse to the client
func (o *GetBlackListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetBlackListNotFoundCode is the HTTP code returned for type GetBlackListNotFound
const GetBlackListNotFoundCode int = 404

/*GetBlackListNotFound get black list not found

swagger:response getBlackListNotFound
*/
type GetBlackListNotFound struct {
}

// NewGetBlackListNotFound creates GetBlackListNotFound with default headers values
func NewGetBlackListNotFound() *GetBlackListNotFound {

	return &GetBlackListNotFound{}
}

// WriteResponse to the client
func (o *GetBlackListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetBlackListInternalServerErrorCode is the HTTP code returned for type GetBlackListInternalServerError
const GetBlackListInternalServerErrorCode int = 500

/*GetBlackListInternalServerError get black list internal server error

swagger:response getBlackListInternalServerError
*/
type GetBlackListInternalServerError struct {
}

// NewGetBlackListInternalServerError creates GetBlackListInternalServerError with default headers values
func NewGetBlackListInternalServerError() *GetBlackListInternalServerError {

	return &GetBlackListInternalServerError{}
}

// WriteResponse to the client
func (o *GetBlackListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
