package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"
)

type field struct {
	FieldName string
	tag
}

type tag struct {
	IsInt     bool
	Required  bool
	Enum      []string
	Min       string
	Max       string
	ParamName string
	Default   string
}

type meta struct {
	URL    string
	Auth   bool
	Method string
}

type handler struct {
	HandlerMethod   string
	meta
	Parameter       string
	Result          string
	StructParameter []field
}

type generationAPI struct {
	Handlers    map[string][]handler
	Validator   map[string][]field
	IsInt       bool
	IsParamName bool
}

func main() {
	finPath := os.Args[1]
	foutPath := os.Args[2]

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, finPath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	fout, err := os.Create(foutPath)
	if err != nil {
		log.Fatal(err)
	}

	addHeader(fout, node.Name.Name)

	// parse
	result := generationAPI{make(map[string][]handler), make(map[string][]field), false, false}

	for _, decl := range node.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			parseStruct(decl, &result)
		case *ast.FuncDecl:
			parseFunc(decl, &result)
		}
	}

	for i1, h := range result.Handlers {
		for i2, v2 := range h {
			result.Handlers[i1][i2].StructParameter = append(v2.StructParameter, result.Validator[v2.Parameter]...)
		}
	}

	// generate
	generate(fout, &result)
}

func addHeader(fout *os.File, packageName string) {
	fmt.Fprintf(fout, `package %s
import (
	"encoding/json"
	"errors"
	//"fmt"
	//"io/ioutil"
 	"net/http"
	//"net/url"
	"strconv"
	"strings"
)
`, packageName)
}

// parsing functions
func parseStruct(f *ast.GenDecl, g *generationAPI) {
	for _, spec := range f.Specs {
		if currType, ok := spec.(*ast.TypeSpec); ok {
			if currStruct, ok := currType.Type.(*ast.StructType); ok {
				handleStruct(currStruct, currType, g)
			}
		}
	}
}

func handleStruct(currStruct *ast.StructType, currType *ast.TypeSpec, g *generationAPI) {
	for _, field := range currStruct.Fields.List {
		if field.Tag != nil {
			tagValue := ""
			if strings.HasPrefix(field.Tag.Value, "`apivalidator:") {
				tagValue = strings.TrimLeft(field.Tag.Value, "`apivalidator:")
			}

			g.IsInt = field.Type.(*ast.Ident).Name == "int"
			g.IsParamName = strings.Contains(field.Tag.Value, "paramname")

			parseField(currType.Name.Name, tagValue, field.Names[0].Name, g)
		}
	}
}

func parseField(currTypeName, tagValue, fieldName string, g *generationAPI) {
	f := field{}
	v := strings.Split(strings.Replace(strings.Trim(tagValue, "/`"), "\"", "", -1), ",")
	for _, value := range v {
		f.IsInt = g.IsInt

		if value == "required" {
			f.Required = true
			f.FieldName = fieldName
		}

		if s := strings.Split(value, "="); len(s) > 1 {
			//check for enum
			if enumFields := strings.Split(s[1], "|"); len(enumFields) > 1 {
				for _, enumField := range enumFields {
					f.FieldName = fieldName
					f.Enum = append(f.Enum, enumField)
				}
			}

			switch s[0] {
			case "min":
				f.FieldName = fieldName
				f.Min = s[1]
			case "max":
				f.FieldName = fieldName
				f.Max = s[1]
			case "paramname":
				f.FieldName = fieldName
				f.ParamName = s[1]
			case "default":
				f.FieldName = fieldName
				f.Default = s[1]
			}
		}
	}

	g.Validator[currTypeName] = append(g.Validator[currTypeName], f)
}

func parseFunc(f *ast.FuncDecl, g *generationAPI) {
	if f.Doc == nil {
		return
	}

	h := handler{}

	for _, comment := range f.Doc.List {
		if strings.HasPrefix(comment.Text, "// apigen:api") {
			h.HandlerMethod = strings.ToLower(f.Name.Name)
			apigenDoc := []byte(strings.TrimLeft(comment.Text, "// apigen:api"))
			handlerMeta := meta{}
			json.Unmarshal(apigenDoc, &handlerMeta)
			h.meta = handlerMeta

			var key string
			if f.Recv != nil {
				switch a := f.Recv.List[0].Type.(type) {
				case *ast.StarExpr:
					key = a.X.(*ast.Ident).Name
				}
			}

			if f.Type.Params.List != nil {
				for _, p := range f.Type.Params.List {
					switch a := p.Type.(type) {
					case *ast.Ident:
						h.Parameter = a.Name
					}
				}
			}

			if f.Type.Results.List != nil && len(f.Type.Results.List) != 0 {
				switch a := f.Type.Results.List[0].Type.(type) {
				case *ast.StarExpr:
					h.Result = a.X.(*ast.Ident).Name
				}
			}

			g.Handlers[key] = append(g.Handlers[key], h)
		}
	}
}

// generating part
func generate(out *os.File, g *generationAPI) {
	fmt.Fprintln(out)
	// define error vriables
	fmt.Fprintln(out,
		`var (
	errorUnknown    = errors.New("unknown method")
	errorBad        = errors.New("bad method")
	errorEmptyLogin = errors.New("login must me not empty")
)

type JsonError struct {`)
	fmt.Fprintln(out, "\tError string `json:\"error\"`")
	fmt.Fprintln(out, "}")


	serveHTTPTpl.Execute(out, g)
	structTpl.Execute(out, g)
	handlerTpl.Execute(out, g)
}

var (
	isMethodPost = func(method string) bool {
		return method == "POST"
	}

	funcMap = template.FuncMap{
		"Title":        strings.Title,
		"isMethodPost": isMethodPost,
		"toLower":      strings.ToLower,
		"joinComma":    func(fields []string) string { return strings.Join(fields, ", ") },
		"FieldNameJoinComma": func(fields []field) string {
			var s []string
			for _, f := range fields {
				s = append(s, strings.ToLower(f.FieldName))
			}
			return strings.Join(s, ", ")
		},
	}

	serveHTTPTpl = template.Must(template.New("serveHTTPTpl").Parse(`
{{ range $key, $val := .Handlers }}
func (h *{{ $key }} ) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	{{range $val }} case "{{ .URL }}":
		h.{{ .HandlerMethod }}(w,r)
	{{end}} default:
		js, _ := json.Marshal(JsonError{errorUnknown.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(js)
		return
	}
}
{{end}}`))

	structTpl = template.Must(template.New("structTpl").Funcs(funcMap).Parse(`
{{ range $key, $val := .Handlers }}{{ range $v := $val }}
type Response{{ $key }}{{ .HandlerMethod | Title  }}  struct {
	*{{ .Result }}` + " `json:\"response\"`\n\t" +
		`JsonError
}
{{end}}{{end}}`))

	handlerTpl = template.Must(template.New("handlerTpl").Funcs(funcMap).Parse(`
{{ range $key, $val := .Handlers }}
{{ range $v := $val }}
func (h *{{ $key }}) {{ .HandlerMethod }}(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
{{if .Method | isMethodPost}}
	if r.Method != "POST" {
		js, _ := json.Marshal(JsonError{errorBad.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(js)
		return
	}{{if .Auth }}
	if r.Header.Get("X-Auth") != "100500" {
		js, _ := json.Marshal(JsonError{"unauthorized"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write(js)
		return
	}{{end}}
	r.ParseForm()
{{ range $p := .StructParameter }}{{ if .IsInt}}
	{{ .FieldName | toLower}} , err := strconv.Atoi(r.Form.Get("{{ .FieldName | toLower}}"))
	if err != nil {
		js, _ := json.Marshal(JsonError{"{{ .FieldName | toLower}} must be int"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(js)
		return
	}
{{ else }}
	{{ .FieldName | toLower }} := r.Form.Get("{{ .FieldName | toLower}}")
{{ end }}{{ if .Required }}{{ if .IsInt}}
	if {{ .FieldName | toLower }} == nil {
		js, _ := json.Marshal(JsonError{errorEmptyLogin.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(js)
		return
	}
{{ else }}	if {{ .FieldName | toLower }} == "" {
		js, _ := json.Marshal(JsonError{errorEmptyLogin.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(js)
		return
	}
{{ end }}{{ end }}{{ if .Min  }}{{ if .IsInt}}
	if {{ .FieldName | toLower }} < {{ .Min }}  {
		js, _ := json.Marshal(JsonError{"{{ .FieldName | toLower }} must be >= {{ .Min }}"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(js)
		return
	}
{{else}}	if len({{ .FieldName | toLower }}) < {{ .Min }}  {
		js, _ := json.Marshal(JsonError{"{{ .FieldName | toLower }} len must be >= {{ .Min }}"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(js)
		return
	}
{{ end }}{{ end }}{{ if .Max  }}{{ if .IsInt}}
	if {{ .FieldName | toLower }} > {{ .Max }}  {
		js, _ := json.Marshal(JsonError{"{{ .FieldName | toLower }} must be <= {{ .Max }}"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(js)
		return
	}
{{else}}	if len({{ .FieldName | toLower }}) > {{ .Max }}  {
		js, _ := json.Marshal(JsonError{"{{ .FieldName | toLower }} len must be <= {{ .Max }}"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(js)
		return
	}
{{ end }}{{ end }}
{{ if .ParamName  }}
	paramname_{{ .FieldName | toLower }} := r.Form.Get("{{ .ParamName }}")
	if paramname_{{ .FieldName | toLower }} == "" {
		{{ .FieldName | toLower }} = strings.ToLower({{ .FieldName | toLower }})
	}  else {
		{{ .FieldName | toLower }} = paramname_{{ .FieldName | toLower }}
	}
{{ end }}{{ if .Enum  }}{{ if .Default }}
	if {{ .FieldName | toLower }} == "" {
		{{ .FieldName | toLower }} = "{{ .Default }}"
	}{{ end }}
	m := make(map[string]bool)
{{ range $p := .Enum }}	m["{{ $p }}"] = true
{{ end }}
	_, prs := m[{{ .FieldName | toLower }}]
	if prs == false {
		js, _ := json.Marshal(JsonError{"{{ .FieldName | toLower }} must be one of [{{ .Enum | joinComma }}]"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(js)
		return
	}
{{ end }}{{ end }}{{ else }}{{ range $p := .StructParameter }}
	var {{ .FieldName | toLower }} string
	switch r.Method {
	case "GET":
		{{ .FieldName | toLower }} = r.URL.Query().Get("{{ .FieldName | toLower }}")
		if {{ .FieldName | toLower }} == "" {
			js, _ := json.Marshal(JsonError{errorEmptyLogin.Error()})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(js)
			return
		}
	case "POST":
		r.ParseForm()
		{{ .FieldName | toLower }} = r.Form.Get("{{ .FieldName | toLower }}")
		if {{ .FieldName | toLower }} == "" {
			js, _ := json.Marshal(JsonError{errorEmptyLogin.Error()})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(js)
			return
		}
	}
{{end}}{{end}}
	{{ .Parameter }} := {{ .Parameter }}{ {{ .StructParameter | FieldNameJoinComma }}   }
	{{ .Result | toLower }}, err := h.{{ .HandlerMethod | Title }}(ctx, {{ .Parameter }})
	if err != nil {
		switch err.(type) {
		case ApiError:
			js, _ := json.Marshal(JsonError{err.(ApiError).Err.Error()})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(err.(ApiError).HTTPStatus)
			w.Write(js)
			return
		default:
			js, _ := json.Marshal(JsonError{"bad user"})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(js)
			return
		}
	}
	js, _ := json.Marshal(Response{{ $key }}{{ .HandlerMethod | Title  }} { {{ .Result | toLower }}, JsonError{""}})
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}{{end}}{{end}}`))
)
