package apierrors

import "errors"

var ErrMissingSymbol = errors.New("missing symbol or symbols")
var ErrMissingSide = errors.New("missing side")
var ErrMissingType = errors.New("missing type")
var ErrEitherOrderIdOrOrigClientOrderId = errors.New("either orderId or origClientOrderId is required")
