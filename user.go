package yuque

import (
	"net/http"

	service "github.com/silverswords/clouds/openapi/yuque"
	con "github.com/silverswords/clouds/pkgs/http/context"
)

// UserInfo API for infomation of the authorized user
func UserInfo(w http.ResponseWriter, r *http.Request) {
	c := con.NewContext(w, r)

	token := c.Request.Header
	t := token.Get("X-Auth-Token")
	if t == "" {
		c.WriteJSON(http.StatusUnauthorized, con.H{"status": http.StatusUnauthorized})
		return
	}

	s := service.NewService(t)

	resp, err := s.UserInfo()
	if err != nil {
		c.WriteJSON(http.StatusRequestTimeout, con.H{"status": http.StatusRequestTimeout})
		return
	}

	c.WriteJSON(http.StatusOK, con.H{"status": http.StatusOK, "user": resp})
}
