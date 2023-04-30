package errors

import "errors"

var (
	ErrInvalidProductID   = errors.New("invalid product id in request")
	ErrNoProductID        = errors.New("no product id in request")
	ErrInvalidUserID      = errors.New("invalid user id")
	ErrNoPermissions      = errors.New("no permission for to do request")
	ErrUserAlreadyCreated = errors.New("user is already created")
	ErrInvalidScore       = errors.New("invalid score")
	ErrInvalidRole        = errors.New("invalid role")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrNoUser             = errors.New("no user")
	ErrNoToken            = errors.New("token not found")
)