package unittest

import (
	"errors"
	"regexp"
	"strings"
)

type Matcher interface {
	Match(t *TestCase) (bool, error)
}

type AndMatcher struct {
	matchers []Matcher
}

func (a AndMatcher) Match(t *TestCase) (v bool, err error) {
	if len(a.matchers) == 0 {
		panic("nil and")
	}
	for i := range a.matchers {
		if v, err = a.matchers[i].Match(t); err != nil || v == false {
			return false, err
		}
	}
	return true, nil
}

func newSelector(path string, c map[string]interface{}) (g Matcher, err error) {
	var m AndMatcher
	for k, v := range c {
		switch v := v.(type) {
		case string:
			g, err = newStringSelector(path, v)
		case map[string]interface{}:
			g, err = newSelector(dotJoin(path, k), v)
		case map[interface{}]interface{}:
			var sc = map[string]interface{}{}
			for k, vv := range v {
				sc[k.(string)] = vv
			}
			g, err = newSelector(dotJoin(path, k), sc)
		}
		if err != nil {
			return
		}
		m.matchers = append(m.matchers, g)
	}
	return
}

type NotExistsMatcher struct {
	path string
}

func (n NotExistsMatcher) MatchString(s string) (bool, error) {
	panic("implement me")
}

func (n NotExistsMatcher) Match(t *TestCase) (bool, error) {
	panic("implement me")
}

type StringMatcher interface {
	Matcher
	MatchString(s string) (bool, error)
}

type RegexpMatcher struct {
	path string
	*regexp.Regexp
}

func (n RegexpMatcher) MatchString(s string) (bool, error) {
	panic("implement me")
}

func (n RegexpMatcher) Match(t *TestCase) (bool, error) {
	panic("implement me")
}

type StringLiteralMatcher struct {
	path string
	v    string
}

func (n StringLiteralMatcher) MatchString(s string) (bool, error) {
	return n.v == s, nil
}

func (n StringLiteralMatcher) Match(t *TestCase) (bool, error) {
	panic("implement me")
}

func newStringSelector(path, v string) (StringMatcher, error) {
	if len(v) == 0 {
		return NotExistsMatcher{path}, nil
	}
	if v[0] == '/' {
		if len(v) > 1 && v[len(v)-1] == '/' {
			return RegexpMatcher{path: path, Regexp: regexp.MustCompile(v[1 : len(v)-1])}, nil
		} else {
			return nil, errors.New("invalid regexp selector: " + v)
		}
	}
	return StringLiteralMatcher{path: path, v: v}, nil
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
	methodField, ok := t.Meta[MetaMethod]
	if !ok || methodField == nil {
		return false, nil
	}
	x, y := m.StringMatcher.MatchString(methodField.(string))
	return !x, y
}

func newMethodNEQMatcher(m StringMatcher) Matcher {
	return MethodNEQMatcher{m}
}
