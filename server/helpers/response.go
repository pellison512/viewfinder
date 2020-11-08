package helpers

import (
	"fmt"
	"net/http"
)

func WriteErrResponse(w http.ResponseWriter, err error, errStr string, status int) {
	http.Error(w, fmt.Sprintf("%s: :%v", errStr, err), status)
}
