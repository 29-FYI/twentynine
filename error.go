package twentynine

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int
	Message string
}

func (err Error) Error() string {
	return fmt.Sprintf("%d %s: %q", err.Code, http.StatusText(err.Code), err.Message)
}
