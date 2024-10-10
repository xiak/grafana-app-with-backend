package binding

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/xiak/grafana-app-with-backend/pkg/common/encoding"
	"github.com/xiak/grafana-app-with-backend/pkg/common/encoding/form"
)

func BadRequest(code int, reason, message string) error {
	return fmt.Errorf("error: code = %d reason = %s message = %s", int32(code), reason, message)
}

// BindQuery bind vars parameters to target.
func BindQuery(vars url.Values, target interface{}) error {
	if err := encoding.GetCodec(form.Name).Unmarshal([]byte(vars.Encode()), target); err != nil {
		return BadRequest(400, "CODEC", err.Error())
	}
	return nil
}

// BindForm bind form parameters to target.
func BindForm(req *http.Request, target interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := encoding.GetCodec(form.Name).Unmarshal([]byte(req.Form.Encode()), target); err != nil {
		return BadRequest(400, "CODEC", err.Error())
	}
	return nil
}
