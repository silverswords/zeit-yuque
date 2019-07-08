package yuque

import (
	"net/http"

	service "github.com/silverswords/clouds/openapi/yuque"
	util "github.com/silverswords/clouds/pkgs/http"
	con "github.com/silverswords/clouds/pkgs/http/context"
)

// BookDetail -
func BookDetail(w http.ResponseWriter, r *http.Request) {
	var (
		yuque struct {
			RepoID string `zeit:"required"`
			ID     string `zeit:"required"`
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

	Token := c.Request.Header
	s := service.NewService(Token["X-Auth-Token"][0])

	resp, err := s.Details(yuque.RepoID, yuque.ID)
	if err != nil {
		c.WriteJSON(http.StatusRequestTimeout, con.H{"status": http.StatusRequestTimeout})
		return
	}

	c.WriteJSON(http.StatusOK, con.H{"status": http.StatusOK, "detail": resp})
}
