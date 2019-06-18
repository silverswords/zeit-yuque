package yuque

import (
	"net/http"

	service "github.com/silverswords/clouds/openapi/yuque"
	core "github.com/silverswords/clouds/pkgs/http/context"
)

// Repostory -
func Repostory(w http.ResponseWriter, r *http.Request) {
	var (
		yuque struct {
			GroupID string `zeit:"required"`
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
		c.WriteJSON(http.StatusNotAcceptable, core.H{"status": http.StatusNotAcceptable})
		return
	}

	Token := c.Request.Header
	s := service.NewService(Token["X-Auth-Token"][0])
	resp, err := s.Repo(yuque.GroupID)
	if err != nil {
		c.WriteJSON(http.StatusNotAcceptable, core.H{"status": http.StatusNotAcceptable})
		return
	}

	c.WriteJSON(http.StatusOK, core.H{"status": http.StatusOK, "Repo": resp})
}
