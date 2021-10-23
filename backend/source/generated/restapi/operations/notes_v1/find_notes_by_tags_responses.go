// Code generated by go-swagger; DO NOT EDIT.

package notes_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
)

// FindNotesByTagsOKCode is the HTTP code returned for type FindNotesByTagsOK
const FindNotesByTagsOKCode int = 200

/*FindNotesByTagsOK Success

swagger:response findNotesByTagsOK
*/
type FindNotesByTagsOK struct {

	/*
	  In: Body
	*/
	Payload models.NotesGetResponse `json:"body,omitempty"`
}

// NewFindNotesByTagsOK creates FindNotesByTagsOK with default headers values
func NewFindNotesByTagsOK() *FindNotesByTagsOK {

	return &FindNotesByTagsOK{}
}

// WithPayload adds the payload to the find notes by tags o k response
func (o *FindNotesByTagsOK) WithPayload(payload models.NotesGetResponse) *FindNotesByTagsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find notes by tags o k response
func (o *FindNotesByTagsOK) SetPayload(payload models.NotesGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindNotesByTagsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.NotesGetResponse{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// FindNotesByTagsBadRequestCode is the HTTP code returned for type FindNotesByTagsBadRequest
const FindNotesByTagsBadRequestCode int = 400

/*FindNotesByTagsBadRequest Bad Request

swagger:response findNotesByTagsBadRequest
*/
type FindNotesByTagsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewFindNotesByTagsBadRequest creates FindNotesByTagsBadRequest with default headers values
func NewFindNotesByTagsBadRequest() *FindNotesByTagsBadRequest {

	return &FindNotesByTagsBadRequest{}
}

// WithPayload adds the payload to the find notes by tags bad request response
func (o *FindNotesByTagsBadRequest) WithPayload(payload *models.ErrResponse) *FindNotesByTagsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find notes by tags bad request response
func (o *FindNotesByTagsBadRequest) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindNotesByTagsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FindNotesByTagsUnauthorizedCode is the HTTP code returned for type FindNotesByTagsUnauthorized
const FindNotesByTagsUnauthorizedCode int = 401

/*FindNotesByTagsUnauthorized Unauthorized

swagger:response findNotesByTagsUnauthorized
*/
type FindNotesByTagsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewFindNotesByTagsUnauthorized creates FindNotesByTagsUnauthorized with default headers values
func NewFindNotesByTagsUnauthorized() *FindNotesByTagsUnauthorized {

	return &FindNotesByTagsUnauthorized{}
}

// WithPayload adds the payload to the find notes by tags unauthorized response
func (o *FindNotesByTagsUnauthorized) WithPayload(payload *models.ErrResponse) *FindNotesByTagsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find notes by tags unauthorized response
func (o *FindNotesByTagsUnauthorized) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindNotesByTagsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FindNotesByTagsForbiddenCode is the HTTP code returned for type FindNotesByTagsForbidden
const FindNotesByTagsForbiddenCode int = 403

/*FindNotesByTagsForbidden Forbidden

swagger:response findNotesByTagsForbidden
*/
type FindNotesByTagsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewFindNotesByTagsForbidden creates FindNotesByTagsForbidden with default headers values
func NewFindNotesByTagsForbidden() *FindNotesByTagsForbidden {

	return &FindNotesByTagsForbidden{}
}

// WithPayload adds the payload to the find notes by tags forbidden response
func (o *FindNotesByTagsForbidden) WithPayload(payload *models.ErrResponse) *FindNotesByTagsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find notes by tags forbidden response
func (o *FindNotesByTagsForbidden) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindNotesByTagsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FindNotesByTagsNotFoundCode is the HTTP code returned for type FindNotesByTagsNotFound
const FindNotesByTagsNotFoundCode int = 404

/*FindNotesByTagsNotFound Not Found

swagger:response findNotesByTagsNotFound
*/
type FindNotesByTagsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewFindNotesByTagsNotFound creates FindNotesByTagsNotFound with default headers values
func NewFindNotesByTagsNotFound() *FindNotesByTagsNotFound {

	return &FindNotesByTagsNotFound{}
}

// WithPayload adds the payload to the find notes by tags not found response
func (o *FindNotesByTagsNotFound) WithPayload(payload *models.ErrResponse) *FindNotesByTagsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find notes by tags not found response
func (o *FindNotesByTagsNotFound) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindNotesByTagsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FindNotesByTagsConflictCode is the HTTP code returned for type FindNotesByTagsConflict
const FindNotesByTagsConflictCode int = 409

/*FindNotesByTagsConflict Conflict

swagger:response findNotesByTagsConflict
*/
type FindNotesByTagsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewFindNotesByTagsConflict creates FindNotesByTagsConflict with default headers values
func NewFindNotesByTagsConflict() *FindNotesByTagsConflict {

	return &FindNotesByTagsConflict{}
}

// WithPayload adds the payload to the find notes by tags conflict response
func (o *FindNotesByTagsConflict) WithPayload(payload *models.ErrResponse) *FindNotesByTagsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find notes by tags conflict response
func (o *FindNotesByTagsConflict) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindNotesByTagsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FindNotesByTagsInternalServerErrorCode is the HTTP code returned for type FindNotesByTagsInternalServerError
const FindNotesByTagsInternalServerErrorCode int = 500

/*FindNotesByTagsInternalServerError Internal Server Error

swagger:response findNotesByTagsInternalServerError
*/
type FindNotesByTagsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewFindNotesByTagsInternalServerError creates FindNotesByTagsInternalServerError with default headers values
func NewFindNotesByTagsInternalServerError() *FindNotesByTagsInternalServerError {

	return &FindNotesByTagsInternalServerError{}
}

// WithPayload adds the payload to the find notes by tags internal server error response
func (o *FindNotesByTagsInternalServerError) WithPayload(payload *models.ErrResponse) *FindNotesByTagsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find notes by tags internal server error response
func (o *FindNotesByTagsInternalServerError) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindNotesByTagsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
