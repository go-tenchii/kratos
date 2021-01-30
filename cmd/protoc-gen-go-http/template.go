package main

import (
	"bytes"
	"html/template"
	"strings"
)

var httpTemplate = `
type {{.ServiceType}}Service interface {
{{range .MethodSets}}
	{{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
{{end}}
}

func Register{{.ServiceType}}HTTPServer(s *http1.Server, srv {{.ServiceType}}Service) {
	r := s.Route("/")
{{range .Methods}}
	r.{{.Method}}("{{.Path}}", func(res http.ResponseWriter, req *http.Request) {
		in := new({{.Request}})
		{{if ne (len .Vars) 0}}
		if err := http1.BindVars(req, in); err != nil {
			s.Error(res, req, err)
			return
		}
		{{end}}
		{{if eq .Body ""}}
		if err := http1.BindForm(req, in); err != nil {
			s.Error(res, req, err)
			return
		}
		{{else if eq .Body ".*"}}
		if err := s.Decode(req, in); err != nil {
			s.Error(res, req, err)
			return
		}
		{{else}}
		if err := s.Decode(req, in{{.Body}}); err != nil {
			s.Error(res, req, err)
			return
		}
		{{end}}
		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.({{$.ServiceType}}Service).{{.Name}}(ctx, in)
		}
		out, err := s.Invoke(req.Context(), in, h)
		if err != nil {
			s.Error(res, req, err)
			return
		}
		s.Encode(res, req, out{{.ResponseBody}})
	})
{{end}}
}
`

type serviceDesc struct {
	ServiceType string // Greeter
	ServiceName string // helloworld.Greeter
	Metadata    string // api/helloworld/helloworld.proto
	Methods     []*methodDesc
	MethodSets  map[string]*methodDesc
}

type methodDesc struct {
	// method
	Name    string
	Num     int
	Vars    []string
	Forms   []string
	Request string
	Reply   string
	// http_rule
	Path         string
	Method       string
	Body         string
	ResponseBody string
}

func (s *serviceDesc) execute() string {
	s.MethodSets = make(map[string]*methodDesc)
	for _, m := range s.Methods {
		s.MethodSets[m.Name] = m
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(httpTemplate))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return string(buf.Bytes())
}
