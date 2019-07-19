package yuque

import (
	"net/http"

	service "github.com/silverswords/clouds/openapi/yuque"
	util "github.com/silverswords/clouds/pkgs/http"
	con "github.com/silverswords/clouds/pkgs/http/context"
)

// Repostory API for a list of repositories in the group
func Repostory(w http.ResponseWriter, r *http.Request) {
	var (
		yuque struct {
			GroupID string `json:"group_id" zeit:"required"`
		}
	)

	c := con.NewContext(w, r)
	err := c.ShouldBind(&yuque)
	if err != nil {
		c.WriteJSON(http.StatusNotAcceptable, con.H{"status": http.StatusNotAcceptable})
		return
	}

	err = util.Validate(&yuque)
	if err != nil {
		c.WriteJSON(http.StatusConflict, con.H{"status": http.StatusConflict})
		return
	}

	token := c.Request.Header
	t := token.Get("X-Auth-Token")
	s := service.NewService(t)

	resp, err := s.Repo(yuque.GroupID)
	if err != nil {
		c.WriteJSON(http.StatusRequestTimeout, con.H{"status": http.StatusRequestTimeout})
		return
	}

	c.WriteJSON(http.StatusOK, con.H{"status": http.StatusOK, "repo": resp})
}
