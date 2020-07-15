package unittest

import (
	"errors"
	"strings"
)

type Option struct {
	metaOperationMap map[string]MetaOperation
	ParseMetaMap     map[string]MetaParser
}

type ErrMetaParserNotFound struct {
	K string
}

func (e ErrMetaParserNotFound) Error() string {
	return "parse function of " + e.K + " not found"
}

type caseContext struct {
	warnings          []error
	Packages          LinkedContext
	gd                *GoDynamicTestData
	testCasePath      string
	inheritFunction   func(k string, v interface{}, t MetaStorage) error
	parseMetaFunction func(k string, v interface{}) (propertyName string, parsedValue interface{}, err error)
	mt                map[string]Matcher
}

func newContext(opt *Option) (*caseContext, error) {
	metaOperationMap := opt.metaOperationMap
	if metaOperationMap == nil {
		metaOperationMap = make(map[string]MetaOperation)
	}
	parseMetaMap := opt.ParseMetaMap
	if parseMetaMap == nil {
		parseMetaMap = make(map[string]MetaParser)
	}

	// todo: check parseMeta propertyName not conflict

	return &caseContext{
		Packages: &linkedContext{
			name: "root",
		},
		inheritFunction: func(k string, v interface{}, t MetaStorage) error {
			fn, ok := metaOperationMap[k]
			if !ok {
				return errors.New("inherit function of " + k + " not found")
			}
			return fn.AssignDefault(v, t)
		},
		parseMetaFunction: func(k string, v interface{}) (string, interface{}, error) {
			fn, ok := parseMetaMap[k]
			if !ok {
				return "", nil, ErrMetaParserNotFound{k}
			}
			pv, err := fn.ParseMeta(v)
			return fn.GetTargetProperty(), pv, err
		},
		gd: new(GoDynamicTestData),
	}, nil
}

func (c *caseContext) inheritPropertyKV(k string, v interface{}, t MetaStorage) error {
	if c.inheritFunction == nil {
		return errors.New("inherit function not registered")
	}
	return c.inheritFunction(k, v, t)
}

func (c *caseContext) inheritProperty(dst MetaStorage, src MetaStorage) (err error) {
	if c.inheritFunction == nil {
		return errors.New("inherit function not registered")
	}
	if src == nil {
		return
	}
	for k, v := range src {
		err = c.inheritFunction(k, v, dst)
		if err != nil {
			return
		}
	}
	return
}

func (c *caseContext) parseMetaKV(k string, v interface{}) (string, interface{}, error) {
	if c.parseMetaFunction == nil {
		return "", nil, errors.New("parse function not registered")
	}
	return c.parseMetaFunction(k, v)
}

func (c *caseContext) parseMeta(meta map[string]interface{}, t *TestCase) error {
	if c.parseMetaFunction == nil {
		return errors.New("parse function not registered")
	}
	if t.Meta == nil {
		t.Meta = make(map[string]interface{})
	}
	for k, v := range meta {
		p, pv, err := c.parseMetaFunction(k, v)
		if err != nil {
			switch warning := err.(type) {
			case ErrMetaParserNotFound:
				c.warnings = append(c.warnings, warning)
			default:
				return err
			}
		} else {
			t.Meta[p] = pv
		}
	}
	return nil
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
