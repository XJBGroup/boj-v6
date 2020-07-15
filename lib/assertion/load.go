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

	var selectors = map[string]Matcher{}

	for _, selector := range s.Selector {
		selectors[selector.Name], err = newSelector("", selector.Case)
		if err != nil {
			return nil, err
		}
	}

	//s.Default
	//s.Selector
	return ctx.gd, nil
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
	parseMetaHeader(ts, meta)
	parseMetaMethod(ts, meta)
	parseMetaData(ts, meta)
	parseMetaEncoding(ts, meta)
}

func parseMetaData(ts *TestCase, meta map[string]interface{}) {
	switch d := meta["data"].(type) {
	case map[string]interface{}:
		ts.Meta.Data = toJSONBody(d).(map[string]interface{})
	case map[interface{}]interface{}:
		ts.Meta.Data = toJSONBody(d).(map[string]interface{})
	case nil:
		return
	default:
		panic("data type error")
	}
}

func parseMetaEncoding(ts *TestCase, meta map[string]interface{}) {
	e, ok := meta["http-encoding"]
	if !ok {
		e = meta["encoding"]
	}
	switch e := e.(type) {
	case string:
		ts.Meta.Encoding = e
	case nil:
		return
	default:
		panic("encoding type error")
	}
}

func parseMetaHeader(ts *TestCase, meta map[string]interface{}) {
	h, ok := meta["http-header"]
	if !ok {
		h = meta["header"]
	}
	switch h := h.(type) {
	case map[string]string:
		ts.Meta.Header = h
	case map[string]interface{}:
		ts.Meta.Header = make(map[string]string)
		for k, v := range h {
			ts.Meta.Header[k] = v.(string)
		}
	case map[interface{}]interface{}:
		ts.Meta.Header = make(map[string]string)
		for k, v := range h {
			ts.Meta.Header[k.(string)] = v.(string)
		}
	case nil:
		return
	default:
		panic("header type error")
	}
}

func parseMetaMethod(ts *TestCase, meta map[string]interface{}) {
	h, ok := meta["http-header"]
	if !ok {
		h = meta["header"]
	}
	switch h := h.(type) {
	case map[string]string:
		ts.Meta.Header = h
	case map[string]interface{}:
		ts.Meta.Header = make(map[string]string)
		for k, v := range h {
			ts.Meta.Header[k] = v.(string)
		}
	case map[interface{}]interface{}:
		ts.Meta.Header = make(map[string]string)
		for k, v := range h {
			ts.Meta.Header[k.(string)] = v.(string)
		}
	case nil:
		return
	default:
		panic("header type error")
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
		fmt.Println(x.Name, x.Path)
	}
}
