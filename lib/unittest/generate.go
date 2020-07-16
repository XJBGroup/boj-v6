package unittest

import "strings"

func generateCaseV1(s *SpecV1) (*GoDynamicTestData, error) {

	var existingPackages = map[string]Package{
		"std-assert":  namespaceStd,
		"json-assert": namespaceJSON,
	}

	var inheritFnMap = map[string]MetaOperation{
		MetaEncoding:     stringMetaPropertyOperation(MetaEncoding),
		MetaMethod:       stringMetaPropertyOperation(MetaMethod),
		MetaUrl:          stringMetaPropertyOperation(MetaUrl),
		MetaData:         dataBodyMetaPropertyOperation(MetaData),
		MetaHeader:       stringMapMetaPropertyOperation(MetaHeader),
		MetaHTTPEncoding: stringMetaPropertyOperation(MetaHTTPEncoding),
		MetaHTTPMethod:   stringMetaPropertyOperation(MetaHTTPMethod),
		MetaHTTPHeader:   stringMapMetaPropertyOperation(MetaHTTPHeader),
	}

	var parseMetaFnMap = map[string]MetaParser{
		"encoding":      parseStringProperty(MetaEncoding),
		"method":        parseStringProperty(MetaMethod),
		"url":           parseStringProperty(MetaUrl),
		"data":          parseDataBodyProperty(MetaData),
		"header":        parseStringMapProperty(MetaHeader),
		"http-encoding": parseStringProperty(MetaHTTPEncoding),
		"http-method":   parseStringProperty(MetaHTTPMethod),
		"http-header":   parseStringMapProperty(MetaHTTPHeader),
	}

	ctx, err := newContext(&Option{metaOperationMap: inheritFnMap, ParseMetaMap: parseMetaFnMap})
	if err != nil {
		return nil, err
	}

	for _, pd := range s.PackageDefs {
		if p, ok := existingPackages[pd.Path]; ok {
			ctx.Packages.Insert(pd.Namespace, newFunctionPackage(pd.Namespace, p))
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

	for i := len(ctx.gd.TestCases) - 1; i >= 0; i-- {
		if ctx.gd.TestCases[i].Abstract {
			for j := i + 1; j < len(ctx.gd.TestCases); j++ {
				ctx.gd.TestCases[j-1] = ctx.gd.TestCases[j]
			}
			ctx.gd.TestCases = ctx.gd.TestCases[:len(ctx.gd.TestCases)-1]
		}
	}

	for i := range s.Default {
		for k, v := range s.Default[i] {
			var matcher Matcher = TrueMatcher{}
			if len(k) != 0 && k[0] == '$' {
				// todo: selector
				li := strings.LastIndex(k[1:], ").")
				if li == -1 {
					panic("wrong selector dot")
				}
				li += 2
				//k, p = xs[0], xs[1]
				matcher = newTestCaseMatcher(ctx, k[1:li])
				k = k[li+1:]
			}

			k, v, err = ctx.parseMetaKV(k, v)
			if err != nil {
				panic(err)
			}

			for _, t := range ctx.gd.TestCases {
				if ok, err := matcher.Match(t); ok && err == nil {
					err = ctx.inheritPropertyKV(k, v, t.Meta)
					if err != nil {
						panic(err)
					}
				} else if err != nil {
					panic(err)
				}
			}
		}
	}
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

func mapString(f func(string) string, x []string) []string {
	for i := range x {
		x[i] = f(x[i])
	}
	return x
}

func generateTestCaseV1(ctx *caseContext, td *TestDef) (err error) {

	var abstract = false
	var inherit []string
	var name = td.Name
	if len(name) > 0 && name[0] == '~' {
		abstract = true
		name = name[1:]
	}
	li := strings.IndexByte(name, '<')
	if li != -1 {
		if name[len(name)-1] != '>' {
			panic("missing rag")
		}
		name, inherit = name[:li], mapString(strings.TrimSpace, strings.Split(name[li+1:len(name)-1], ","))
	} else {
		ss := strings.SplitN(name, ".", 2)
		if len(ss) == 2 {
			name, inherit = ss[0], ss[1:]
			if len(inherit[0]) == 0 {
				inherit[0] = "root"
			}
		}
	}

	lc := &linkedContext{name: name}
	ctx.Packages.Insert(name, lc)
	ctx.Packages = lc

	var ts = TestCase{Path: ctx.testCasePath, Name: name, Abstract: abstract || td.Abstract}

	ctx.gd.TestCases = append(ctx.gd.TestCases, &ts)
	if td.Using != nil {
		if err = insertUsingV1(ctx, td.Using); err != nil {
			return
		}
	}

	//debugPrint(os.Stdout, td.Meta, name, 0)
	//fmt.Println()

	err = ctx.parseMeta(td.Meta, &ts)
	if err != nil {
		return
	}
	for _, i := range append(td.Inherit, inherit...) {
		inheritTarget := ctx.findTestCase(i)
		if inheritTarget == nil {
			panic("miss inherit object")
		}
		err = ctx.inheritProperty(ts.Meta, inheritTarget.Meta)
		if err != nil {
			return
		}

		ts.Assertions = append(ts.Assertions, inheritTarget.Assertions...)
	}

	var assertions = td.Assertion
	if len(td.Assert) != 0 {
		assertions = append(assertions, append([]string{"Assert"}, td.Assert...))
	}
	for _, a := range assertions {
		if len(a) == 0 {
			panic("nil assertion")
		}
		fn := findCheckFunc(ctx.Packages, a[0])
		if fn == nil {
			panic("nil assertion")
		}
		ts.Assertions = append(ts.Assertions, Assertion{
			F:     fn,
			FN:    a[0],
			VArgs: a[1:],
		})
	}
	//Assertion  [][]string             `yaml:"assertion"`
	//Assert     []string               `yaml:"assert"`

	ctx.testCasePath = dotJoin(ts.Name, ts.Path)
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
