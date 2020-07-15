package assertion

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
	"strings"
)

func dotJoin(u, v string) string {
	if len(u) == 0 {
		return v
	}
	if len(v) == 0 {
		return u
	}
	return u + "." + v
}

type JSONBody map[string]interface{}

type Request struct {
	Body      []byte
	CacheBody interface{}
}

type caseContext struct {
	Packages     LinkedContext
	gd           *GoDynamicTestData
	testCasePath string
	mt           map[string]Matcher
}

func (c *caseContext) findTestCase(p string) *TestCase {
	for i := len(c.gd.TestCases) - 1; i >= 0; i-- {
		// todo: nearest match
		t := c.gd.TestCases[i]
		if strings.HasPrefix(dotJoin(t.Name, t.Path), p) {
			return t
		}
	}
	return nil
}

func newContext() (*caseContext, error) {
	return &caseContext{
		Packages: &linkedContext{
			name: "root",
		},
		gd: new(GoDynamicTestData),
	}, nil
}

type TestCase struct {
	Path string
	Name string
	Meta TestCaseHTTPMeta
}

type TestCaseHTTPMeta struct {
	Header   map[string]string
	Method   string
	Data     map[string]interface{}
	Encoding string
}

const (
	MetaHTTPHeader   = "HTTPHeader"
	MetaHTTPMethod   = "HTTPMethod"
	MetaHTTPData     = "HTTPData"
	MetaHTTPEncoding = "HTTPEncoding"
)

func (t *TestCase) GetMeta(k string) (v interface{}) {
	if k == MetaHTTPHeader {
		return t.Meta.Header
	}
	if k == MetaHTTPMethod {
		return t.Meta.Method
	}
	if k == MetaHTTPData {
		return t.Meta.Data
	}
	if k == MetaHTTPEncoding {
		return t.Meta.Encoding
	}
	return nil
}

type GoDynamicTestData struct {
	TestCases []*TestCase
}

func ensureJSONBody(req *Request, err *error) (body JSONBody) {
	if *err != nil {
		return
	}
	if req.CacheBody != nil {
		body = req.CacheBody.(JSONBody)
	} else {
		*err = json.Unmarshal(req.Body, body)
		if *err != nil {
			req.CacheBody = body
		}
	}
	return
}

func ensureVarLength(s []string, wantLen int, err *error) {
	if *err != nil {
		return
	}
	if len(s) != wantLen {
		panic("var length error")
	}
}

func generateCaseV1(s *SpecV1) (*GoDynamicTestData, error) {
	ctx, err := newContext()
	if err != nil {
		return nil, err
	}

	var existingPackages = map[string]Package{
		"std-assert":  namespaceStd,
		"json-assert": namespaceJSON,
	}

	for _, pd := range s.PackageDefs {
		if p, ok := existingPackages[pd.Path]; ok {
			ctx.Packages.Insert(pd.Namespace, newFunctionPackage(p))
		} else {
			panic("not found existing package")
		}
	}

	for _, td := range s.TestDefs {
		err = generateTestCaseV1(ctx, &td)
		if err != nil {
			return nil, err
		}
	}

	ctx.mt = map[string]Matcher{}
	for _, selector := range s.Selector {
		ctx.mt[selector.Name], err = newSelector("", selector.Case)
		if err != nil {
			return nil, err
		}
	}

	var inheritFn = func(k string, v interface{}, t *TestCase) {
		switch k {
		case "encoding":
			fallthrough
		case "http-encoding":
			for _, t := range ctx.gd.TestCases {
				if len(v.(string)) != 0 && len(t.Meta.Encoding) == 0 {
					t.Meta.Encoding = v.(string)
				}
			}
		case "method":
			fallthrough
		case "http-method":
			for _, t := range ctx.gd.TestCases {
				if len(v.(string)) != 0 && len(t.Meta.Method) == 0 {
					t.Meta.Method = v.(string)
				}
			}
		}
	}

	for i := range s.Default {
		for k, v := range s.Default[i] {
			switch k {
			case "encoding", "http-encoding", "method", "http-method":
				for _, t := range ctx.gd.TestCases {
					inheritFn(k, v, t)
				}
			default:
				if len(k) != 0 && k[0] == '$' {
					// todo: selector
					li := strings.LastIndex(k[1:], ").")
					if li == -1 {
						panic("wrong selector dot")
					}
					li += 2
					//k, p = xs[0], xs[1]
					matcher := newTestCaseMatcher(ctx, k[1:li])
					k := k[li+1:]
					for _, t := range ctx.gd.TestCases {
						if ok, err := matcher.Match(t); ok && err == nil {
							inheritFn(k, v, t)
						} else if err != nil {
							panic(err)
						}
					}
				} else {
					panic("not found")
				}
			}
		}
	}
	//s.Selector
	return ctx.gd, nil
}

func newTestCaseMatcher(ctx *caseContext, k string) Matcher {
	//$(.[!=(method, GET)]).encoding
	if len(k) < 4 || k[0] != '(' || k[len(k)-1] != ')' || k[len(k)-2] != ']' {
		panic("selector no paren")
	}
	xs := strings.SplitN(k[1:len(k)-2], "[", 2)
	if len(xs) <= 1 {
		panic("selector no bracket")
	}
	p, ms := strings.TrimSpace(xs[0]), xs[1]
	var balance, j, li = 0, 0, 0
	var sms []string
	for i := 0; i < len(ms); i++ {
		if ms[i] == '(' {
			balance = 1
			for j = i + 1; j < len(ms); j++ {
				if ms[j] == '(' {
					balance++
				} else if ms[j] == ')' {
					balance--
				}
				if (balance) == 0 {
					break
				}
			}
			if balance != 0 {
				panic("unbalanced selector")
			}
			for ; j < len(ms); j++ {
				if ms[j] == ',' {
					break
				}
			}
			sms = append(sms, strings.TrimSpace(ms[li:j]))
			li = j + 1
		}
	}
	var am AndMatcher
	am.matchers = append(am.matchers, newPathMatcher(p))
	for _, sm := range sms {
		if len(sm) < 2 || sm[len(sm)-1] != ')' {
			panic("sm selector no paren")
		}
		xs := strings.SplitN(sm[:len(sm)-1], "(", 2)
		if len(xs) <= 1 {
			panic("sm selector no paren")
		}
		fn, args := strings.TrimSpace(xs[0]), strings.Split(xs[1], ",")
		switch fn {
		case "!=":
			am.matchers = append(am.matchers, newNEQMatcher(args))
		default:
			panic("todo fn")
		}
	}
	return am
}

type TrueMatcher struct {
}

func (t2 TrueMatcher) Match(t *TestCase) (bool, error) {
	return true, nil
}

func newPathMatcher(p string) Matcher {
	if p != "." {
		panic("todo")
	}
	return TrueMatcher{}
}

func newNEQMatcher(args []string) Matcher {
	if len(args) != 2 {
		panic("neq accept 2 string arg")
	}
	switch strings.TrimSpace(args[0]) {
	case "method":
		fallthrough
	case "http-method":
		m, err := newStringSelector("", strings.TrimSpace(args[1]))
		if err != nil {
			panic(err)
		}
		return newMethodNEQMatcher(m)
	}
	panic("todo")
}

type MethodNEQMatcher struct{ StringMatcher }

func (m MethodNEQMatcher) Match(t *TestCase) (bool, error) {
	x, y := m.StringMatcher.MatchString(t.Meta.Method)
	return !x, y
}

func newMethodNEQMatcher(m StringMatcher) Matcher {
	return MethodNEQMatcher{m}
}

func insertUsingV1(ctx *caseContext, us map[string]string) (err error) {
	if us == nil {
		return nil
	}
	for k, v := range us {
		if len(k) == 0 {
			panic("nil k")
		}
		if k[0] == '$' {
			p := RepositionCtx(ctx.Packages, v)
			if p == nil {
				p = RepositionCtx(ctx.Packages, dotJoin("root", v))
			}
			if p == nil {
				panic("package not found")
			}
			ctx.Packages.Insert(k[1:], copyLink(p))
		} else {
			f := findCheckFunc(ctx.Packages, v)
			if f == nil {
				panic("func not found")
			}
			ctx.Packages.InsertFunc(k, f)
		}
	}
	return nil
}

func generateTestCaseV1(ctx *caseContext, td *TestDef) (err error) {
	var inherit, name = "", td.Name
	ss := strings.SplitN(name, ".", 2)
	if len(ss) == 2 {
		name, inherit = ss[0], ss[1]
		if len(inherit) == 0 {
			inherit = "root"
		}
	}

	lc := &linkedContext{name: name}
	ctx.Packages.Insert(name, lc)
	ctx.Packages = lc

	var ts = TestCase{Path: ctx.testCasePath, Name: name}

	ctx.gd.TestCases = append(ctx.gd.TestCases, &ts)
	if td.Using != nil || td.UsingForce != nil {
		//ctx.Packages = clonePackages(ctx.Packages)

		if err = insertUsingV1(ctx, td.Using); err != nil {
			return
		}

		//if err = insertUsingV1(ctx, td.UsingForce, true); err != nil {
		//	return
		//}
	}

	debugPrint(os.Stdout, td.Meta, name, 0)
	fmt.Println()
	parseMeta(&ts, td.Meta)
	for _, i := range td.Inherit {
		inheritTestCase(&ts, ctx.findTestCase(i))
	}
	if len(inherit) != 0 {
		inheritTestCase(&ts, ctx.findTestCase(inherit))
	}

	ctx.testCasePath = dotJoin(td.Name, ts.Path)
	for _, std := range td.Cases {

		err = generateTestCaseV1(ctx, &std)
		if err != nil {
			return
		}
	}
	ctx.testCasePath = ts.Path
	ctx.Packages = ctx.Packages.Last()
	return
}

func inheritTestCase(dst *TestCase, src *TestCase) {
	inheritMetaHeader(dst, src)
	inheritMetaMethod(dst, src)
	inheritMetaData(dst, src)
	inheritMetaEncoding(dst, src)
}

func inheritMetaEncoding(dst *TestCase, src *TestCase) {
	if len(dst.Meta.Encoding) == 0 {
		if len(src.Meta.Encoding) != 0 {
			dst.Meta.Encoding = src.Meta.Encoding
		}
	}
}

func inheritMetaData(dst *TestCase, src *TestCase) {
	if len(src.Meta.Data) != 0 {
		if dst.Meta.Data == nil {
			dst.Meta.Data = make(map[string]interface{})
		}
		inheritMapType(dst.Meta.Data, src.Meta.Data)
	}
}

func deepCopy(v interface{}) interface{} {
	switch v := v.(type) {
	case map[string]interface{}:
		var nv = make(map[string]interface{})
		for k, vv := range v {
			nv[k] = deepCopy(vv)
		}
		return nv
	default:
		t := reflect.TypeOf(v)
		if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
			xv := reflect.ValueOf(v)
			var l = xv.Len()
			var nv = reflect.MakeSlice(t, l, l)
			var values []reflect.Value
			for i := 0; i < l; i++ {
				values = append(values, reflect.ValueOf(deepCopy(xv.Index(i).Interface())))
			}
			nv = reflect.Append(nv, values...)
			return nv.Interface()
		}
		return v
	}
}

func toJSONBody(v interface{}) interface{} {
	switch v := v.(type) {
	case map[interface{}]interface{}:
		var nv = make(map[string]interface{})
		for k, vv := range v {
			nv[k.(string)] = toJSONBody(vv)
		}
		return nv
	default:
		return v
	}
}

func inheritMapType(dst map[string]interface{}, src map[string]interface{}) {
	for k, v := range src {
		switch v := v.(type) {
		case map[string]interface{}:
			if dv, ok := dst[k]; !ok {
				dst[k] = deepCopy(v)
			} else if dv, ok := dv.(map[string]interface{}); ok {
				inheritMapType(dv, v)
			}
		default:
			if reflect.TypeOf(v).Kind() == reflect.Array {
				dst[k] = deepCopy(v)
			}
			if _, ok := dst[k]; !ok {
				dst[k] = v
			}
		}
	}
}

func inheritMetaHeader(dst *TestCase, src *TestCase) {
	if src.Meta.Header != nil && len(src.Meta.Header) != 0 {
		if dst.Meta.Header == nil {
			dst.Meta.Header = make(map[string]string)
		}
		for k, v := range src.Meta.Header {
			if _, ok := dst.Meta.Header[k]; !ok {
				dst.Meta.Header[k] = v
			}
		}
	}
}

func inheritMetaMethod(dst *TestCase, src *TestCase) {
	if len(dst.Meta.Method) == 0 {
		if len(src.Meta.Method) != 0 {
			dst.Meta.Method = src.Meta.Method
		}
	}
}

func parseMeta(ts *TestCase, meta map[string]interface{}) {
	ts.Meta.Header = parseMetaHeader(meta)
	ts.Meta.Method = parseMetaMethod(meta)
	ts.Meta.Data = parseMetaData(meta)
	ts.Meta.Encoding = parseMetaEncoding(meta)
}

func parseMetaData(meta map[string]interface{}) map[string]interface{} {
	switch d := meta["data"].(type) {
	case map[string]interface{}:
		return toJSONBody(d).(map[string]interface{})
	case map[interface{}]interface{}:
		return toJSONBody(d).(map[string]interface{})
	case nil:
		return nil
	default:
		panic("data type error")
	}
}

func parseMetaEncoding(meta map[string]interface{}) string {
	e, ok := meta["http-encoding"]
	if !ok {
		e = meta["encoding"]
	}
	switch e := e.(type) {
	case string:
		return e
	case nil:
		return ""
	default:
		panic("encoding type error")
	}
}

func parseMetaHeader(meta map[string]interface{}) (nv map[string]string) {
	h, ok := meta["http-header"]
	if !ok {
		h = meta["header"]
	}
	switch h := h.(type) {
	case map[string]string:
		nv = h
		return
	case map[string]interface{}:
		nv = make(map[string]string)
		for k, v := range h {
			nv[k] = v.(string)
		}
		return
	case map[interface{}]interface{}:
		nv = make(map[string]string)
		for k, v := range h {
			nv[k.(string)] = v.(string)
		}
		return
	case nil:
		return
	default:
		panic("header type error")
	}
}

func parseMetaMethod(meta map[string]interface{}) string {
	e, ok := meta["http-method"]
	if !ok {
		e = meta["method"]
	}
	switch e := e.(type) {
	case string:
		return e
	case nil:
		return ""
	default:
		panic("encoding type error")
	}
}

func Load() {
	f, err := os.Open("test.yaml")
	if err != nil {
		panic(err)
	}
	var spec SpecV1
	err = yaml.NewDecoder(f).Decode(&spec)
	if err != nil {
		panic(err)
	}
	//debugPrint(os.Stdout, spec, "", 0)
	gd, err := generateCaseV1(&spec)
	if err != nil {
		panic(err)
	}
	for _, x := range gd.TestCases {
		fmt.Println("Name:", x.Name)
		fmt.Println("Path:", x.Path)
		fmt.Println("Method:", x.Meta.Method)
		fmt.Println("Data:", x.Meta.Data)
		fmt.Println("Encoding:", x.Meta.Encoding)
		fmt.Println("Header:", x.Meta.Header)
		fmt.Println("----------------------------------------------------------------------")
	}
}
