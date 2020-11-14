package server

import "net/http"

func checkRequest(m string, rw http.ResponseWriter, r *http.Request) bool {
	if r.Method != m {
		http.Error(rw, http.ErrNotSupported.ErrorString, http.StatusForbidden)
		return false
	}
	return true
}

// CheckPOST func
func CheckPOST(rw http.ResponseWriter, r *http.Request) bool {
	return checkRequest("POST", rw, r)
}

// CheckGET func
func CheckGET(rw http.ResponseWriter, r *http.Request) bool {
	return checkRequest("GET", rw, r)
}
