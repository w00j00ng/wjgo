package wjerr

// Starts from 10000 (exception: E_ERROR_OK = 0, E_ERROR_UNDEFINED = 1)
type EError uint

const (
	E_ERROR_OK EError = iota + 0
	E_ERROR_UNDEFINED
)

const (
	_ EError = iota + 11000
	E_ERROR_UNAUTHORIZED
)
