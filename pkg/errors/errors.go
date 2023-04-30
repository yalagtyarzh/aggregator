package errors

import "errors"

var (
	ErrInvalidProductID   = errors.New("invalid product id in request")
	ErrNoProductID        = errors.New("no product id in request")
	ErrNoProduct          = errors.New("product not found")
	ErrInvalidUserID      = errors.New("invalid user id")
	ErrNoPermissions      = errors.New("no permission to do request")
	ErrUserAlreadyCreated = errors.New("user is already created")
	ErrInvalidRole        = errors.New("invalid role")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrNoUser             = errors.New("no user")
	ErrNoToken            = errors.New("token not found")
	ErrTooManyReviews     = errors.New("too many reviews for user")
	ErrNoReview           = errors.New("review not found")
)
