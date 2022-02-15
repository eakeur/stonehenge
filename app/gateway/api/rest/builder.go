package rest

import (
	"net/http"

	"github.com/rs/zerolog"
)

type ResponseBuilder struct {
	Logger zerolog.Logger
}

func (b ResponseBuilder) BuildResult(status int, content interface{}, err error) Response {
	return Response{
		HTTPStatus: status,
		Error:      err,
		Content:    content,
		logger:     b.Logger,
	}
}

func (b ResponseBuilder) BuildCreatedResult(content interface{}) Response {
	return b.BuildResult(http.StatusCreated, content, nil)
}

func (b ResponseBuilder) BuildOKResult(content interface{}) Response {
	return b.BuildResult(http.StatusOK, content, nil)
}

func (b ResponseBuilder) BuildNoContentResult() Response {
	return b.BuildResult(http.StatusNoContent, nil, nil)
}

func (b ResponseBuilder) BuildNotFoundResult(err error) Response {
	return b.BuildResult(http.StatusNotFound, nil, err)
}

func (b ResponseBuilder) BuildBadRequestResult(err error) Response {
	return b.BuildResult(http.StatusBadRequest, nil, err)
}

func (b ResponseBuilder) BuildForbiddenResult(err error) Response {
	return b.BuildResult(http.StatusForbidden, nil, err)
}

func (b ResponseBuilder) BuildUnauthorizedResult(err error) Response {
	return b.BuildResult(http.StatusUnauthorized, nil, err)
}

func (b ResponseBuilder) BuildInternalErrorResult(err error) Response {
	return b.BuildResult(http.StatusInternalServerError, nil, err)
}

func (b ResponseBuilder) BuildErrorResult(err error) Response {
	e := FindMatchingDomainError(err)
	var res Response
	switch e.Status {
	case http.StatusBadRequest:
		res = b.BuildBadRequestResult(e)
	case http.StatusUnauthorized:
		res = b.BuildUnauthorizedResult(e)
	case http.StatusForbidden:
		res = b.BuildForbiddenResult(e)
	case http.StatusNotFound:
		res = b.BuildNotFoundResult(e)
	default:
		res = b.BuildInternalErrorResult(e)
	}
	res.err = err
	return res
}
