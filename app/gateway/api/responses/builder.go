package responses

import "net/http"

type Builder struct {
}

func BuildResult(status int, content interface{}, err error) Response {
	return Response{
		HTTPStatus: status,
		Error:      err,
		Content:    content,
	}
}

func BuildCreatedResult(content interface{}) Response {
	return BuildResult(http.StatusCreated, content, nil)
}

func BuildOKResult(content interface{}) Response {
	return BuildResult(http.StatusOK, content, nil)
}

func BuildNoContentResult() Response {
	return BuildResult(http.StatusNoContent, nil, nil)
}

func BuildNotFoundResult(err error) Response {
	return BuildResult(http.StatusNotFound, nil, err)
}

func BuildBadRequestResult(err error) Response {
	return BuildResult(http.StatusBadRequest, nil, err)
}

func BuildForbiddenResult(err error) Response {
	return BuildResult(http.StatusForbidden, nil, err)
}

func BuildUnauthorizedResult(err error) Response {
	return BuildResult(http.StatusUnauthorized, nil, err)
}

func BuildInternalErrorResult(err error) Response {
	return BuildResult(http.StatusInternalServerError, nil, err)
}

func BuildErrorResult(err error) Response {
	e := FindMatchingDomainError(err)
	switch e.Status {
	case http.StatusBadRequest:
		return BuildBadRequestResult(e)
	case http.StatusUnauthorized:
		return BuildUnauthorizedResult(e)
	case http.StatusForbidden:
		return BuildForbiddenResult(e)
	case http.StatusNotFound:
		return BuildNotFoundResult(e)
	default:
		return BuildInternalErrorResult(e)
	}
}