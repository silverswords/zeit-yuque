package yuque

import (
	"net/http"

	service "github.com/silverswords/clouds/openapi/yuque"
	util "github.com/silverswords/clouds/pkgs/http"
	con "github.com/silverswords/clouds/pkgs/http/context"
)

// UserRepos API for a list of repositories for a user
func UserRepos(w http.ResponseWriter, r *http.Request) {
	var (
		yuque struct {
			UserID string `json:"user_id" zeit:"required"`
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
	if t == "" {
		c.WriteJSON(http.StatusUnauthorized, con.H{"status": http.StatusUnauthorized})
		return
	}

	s := service.NewService(t)

	resp, err := s.UserRepos(yuque.UserID)
	if err != nil {
		c.WriteJSON(http.StatusRequestTimeout, con.H{"status": http.StatusRequestTimeout})
		return
	}

	c.WriteJSON(http.StatusOK, con.H{"status": http.StatusOK, "user_repos": resp})
}
