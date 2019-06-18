package yuque

import (
	"net/http"

	service "github.com/silverswords/clouds/openapi/yuque"
	core "github.com/silverswords/clouds/pkgs/http"
	con "github.com/silverswords/clouds/pkgs/http/context"
)

// BookList -
func BookList(w http.ResponseWriter, r *http.Request) {
	var (
		yuque struct {
			RepoID string `zeit:"required"`
		}
	)

	c := con.NewContext(w, r)
	if err := c.BindJSON(&yuque); err != nil {
		c.WriteJSON(http.StatusBadRequest, con.H{"status": http.StatusBadRequest})
		return
	}

	if err := core.Validate(&yuque); err != nil {
		c.WriteJSON(http.StatusNotAcceptable, con.H{"status": http.StatusNotAcceptable})
		return
	}

	Token := c.Request.Header
	s := service.NewService(Token["X-Auth-Token"][0])
	resp, err := s.List(yuque.RepoID)
	if err != nil {
		c.WriteJSON(http.StatusNotAcceptable, con.H{"status": http.StatusNotAcceptable})
		return
	}

	c.WriteJSON(http.StatusOK, con.H{"status": http.StatusOK, "List": resp})
}
