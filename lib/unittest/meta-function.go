package unittest

import "errors"

type MetaStorage = map[string]interface{}

type MetaParser interface {
	GetTargetProperty() string
	ParseMeta(v interface{}) (parsedValue interface{}, err error)
}

type MetaOperation interface {
	ZeroValue(t MetaStorage) bool
	AssignDefault(v interface{}, t MetaStorage) error
}

type StringPropertyParser struct {
	metaKey string
}

func (p StringPropertyParser) GetTargetProperty() string {
	return p.metaKey
}

func (p StringPropertyParser) ParseMeta(v interface{}) (parsedValue interface{}, err error) {
	switch e := v.(type) {
	case string:
		return e, nil
	case nil:
		return "", nil
	default:
		panic("encoding type error")
	}
}

func ParseStringProperty(metaKey string) MetaParser {
	return StringPropertyParser{metaKey: metaKey}
}

type StringMapPropertyParser struct {
	metaKey string
}

func (p StringMapPropertyParser) GetTargetProperty() string {
	return p.metaKey
}

func (p StringMapPropertyParser) ParseMeta(v interface{}) (parsedValue interface{}, err error) {
	var nv map[string]string
	switch h := v.(type) {
	case map[string]string:
		nv = h
		return nv, nil
	case map[string]interface{}:
		nv = make(map[string]string)
		for k, v := range h {
			nv[k] = v.(string)
		}
		return nv, nil
	case map[interface{}]interface{}:
		nv = make(map[string]string)
		for k, v := range h {
			nv[k.(string)] = v.(string)
		}
		return nv, nil
	case nil:
		return
	default:
		return nil, errors.New("string map type error")
	}
}

func ParseStringMapProperty(metaKey string) MetaParser {
	return StringMapPropertyParser{metaKey: metaKey}
}

type DataBodyPropertyParser struct {
	metaKey string
}

func (p DataBodyPropertyParser) GetTargetProperty() string {
	return p.metaKey
}

func (p DataBodyPropertyParser) ParseMeta(v interface{}) (parsedValue interface{}, err error) {
	switch d := v.(type) {
	case map[string]interface{}:
		return toDataBody(d).(map[string]interface{}), nil
	case map[interface{}]interface{}:
		return toDataBody(d).(map[string]interface{}), nil
	case nil:
		return nil, nil
	default:
		return nil, errors.New("data map type error")
	}
}

func ParseDataBodyProperty(metaKey string) MetaParser {
	return DataBodyPropertyParser{metaKey: metaKey}
}

type StringMetaOperation struct {
	metaKey string
}

func (o StringMetaOperation) ZeroValue(t MetaStorage) bool {
	if tv, ok := t[o.metaKey]; !ok || len(tv.(string)) == 0 {
		return true
	} else {
		return false
	}
}

func (o StringMetaOperation) AssignDefault(v interface{}, t MetaStorage) error {
	if v != nil && len(v.(string)) != 0 {
		if o.ZeroValue(t) {
			t[o.metaKey] = v.(string)
		}
	}
	return nil
}

func StringMetaPropertyOperation(metaKey string) MetaOperation {
	return StringMetaOperation{metaKey: metaKey}
}

type DataBodyMetaOperation struct {
	metaKey string
}

func (o DataBodyMetaOperation) ZeroValue(t MetaStorage) bool {
	if tv, ok := t[o.metaKey]; !ok || len(tv.(DataBody)) == 0 {
		return true
	} else {
		return false
	}
}

func (o DataBodyMetaOperation) AssignDefault(v interface{}, t MetaStorage) error {
	if v != nil && len(v.(DataBody)) != 0 {
		if t[o.metaKey] == nil {
			t[o.metaKey] = make(DataBody)
		}
		inheritMapType(t[o.metaKey].(DataBody), v.(DataBody))
	}
	return nil
}

func DataBodyMetaPropertyOperation(metaKey string) MetaOperation {
	return DataBodyMetaOperation{metaKey: metaKey}
}

type StringMap = map[string]string
type StringMapMetaOperation struct {
	metaKey string
}

func (o StringMapMetaOperation) ZeroValue(t MetaStorage) bool {
	if tv, ok := t[o.metaKey]; !ok || len(tv.(StringMap)) == 0 {
		return true
	} else {
		return false
	}
}

func (o StringMapMetaOperation) AssignDefault(v interface{}, t MetaStorage) error {
	if v != nil && len(v.(StringMap)) != 0 {
		if t[o.metaKey] == nil {
			t[o.metaKey] = make(map[string]string)
		}
		dst := t[o.metaKey].(map[string]string)
		for k, v := range v.(StringMap) {
			if _, ok := dst[k]; !ok {
				dst[k] = v
			}
		}
	}
	return nil
}

func StringMapMetaPropertyOperation(metaKey string) MetaOperation {
	return StringMapMetaOperation{metaKey: metaKey}
}
