package errorcode

import "errors"

var (
	ErrInternal              = errors.New("internal error")
	ErrInconsistentIDs       = errors.New("inconsistent IDs")
	ErrAlreadyExists         = errors.New("already exists")
	ErrNotFound              = errors.New("not found")
	ErrTokenNull             = errors.New("null token")
	ErrAccountNull           = errors.New("null primary token number")
	ErrTokenInValid          = errors.New("token invalid")
	ErrUsernamelengthInvalid = errors.New("username length invalid")
	ErrUsernameformatInvalid = errors.New("username format invalid")
	ErrEmailformatInvalid    = errors.New("email format invalid")
	ErrPasslengthInvalid     = errors.New("pass length invalid")
	ErrPassInvalid           = errors.New("invalid password")
	ErrParsing               = errors.New("Error Parsing")
	ErrBadRouting            = errors.New("inconsistent mapping between route and handler")
	ErrInvalidID             = errors.New("Invalid id value")
	ErrUnknown               = errors.New("unknown uri error")
	ErrInvalidArgument       = errors.New("invalid argument")
	ErrUnAuthorized          = errors.New("unauthorized")
	ErrAuth                  = errors.New("Incorrect credentials")
	ErrDuplicateValue        = errors.New("duplicate value")
)
