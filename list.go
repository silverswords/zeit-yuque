package yuque

import (
	"net/http"

	service "github.com/silverswords/clouds/openapi/yuque"
	core "github.com/silverswords/clouds/pkgs/http/context"
)

// BookList -
func BookList(w http.ResponseWriter, r *http.Request) {
	var (
		yuque struct {
			RepoID string `zeit:"required"`
		}
	)

	c := core.NewContext(w, r)
	if err := c.BindJSON(&yuque); err != nil {
		c.WriteJSON(http.StatusBadRequest, core.H{"status": http.StatusBadRequest})
		return
	}

	if err := core.Validate(&yuque); err != nil {
		c.WriteJSON(http.StatusNotAcceptable, core.H{"status": http.StatusNotAcceptable})
		return
	}

	Token := c.Request.Header
	s := service.NewService(Token["X-Auth-Token"][0])
	resp, err := s.List(yuque.RepoID)
	if err != nil {
		c.WriteJSON(http.StatusNotAcceptable, core.H{"status": http.StatusNotAcceptable})
		return
	}

	c.WriteJSON(http.StatusOK, core.H{"status": http.StatusOK, "List": resp})
}
