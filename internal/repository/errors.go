package repository

import "errors"

// establishment
var ErrDuplicateCorporateNumber = errors.New("já existe um estabelecimento com este CNPJ")

// store
var ErrDuplicateStoreCorporateNumber = errors.New("já existe um loja com este CNPJ")


