{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

{{- range .MethodSets}}
const Operation{{$svrType}}{{.OriginalName}} = "/{{$svrName}}/{{.OriginalName}}"
{{- end}}

type {{.ServiceType}}HTTPServer interface {
{{- range .MethodSets}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	{{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

func Register{{.ServiceType}}HTTPServer(s *http.Server, srv {{.ServiceType}}HTTPServer) {
	r := s.Route("/")
	{{- range .Methods}}
	r.{{.Method}}("{{.Path}}", _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler(srv))
	{{- end}}
}

{{range .Methods}}
func _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler(srv {{$svrType}}HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in {{.Request}}
		{{- if .HasBody}}
		if err := ctx.Bind(&in{{.Body}}); err != nil {
			return err
		}
		{{- end}}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		{{- if .HasVars}}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		{{- end}}
		http.SetOperation(ctx,Operation{{$svrType}}{{.OriginalName}})
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.{{.Name}}(ctx, req.(*{{.Request}}))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*{{.Reply}})
		return ctx.Result(200, reply{{.ResponseBody}})
	}
}
{{end}}
