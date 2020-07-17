package unittest_types

type MetaStorage = map[string]interface{}

type MetaParser interface {
	GetTargetProperty() string
	ParseMeta(v interface{}) (parsedValue interface{}, err error)
}

type MetaOperation interface {
	ZeroValue(t MetaStorage) bool
	AssignDefault(v interface{}, t MetaStorage) error
}
