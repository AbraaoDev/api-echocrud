package service

import "errors"

// establishment
var ErrEstablishmentNotFound = errors.New("establishment not found")
var ErrEstablishmentHasStores = errors.New("cannot delete establishment: there are stores associated with this establishment")

// store
var ErrStoreNotFound = errors.New("store not found")

