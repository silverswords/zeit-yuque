package yuque

import (
	"net/http"

	service "github.com/silverswords/clouds/openapi/yuque"
	core "github.com/silverswords/clouds/pkgs/http/context"
)

// BookDetail -
func BookDetail(w http.ResponseWriter, r *http.Request) {
	var (
		yuque struct {
			RepoID string `zeit:"required"`
			ID     string `zeit:"required"`
		}
	)

	c := core.NewContext(w, r)
	err := c.BindJSON(&yuque)
	if err != nil {
		c.WriteJSON(http.StatusBadRequest, core.H{"status": http.StatusBadRequest})
		return
	}

	err = core.Validate(&yuque)
	if err != nil {
		c.WriteJSON(http.StatusRequestTimeout, core.H{"status": http.StatusRequestTimeout})
		return
	}

	Token := c.Request.Header
	s := service.NewService(Token["X-Auth-Token"][0])
	resp, err := s.Details(yuque.RepoID, yuque.ID)
	if err != nil {
		c.WriteJSON(http.StatusNotAcceptable, core.H{"status": err, "token": s.Token})
		return
	}

	c.WriteJSON(http.StatusOK, core.H{"status": http.StatusOK, "GroupRepo": resp})
}
