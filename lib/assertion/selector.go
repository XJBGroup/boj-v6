package assertion

import (
	"errors"
	"regexp"
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
