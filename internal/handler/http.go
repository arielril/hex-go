package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arielril/hexgo/internal/container/user"
	uController "github.com/arielril/hexgo/internal/controller/user"
)

type HttpHandler interface {
	Serve() error
}

type httpHandler struct {
	UserController uController.UserController
}

func NewHttpServer(ctx *HandlerContext) HttpHandler {
	return &httpHandler{
		UserController: ctx.UserController,
	}
}

func (h *httpHandler) Serve() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!! :P")
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var user user.User

			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Failed to parse body. ERROR: %v", err)
				return
			}

			err = h.UserController.CreateUser(&user)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Failed to insert user")
				return
			}

			resp, _ := json.Marshal(&user)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(resp)
		}
	})

	return http.ListenAndServe(":3000", nil)
}
