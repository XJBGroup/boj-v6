package unittest_statics

import (
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest/inner"
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_types"
	"strings"
)

type LinkedContext struct {
	last unittest_types.LinkedContext
	name string
	next map[string]unittest_types.LinkedContext
	fns  unittest_types.Package
}

func NewFunctionPackage(name string, fns unittest_types.Package) *LinkedContext {
	return &LinkedContext{name: name, fns: fns}
}

func CopyLink(l unittest_types.LinkedContext) *CopiedLinkedContext {
	return &CopiedLinkedContext{attached: l}
}

type CopiedLinkedContext struct {
	last     unittest_types.LinkedContext
	name     string
	attached unittest_types.LinkedContext
}

func (c CopiedLinkedContext) Last() unittest_types.LinkedContext {
	return c.last
}

func (c CopiedLinkedContext) Name() string {
	return c.name
}

func (c CopiedLinkedContext) Get(s string) unittest_types.LinkedContext {
	return c.attached.Get(s)
}

func (c CopiedLinkedContext) GetFunc(s string) unittest_types.CheckFunc {
	return c.attached.GetFunc(s)
}

func (c CopiedLinkedContext) GetFunctions(mf func(s string, f unittest_types.CheckFunc) error) error {
	return c.attached.GetFunctions(mf)
}

func (c *CopiedLinkedContext) SetName(s string) {
	c.name = s
}

func (c *CopiedLinkedContext) SetLast(l unittest_types.LinkedContext) {
	c.last = l
}

func (c CopiedLinkedContext) Insert(string, unittest_types.LinkedContext) {
	panic("not changeable")
}

func (c CopiedLinkedContext) InsertFunc(string, unittest_types.CheckFunc) {
	panic("not changeable")
}
func (c LinkedContext) Name() string {
	return c.name
}

func (c *LinkedContext) SetName(s string) {
	c.name = s
}

func (c LinkedContext) Last() unittest_types.LinkedContext {
	return c.last
}

func (c LinkedContext) Get(s string) unittest_types.LinkedContext {
	if c.next == nil {
		return nil
	}
	return c.next[s]
}

func (c LinkedContext) GetFunc(s string) unittest_types.CheckFunc {
	if c.fns == nil {
		return nil
	}
	return c.fns[s]
}

func (c LinkedContext) GetFunctions(mf func(s string, f unittest_types.CheckFunc) error) (err error) {
	for k, v := range c.fns {
		err = mf(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *LinkedContext) SetLast(ctx unittest_types.LinkedContext) {
	c.last = ctx
}

func (c *LinkedContext) Insert(k string, ctx unittest_types.LinkedContext) {
	if len(k) == 0 {
		if c.fns == nil {
			c.fns = make(map[string]unittest_types.CheckFunc)
		}
		err := ctx.GetFunctions(func(k string, f unittest_types.CheckFunc) error {
			if _, ok := c.fns[k]; ok {
				// todo conflict
				// panic("conflict")
			} else {
				c.fns[k] = f
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
		return
	}

	if c.next == nil {
		c.next = make(map[string]unittest_types.LinkedContext)
	}
	if _, ok := c.next[k]; ok {
		panic("conflict")
	}
	c.next[k] = ctx
	ctx.SetLast(c)
}

func (c *LinkedContext) InsertFunc(k string, f unittest_types.CheckFunc) {
	if c.fns == nil {
		c.fns = make(map[string]unittest_types.CheckFunc)
	}
	if _, ok := c.fns[k]; ok {
		panic("conflict")
	}
	c.fns[k] = f
}

func RepositionCtx(p unittest_types.LinkedContext, pn string) unittest_types.LinkedContext {
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
}

func FindCheckFunc(p unittest_types.LinkedContext, k string) unittest_types.CheckFunc {
	var ls = strings.LastIndexByte(k, '.')
	if ls == -1 {
		return GetFuncRecursive(p, k)
	}
	pn, fn := k[:ls], k[ls+1:]
	if len(pn) == 0 {
		return GetFuncRecursive(p, fn)
	}

	sp := RepositionCtx(p, pn)
	if sp == nil {
		sp = RepositionCtx(p, inner.DotJoin("root", pn))
	}
	if sp != nil {
		return sp.GetFunc(fn)
	}
	return nil
}

func GetFuncRecursive(p unittest_types.LinkedContext, k string) (f unittest_types.CheckFunc) {
	for f == nil {
		if p == nil {
			return
		}
		f = p.GetFunc(k)
		p = p.Last()
	}
	return
}
