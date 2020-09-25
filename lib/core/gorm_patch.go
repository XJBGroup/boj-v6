package mcore

import (
	"github.com/jinzhu/gorm"
	"reflect"
	"unsafe"
)

func changeableField(scope *gorm.Scope, field *gorm.Field) bool {
	if selectAttrs := scope.SelectAttrs(); len(selectAttrs) > 0 {
		for _, attr := range selectAttrs {
			if field.Name == attr || field.DBName == attr {
				return true
			}
		}
		return false
	}

	for _, attr := range scope.OmitAttrs() {
		if field.Name == attr || field.DBName == attr {
			return false
		}
	}

	return true
}

type searchable interface {
	Search(searching string) bool
	Len() int
}

type serialSearchContainer []string

func (s serialSearchContainer) Len() int {
	return len(s)
}

func (s serialSearchContainer) Search(searching string) bool {
	for i := range s {
		if s[i] == searching {
			return true
		}
	}
	return false
}

type mapSearchContainer map[string]bool

func (s mapSearchContainer) Len() int {
	return len(s)
}

func (s mapSearchContainer) Search(searching string) bool {
	_, ok := s[searching]
	return ok
}

func getFields(scope *gorm.Scope) []*gorm.Field {
	var (
		fields             []*gorm.Field
		indirectScopeValue = scope.IndirectValue()
		isStruct           = indirectScopeValue.Kind() == reflect.Struct
	)

	// offset = 8 + 24 * 7 = 176
	search := *(*map[string]interface{})(
		unsafe.Pointer(uintptr(unsafe.Pointer(scope.Search)) + 176))
	var queries searchable
	if search != nil {
		var v = search["query"]
		if v != nil {
			queries = serialSearchContainer(v.([]string))

			if queries.Len() > 10 {
				var oq = make(map[string]bool)
				for _, vv := range v.([]string) {
					oq[vv] = true
				}
				queries = mapSearchContainer(oq)
			}
		}
	}

	if isStruct {
		for _, structField := range scope.GetModelStruct().StructFields {
			fieldValue := indirectScopeValue
			for _, name := range structField.Names {
				if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
					fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
				}
				fieldValue = reflect.Indirect(fieldValue).FieldByName(name)
			}

			if queries == nil || queries.Search(structField.DBName) {
				fields = append(fields, &gorm.Field{StructField: structField, Field: fieldValue})
			}
		}
	} else {
		for _, structField := range scope.GetModelStruct().StructFields {
			if queries == nil || queries.Search(structField.DBName) {
				fields = append(fields, &gorm.Field{StructField: structField, IsBlank: true})
			}
		}
	}
	return fields
}

func convertInterfaceToMap(values interface{}, withIgnoredField bool, scope *gorm.Scope) map[string]interface{} {
	var attrs = map[string]interface{}{}

	switch value := values.(type) {
	case map[string]interface{}:
		return value
	case []interface{}:
		for _, v := range value {
			for key, value := range convertInterfaceToMap(v, withIgnoredField, scope) {
				attrs[key] = value
			}
		}
	case interface{}:
		reflectValue := reflect.ValueOf(values)

		switch reflectValue.Kind() {
		case reflect.Map:
			for _, key := range reflectValue.MapKeys() {
				attrs[gorm.ToColumnName(key.Interface().(string))] = reflectValue.MapIndex(key).Interface()
			}
		default:
			for _, field := range getFields(scope) {
				if !field.IsBlank && (withIgnoredField || !field.IsIgnored) {
					attrs[field.DBName] = field.Field.Interface()
				}
			}
		}
	}
	return attrs
}

func updatedAttrsWithValues(scope *gorm.Scope, value interface{}) (results map[string]interface{}, hasUpdate bool) {
	if scope.IndirectValue().Kind() != reflect.Struct {
		return convertInterfaceToMap(value, false, scope), true
	}

	results = map[string]interface{}{}

	for key, value := range convertInterfaceToMap(value, true, scope) {
		if field, ok := scope.FieldByName(key); ok {
			if changeableField(scope, field) {
				if _, ok := value.(*gorm.SqlExpr); ok {
					hasUpdate = true
					results[field.DBName] = value
				} else {
					err := field.Set(value)
					if field.IsNormal && !field.IsIgnored {
						hasUpdate = true
						if err == gorm.ErrUnaddressable {
							results[field.DBName] = value
						} else {
							results[field.DBName] = field.Field.Interface()
						}
					}
				}
			}
		} else {
			results[key] = value
		}
	}
	return
}

func init() {
	// hack terrible assign_updating_attributes impl of callback
	gorm.DefaultCallback.Update().Remove("gorm:assign_updating_attributes")
	gorm.DefaultCallback.Update().Before("after_create1").Register(
		"gorm:assign_updating_attributes", func(scope *gorm.Scope) {
			if attrs, ok := scope.InstanceGet("gorm:update_interface"); ok {
				if updateMaps, hasUpdate := updatedAttrsWithValues(scope, attrs); hasUpdate {
					scope.InstanceSet("gorm:update_attrs", updateMaps)
				} else {
					scope.SkipLeft()
				}
			}
		})
}
