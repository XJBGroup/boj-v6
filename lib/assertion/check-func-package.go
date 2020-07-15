package assertion

import (
	"fmt"
	"strings"
)

type CheckFunc = func(Request, ...string) (string, error)
type Package = map[string]CheckFunc
type LinkedContext interface {
	Last() LinkedContext
	Name() string
	Get(s string) LinkedContext
	GetFunc(s string) CheckFunc
	SetName(s string)
	SetLast(LinkedContext)
	Insert(k string, ctx LinkedContext)
	InsertFunc(s string, f CheckFunc)
}

type linkedContext struct {
	last LinkedContext
	name string
	next map[string]LinkedContext
	fns  Package
}

func newFunctionPackage(fns Package) *linkedContext {
	return &linkedContext{fns: fns}
}

func copyLink(l LinkedContext) *copiedLinkedContext {
	return &copiedLinkedContext{attached: l}
}

type copiedLinkedContext struct {
	last     LinkedContext
	name     string
	attached LinkedContext
}

func (c copiedLinkedContext) Last() LinkedContext {
	return c.last
}

func (c copiedLinkedContext) Name() string {
	return c.name
}

func (c copiedLinkedContext) Get(s string) LinkedContext {
	return c.attached.Get(s)
}

func (c copiedLinkedContext) GetFunc(s string) CheckFunc {
	return c.attached.GetFunc(s)
}

func (c *copiedLinkedContext) SetName(s string) {
	c.name = s
}

func (c *copiedLinkedContext) SetLast(l LinkedContext) {
	c.last = l
}

func (c copiedLinkedContext) Insert(string, LinkedContext) {
	panic("not changeable")
}

func (c copiedLinkedContext) InsertFunc(string, CheckFunc) {
	panic("not changeable")
}
func (c linkedContext) Name() string {
	return c.name
}

func (c *linkedContext) SetName(s string) {
	c.name = s
}

func (c linkedContext) Last() LinkedContext {
	return c.last
}

func (c linkedContext) Get(s string) LinkedContext {
	if c.next == nil {
		return nil
	}
	return c.next[s]
}

func (c linkedContext) GetFunc(s string) CheckFunc {
	if c.fns == nil {
		return nil
	}
	return c.fns[s]
}

func (c *linkedContext) SetLast(ctx LinkedContext) {
	c.last = ctx
}

func (c *linkedContext) Insert(k string, ctx LinkedContext) {
	if c.next == nil {
		c.next = make(map[string]LinkedContext)
	}
	if _, ok := c.next[k]; ok {
		panic("conflict")
	}
	c.next[k] = ctx
	ctx.SetLast(c)
}

func (c *linkedContext) InsertFunc(k string, f CheckFunc) {
	if c.fns == nil {
		c.fns = make(map[string]CheckFunc)
	}
	if _, ok := c.fns[k]; ok {
		panic("conflict")
	}
	c.fns[k] = f
}

func RepositionCtx(p LinkedContext, pn string) LinkedContext {
	if p == nil {
		return nil
	}
	var firstPn string
	var pnIndex = strings.IndexByte(pn, '.')
	if pnIndex != -1 {
		firstPn, pn = pn[:pnIndex], pn[pnIndex+1:]
	} else {
		firstPn, pn = pn, ""
	}

	if len(firstPn) == 0 {
		firstPn = "root"
	}

	for p != nil && p.Name() != firstPn {
		p = p.Last()
	}
	if p == nil {
		return nil
	}

	if len(pn) == 0 {
		return p
	}

	for p != nil && len(pn) != 0 {
		pnIndex = strings.IndexByte(pn, '.')
		if pnIndex != -1 {
			firstPn, pn = pn[:pnIndex], pn[pnIndex+1:]
		} else {
			firstPn, pn = pn, ""
		}
		p = p.Get(firstPn)
	}

	if p == nil {
		return nil
	}

	return p
	//if pn[0] == '.' {
	//	pn = pn[1:]
	//	//if !strings.HasPrefix(thisPath, pn) {
	//	//	panic("root select error")
	//	//}
	//	//thisPath = thisPath[len(pn):]
	//	//if thisPath[0] == '.' {
	//	//	thisPath = thisPath[1:]
	//	//}
	//	return p[pn].(Package)
	//}
	//
	//n := strings.LastIndex(thisPath, pn)
	//if n > 0 {
	//	y := strings.Index(thisPath[:n+1], pn)
	//	if y != -1 {
	//		panic("multiple definition")
	//	}
	//}
	//
	//if sp,ok := p[pn]; ok {
	//	return sp.(Package)
	//}
	//
	//if n == -1 {
	//	panic("not found")
	//}
	//return p[thisPath[:len(pn)+n]].(Package)
}

//func consume(p PackageSetOrPackage, paths... string) PackageSetOrPackage {
//	var dotIndex, lastDotIndex int
//	if p == nil {
//		return p
//	}
//	for _, path := range paths {
//		dotIndex, lastDotIndex = 0, 0
//		for dotIndex != -1 && dotIndex < len(path) {
//			dotIndex = strings.IndexByte(path[dotIndex:], '.')
//
//			p = p.Get(path[lastDotIndex:dotIndex])
//			if p == nil {
//				return p
//			}
//			lastDotIndex = dotIndex + 1
//		}
//	}
//	return p
//}

func findCheckFunc(p LinkedContext, k string) CheckFunc {
	var ls = strings.LastIndexByte(k, '.')
	if ls == -1 {
		return p.GetFunc(k)
	}
	pn, fn := k[:ls], k[ls+1:]
	if len(pn) == 0 {
		return p.GetFunc(fn)
	}

	sp := RepositionCtx(p, pn)
	if sp == nil {
		sp = RepositionCtx(p, dotJoin("root", pn))
	}
	if sp != nil {
		return sp.GetFunc(fn)
	}
	return nil
}

var namespaceStd = Package{
	"Assert": func(req Request, s2 ...string) (s string, err error) {
		ensureVarLength(s2, 2, &err)
		if body := ensureJSONBody(&req, &err); body != nil {
			fmt.Println("asserting", body[s2[0]], s2[1])
		}
		return
	},
}

var namespaceJSON = Package{
	"Assert": func(req Request, s2 ...string) (s string, err error) {
		ensureVarLength(s2, 2, &err)
		if body := ensureJSONBody(&req, &err); body != nil {
			fmt.Println("asserting", body[s2[0]], s2[1])
		}
		return
	},
}
