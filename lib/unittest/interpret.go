package unittest

import (
	"fmt"
	"strings"
)

func interpretPackageDef(ctx *caseContext, s *SpecV1) (err error) {

	var existingPackages = map[string]Package{
		"std-assert":  namespaceStd,
		"json-assert": namespaceJSON,
	}

	for _, pd := range s.PackageDefs {
		if p, ok := existingPackages[pd.Path]; ok {
			ctx.Packages.Insert(pd.Namespace, newFunctionPackage(pd.Namespace, p))
		} else {
			panic("not found existing package")
		}
	}
	return
}

func interpretSelectors(ctx *caseContext, s *SpecV1) (err error) {

	if ctx.Selectors == nil {
		ctx.Selectors = map[string]Matcher{}
	}

	for _, selector := range s.Selector {
		ctx.Selectors[selector.Name], err = newSelector("", selector.Case)
		if err != nil {
			return err
		}
	}
	return
}

func interpretTestCases(ctx *caseContext, s *SpecV1) (err error) {
	for _, td := range s.TestDefs {
		err = interpretTestCaseV1(ctx, &td)
		if err != nil {
			return err
		}
	}
	return
}

func interpretTestCaseV1(ctx *caseContext, td *TestDef) (err error) {
	// generate recursively
	// use backtracking algorithm

	// parse test case name grammar
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

	// create linked context
	lc := &linkedContext{name: name}
	// update ctx.Packages
	ctx.Packages.Insert(name, lc)
	ctx.Packages = lc

	var ts = TestCase{Path: ctx.TestCasePath, Name: name, Abstract: abstract || td.Abstract}
	ctx.Gd.TestCases = append(ctx.Gd.TestCases, &ts)

	// interpret using statement
	if td.Using != nil {
		if err = interpretUsingV1(ctx, td.Using); err != nil {
			return
		}
	}

	// parse metas
	err = ctx.parseMeta(td.Meta, &ts)
	if err != nil {
		return
	}

	// interpret inherit statement
	for _, i := range append(td.Inherit, inherit...) {
		inheritTarget := ctx.findTestCase(i)
		if inheritTarget == nil {
			panic("miss inherit object")
		}
		err = ctx.inheritProperty(ts.Meta, inheritTarget.Meta)
		if err != nil {
			return
		}

		ts.Script = append(ts.Script, inheritTarget.Script...)
	}

	// interpret script block (assert statement appended)
	var scriptLines = td.Script
	if len(td.Assert) != 0 {
		scriptLines = append(scriptLines, append([]interface{}{"Assert"}, td.Assert...))
	}
	for _, a := range scriptLines {
		if len(a) == 0 {
			panic("nil assertion")
		}
		fn := findCheckFunc(ctx.Packages, a[0].(string))
		if fn == nil {
			panic("nil assertion")
		}
		ts.Script = append(ts.Script, Assertion{
			F:     fn,
			FN:    a[0].(string),
			VArgs: a[1:],
		})
	}

	// update ctx.testCasePath
	ctx.TestCasePath = dotJoin(ts.Name, ts.Path)

	for _, std := range td.Cases {

		err = interpretTestCaseV1(ctx, &std)
		if err != nil {
			return
		}
	}

	// remove ctx testCasePath/package value
	ctx.TestCasePath = ts.Path
	ctx.Packages = ctx.Packages.Last()
	return
}

func interpretUsingV1(ctx *caseContext, us map[string]string) (err error) {
	if us == nil {
		return nil
	}
	for k, v := range us {
		if len(k) == 0 {
			return fmt.Errorf("nil k")
		}
		if k[0] == '$' {
			p := RepositionCtx(ctx.Packages, v)
			if p == nil {
				p = RepositionCtx(ctx.Packages, dotJoin("root", v))
			}
			if p == nil {
				return fmt.Errorf("package not found")
			}
			ctx.Packages.Insert(k[1:], copyLink(p))
		} else {
			f := findCheckFunc(ctx.Packages, v)
			if f == nil {
				return fmt.Errorf("func not found")
			}
			ctx.Packages.InsertFunc(k, f)
		}
	}
	return nil
}
