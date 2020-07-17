package unittest

var inheritFnMap = map[string]MetaOperation{
	MetaEncoding:     StringMetaPropertyOperation(MetaEncoding),
	MetaMethod:       StringMetaPropertyOperation(MetaMethod),
	MetaUrl:          StringMetaPropertyOperation(MetaUrl),
	MetaData:         DataBodyMetaPropertyOperation(MetaData),
	MetaHeader:       StringMapMetaPropertyOperation(MetaHeader),
	MetaHTTPEncoding: StringMetaPropertyOperation(MetaHTTPEncoding),
	MetaHTTPMethod:   StringMetaPropertyOperation(MetaHTTPMethod),
	MetaHTTPHeader:   StringMapMetaPropertyOperation(MetaHTTPHeader),
}

var parseMetaFnMap = map[string]MetaParser{
	"encoding":      ParseStringProperty(MetaEncoding),
	"method":        ParseStringProperty(MetaMethod),
	"url":           ParseStringProperty(MetaUrl),
	"data":          ParseDataBodyProperty(MetaData),
	"header":        ParseStringMapProperty(MetaHeader),
	"http-encoding": ParseStringProperty(MetaHTTPEncoding),
	"http-method":   ParseStringProperty(MetaHTTPMethod),
	"http-header":   ParseStringMapProperty(MetaHTTPHeader),
}

var V1Opt = &Option{MetaOperationMap: inheritFnMap, ParseMetaMap: parseMetaFnMap}
