package psqlerrors

import (
	"fmt"
	"strings"

	"github.com/ashishkhuraishy/blogge/src/utils/errors/resterror"
	"github.com/lib/pq"
)

const (
	duplicateKey = "23505"
	notFound     = "no rows in result set"
)

// ParseErrors will take in an input error and will
// try to convert the error to a psql error and find
// out what the error is and returns a rest error
// based on what the error is about
func ParseErrors(err error) *resterror.RestError {
	psqlErr, ok := err.(*pq.Error)

	if !ok {
		if strings.Contains(err.Error(), notFound) {
			return resterror.NewNotFoundError("could not find the data")
		}
		return resterror.NewInternalServerError(fmt.Sprintf("database error -> %s", err))
	}

	switch psqlErr.Code {
	case duplicateKey:
		return resterror.NewBadRequest("data already in use")
	default:
		return resterror.NewInternalServerError(fmt.Sprintf("database error \n Code : %s Message : %s", psqlErr.Code, psqlErr.Message))
	}

}
