package yuque

import (
	"net/http"

	service "github.com/silverswords/clouds/openapi/yuque"
	util "github.com/silverswords/clouds/pkgs/http"
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
	if err := c.ShouldBind(&yuque); err != nil {
		c.WriteJSON(http.StatusNotAcceptable, con.H{"status": http.StatusNotAcceptable})
		return
	}

	if err := util.Validate(&yuque); err != nil {
		c.WriteJSON(http.StatusConflict, con.H{"status": http.StatusConflict})
		return
	}

	Token := c.Request.Header
	s := service.NewService(Token["X-Auth-Token"][0])

	resp, err := s.List(yuque.RepoID)
	if err != nil {
		c.WriteJSON(http.StatusRequestTimeout, con.H{"status": http.StatusRequestTimeout})
		return
	}

	c.WriteJSON(http.StatusOK, con.H{"status": http.StatusOK, "list": resp})
}
