package yuque

import (
	"net/http"

	service "github.com/silverswords/clouds/openapi/yuque"
	util "github.com/silverswords/clouds/pkgs/http"
	con "github.com/silverswords/clouds/pkgs/http/context"
)

// BookList API for a list of repositories in the group
func BookList(w http.ResponseWriter, r *http.Request) {
	var (
		yuque struct {
			RepoID string `json:"repo_id" zeit:"required"`
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

	token := c.Request.Header
	t := token.Get("X-Auth-Token")
	s := service.NewService(t)

	resp, err := s.List(yuque.RepoID)
	if err != nil {
		c.WriteJSON(http.StatusRequestTimeout, con.H{"status": http.StatusRequestTimeout})
		return
	}

	c.WriteJSON(http.StatusOK, con.H{"status": http.StatusOK, "list": resp})
}
