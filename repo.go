package yuque

import (
	"net/http"

	service "github.com/silverswords/clouds/openapi/yuque"
	core "github.com/silverswords/clouds/pkgs/http"
	con "github.com/silverswords/clouds/pkgs/http/context"
)

// Repostory -
func Repostory(w http.ResponseWriter, r *http.Request) {
	var (
		yuque struct {
			GroupID string `zeit:"required"`
		}
	)

	c := con.NewContext(w, r)
	err := c.BindJSON(&yuque)
	if err != nil {
		c.WriteJSON(http.StatusBadRequest, con.H{"status": http.StatusBadRequest})
		return
	}

	err = core.Validate(&yuque)
	if err != nil {
		c.WriteJSON(http.StatusNotAcceptable, con.H{"status": http.StatusNotAcceptable})
		return
	}

	Token := c.Request.Header
	s := service.NewService(Token["X-Auth-Token"][0])
	resp, err := s.Repo(yuque.GroupID)
	if err != nil {
		c.WriteJSON(http.StatusNotAcceptable, con.H{"status": http.StatusNotAcceptable})
		return
	}

	c.WriteJSON(http.StatusOK, con.H{"status": http.StatusOK, "Repo": resp})
}
