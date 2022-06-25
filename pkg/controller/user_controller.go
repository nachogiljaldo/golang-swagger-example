package controller

import (
	"fmt"
	"net/http"
)

func Users(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		GetUsers(w, req)
	} else {
		CreateUser(w, req)
	}
}

// CreateUser godoc
// @Summary      creates an user
// @Description  Allows creating an user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        message  body  model.User  true  "User to create"
// @Success      200   {object} model.User
// @Header       200   string   Location "The path in which the newly created resource can be found"
// @Failure      400   {object} model.Error
// @Failure      404   {object} model.Error
// @Failure      500   {object} model.Error
// @Router       /users [post]
func CreateUser(w http.ResponseWriter, req *http.Request) {

	// This handler does something a little more
	// sophisticated by reading all the HTTP request
	// headers and echoing them into the response body.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// GetUsers godoc
// @Summary      gets the users
// @Description  Retrieves the list of users
// @Tags         users
// @Produce      json
// @Success      200   {array}  model.User
// @Failure      500   {object} model.Error
// @Router       /users [get]
func GetUsers(w http.ResponseWriter, req *http.Request) {

	// This handler does something a little more
	// sophisticated by reading all the HTTP request
	// headers and echoing them into the response body.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// GetUser godoc
// @Summary      Gets a user given its id
// @Description  Retrieves the list of users
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200   {object}  model.User
// @Failure      404   {object} model.Error
// @Failure      500   {object} model.Error
// @Router       /users/{id} [get]
func GetUser(w http.ResponseWriter, req *http.Request) {

	// This handler does something a little more
	// sophisticated by reading all the HTTP request
	// headers and echoing them into the response body.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// UpdateUser godoc
// @Summary      Updates a user given its id
// @Description  Updates a user given its id, performs a full update of the user.
// @Tags         users
// @Produce      json
// @Param        message  body  model.User  true  "New content for the user"
// @Param        id   path      int  true  "User ID"
// @Success      204 "success"
// @Failure      404   {object} model.Error
// @Failure      500   {object} model.Error
// @Router       /users/{id} [put]
func UpdateUser(w http.ResponseWriter, req *http.Request) {

	// This handler does something a little more
	// sophisticated by reading all the HTTP request
	// headers and echoing them into the response body.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// DeleteUser godoc
// @Summary      Deletes a user given its id
// @Description  Deletes a user given its id, performs a full delete of the user. Method is idempotent.
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      204 "success"
// @Failure      500   {object} model.Error
// @Router       /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, req *http.Request) {

	// This handler does something a little more
	// sophisticated by reading all the HTTP request
	// headers and echoing them into the response body.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
