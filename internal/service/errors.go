package service

import "errors"

var ErrEstablishmentNotFound = errors.New("establishment not found")
var ErrEstablishmentHasStores = errors.New("cannot delete establishment: there are stores associated with this establishment")
