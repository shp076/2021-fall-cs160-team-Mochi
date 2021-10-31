// Code generated by go-swagger; DO NOT EDIT.

package friends_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
)

// AddFriendsV1OKCode is the HTTP code returned for type AddFriendsV1OK
const AddFriendsV1OKCode int = 200

/*AddFriendsV1OK Success

swagger:response addFriendsV1OK
*/
type AddFriendsV1OK struct {

	/*
	  In: Body
	*/
	Payload *models.FriendResponse `json:"body,omitempty"`
}

// NewAddFriendsV1OK creates AddFriendsV1OK with default headers values
func NewAddFriendsV1OK() *AddFriendsV1OK {

	return &AddFriendsV1OK{}
}

// WithPayload adds the payload to the add friends v1 o k response
func (o *AddFriendsV1OK) WithPayload(payload *models.FriendResponse) *AddFriendsV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add friends v1 o k response
func (o *AddFriendsV1OK) SetPayload(payload *models.FriendResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddFriendsV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddFriendsV1BadRequestCode is the HTTP code returned for type AddFriendsV1BadRequest
const AddFriendsV1BadRequestCode int = 400

/*AddFriendsV1BadRequest Bad Request

swagger:response addFriendsV1BadRequest
*/
type AddFriendsV1BadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewAddFriendsV1BadRequest creates AddFriendsV1BadRequest with default headers values
func NewAddFriendsV1BadRequest() *AddFriendsV1BadRequest {

	return &AddFriendsV1BadRequest{}
}

// WithPayload adds the payload to the add friends v1 bad request response
func (o *AddFriendsV1BadRequest) WithPayload(payload *models.ErrResponse) *AddFriendsV1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add friends v1 bad request response
func (o *AddFriendsV1BadRequest) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddFriendsV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddFriendsV1UnauthorizedCode is the HTTP code returned for type AddFriendsV1Unauthorized
const AddFriendsV1UnauthorizedCode int = 401

/*AddFriendsV1Unauthorized Unauthorized

swagger:response addFriendsV1Unauthorized
*/
type AddFriendsV1Unauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewAddFriendsV1Unauthorized creates AddFriendsV1Unauthorized with default headers values
func NewAddFriendsV1Unauthorized() *AddFriendsV1Unauthorized {

	return &AddFriendsV1Unauthorized{}
}

// WithPayload adds the payload to the add friends v1 unauthorized response
func (o *AddFriendsV1Unauthorized) WithPayload(payload *models.ErrResponse) *AddFriendsV1Unauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add friends v1 unauthorized response
func (o *AddFriendsV1Unauthorized) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddFriendsV1Unauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddFriendsV1ForbiddenCode is the HTTP code returned for type AddFriendsV1Forbidden
const AddFriendsV1ForbiddenCode int = 403

/*AddFriendsV1Forbidden Forbidden

swagger:response addFriendsV1Forbidden
*/
type AddFriendsV1Forbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewAddFriendsV1Forbidden creates AddFriendsV1Forbidden with default headers values
func NewAddFriendsV1Forbidden() *AddFriendsV1Forbidden {

	return &AddFriendsV1Forbidden{}
}

// WithPayload adds the payload to the add friends v1 forbidden response
func (o *AddFriendsV1Forbidden) WithPayload(payload *models.ErrResponse) *AddFriendsV1Forbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add friends v1 forbidden response
func (o *AddFriendsV1Forbidden) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddFriendsV1Forbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddFriendsV1NotFoundCode is the HTTP code returned for type AddFriendsV1NotFound
const AddFriendsV1NotFoundCode int = 404

/*AddFriendsV1NotFound Not Found

swagger:response addFriendsV1NotFound
*/
type AddFriendsV1NotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewAddFriendsV1NotFound creates AddFriendsV1NotFound with default headers values
func NewAddFriendsV1NotFound() *AddFriendsV1NotFound {

	return &AddFriendsV1NotFound{}
}

// WithPayload adds the payload to the add friends v1 not found response
func (o *AddFriendsV1NotFound) WithPayload(payload *models.ErrResponse) *AddFriendsV1NotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add friends v1 not found response
func (o *AddFriendsV1NotFound) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddFriendsV1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddFriendsV1ConflictCode is the HTTP code returned for type AddFriendsV1Conflict
const AddFriendsV1ConflictCode int = 409

/*AddFriendsV1Conflict Conflict

swagger:response addFriendsV1Conflict
*/
type AddFriendsV1Conflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewAddFriendsV1Conflict creates AddFriendsV1Conflict with default headers values
func NewAddFriendsV1Conflict() *AddFriendsV1Conflict {

	return &AddFriendsV1Conflict{}
}

// WithPayload adds the payload to the add friends v1 conflict response
func (o *AddFriendsV1Conflict) WithPayload(payload *models.ErrResponse) *AddFriendsV1Conflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add friends v1 conflict response
func (o *AddFriendsV1Conflict) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddFriendsV1Conflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddFriendsV1InternalServerErrorCode is the HTTP code returned for type AddFriendsV1InternalServerError
const AddFriendsV1InternalServerErrorCode int = 500

/*AddFriendsV1InternalServerError Internal Server Error

swagger:response addFriendsV1InternalServerError
*/
type AddFriendsV1InternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewAddFriendsV1InternalServerError creates AddFriendsV1InternalServerError with default headers values
func NewAddFriendsV1InternalServerError() *AddFriendsV1InternalServerError {

	return &AddFriendsV1InternalServerError{}
}

// WithPayload adds the payload to the add friends v1 internal server error response
func (o *AddFriendsV1InternalServerError) WithPayload(payload *models.ErrResponse) *AddFriendsV1InternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add friends v1 internal server error response
func (o *AddFriendsV1InternalServerError) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddFriendsV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
