package unittest

import (
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_statics"
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_types"
)

var inheritFnMap = map[string]unittest_types.MetaOperation{
	MetaEncoding:     unittest_statics.StringMetaPropertyOperation(MetaEncoding),
	MetaMethod:       unittest_statics.StringMetaPropertyOperation(MetaMethod),
	MetaUrl:          unittest_statics.StringMetaPropertyOperation(MetaUrl),
	MetaData:         unittest_statics.DataBodyMetaPropertyOperation(MetaData),
	MetaHeader:       unittest_statics.StringMapMetaPropertyOperation(MetaHeader),
	MetaHTTPEncoding: unittest_statics.StringMetaPropertyOperation(MetaHTTPEncoding),
	MetaHTTPMethod:   unittest_statics.StringMetaPropertyOperation(MetaHTTPMethod),
	MetaHTTPHeader:   unittest_statics.StringMapMetaPropertyOperation(MetaHTTPHeader),
}

var parseMetaFnMap = map[string]unittest_types.MetaParser{
	"encoding":      unittest_statics.ParseStringProperty(MetaEncoding),
	"method":        unittest_statics.ParseStringProperty(MetaMethod),
	"url":           unittest_statics.ParseStringProperty(MetaUrl),
	"data":          unittest_statics.ParseDataBodyProperty(MetaData),
	"header":        unittest_statics.ParseStringMapProperty(MetaHeader),
	"http-encoding": unittest_statics.ParseStringProperty(MetaHTTPEncoding),
	"http-method":   unittest_statics.ParseStringProperty(MetaHTTPMethod),
	"http-header":   unittest_statics.ParseStringMapProperty(MetaHTTPHeader),
}

var V1Opt = &Option{MetaOperationMap: inheritFnMap, ParseMetaMap: parseMetaFnMap}
