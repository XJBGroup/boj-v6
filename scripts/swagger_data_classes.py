from typing import *


class Schema(object):
    id: Union[None, 'str']
    ref: Union[None, 'str']
    schemaUrl: Union[None, 'str']
    description: Union[None, 'str']
    type: Union[None, 'Union[str, List[str]]']
    nullable: Union[None, 'bool']
    format: Union[None, 'str']
    title: Union[None, 'str']
    default: Union[None, 'object']
    maximum: Union[None, 'float']
    exclusiveMaximum: Union[None, 'bool']
    minimum: Union[None, 'float']
    exclusiveMinimum: Union[None, 'bool']
    maxLength: Union[None, 'int']
    minLength: Union[None, 'int']
    pattern: Union[None, 'str']
    maxItems: Union[None, 'int']
    minItems: Union[None, 'int']
    uniqueItems: Union[None, 'bool']
    multipleOf: Union[None, 'float']
    enum: Union[None, 'List[object]']
    maxProperties: Union[None, 'int']
    minProperties: Union[None, 'int']
    required: Union[None, 'List[str]']
    items: Union[None, 'Union[Schema, List[Schema]]']
    allOf: Union[None, 'List[Schema]']
    oneOf: Union[None, 'List[Schema]']
    anyOf: Union[None, 'List[Schema]']
    not_: Union[None, 'Schema']
    properties: Union[None, 'Dict[str, Schema]']
    additionalProperties: Union[None, 'Union[Schema, bool]']
    patternProperties: Union[None, 'Dict[str, Schema]']
    dependencies: Union[None, 'Dict[str, Union[Schema, List[str]]]']
    additionalItems: Union[None, 'Union[Schema, bool]']
    definitions: Union[None, 'Dict[str, Schema]']

    def __init__(self):
        self.id = None
        self.ref = None
        self.schemaUrl = None
        self.description = None
        self.type = None
        self.nullable = None
        self.format = None
        self.title = None
        self.default = None
        self.maximum = None
        self.exclusiveMaximum = None
        self.minimum = None
        self.exclusiveMinimum = None
        self.maxLength = None
        self.minLength = None
        self.pattern = None
        self.maxItems = None
        self.minItems = None
        self.uniqueItems = None
        self.multipleOf = None
        self.enum = None
        self.maxProperties = None
        self.minProperties = None
        self.required = None
        self.items = None
        self.allOf = None
        self.oneOf = None
        self.anyOf = None
        self.not_ = None
        self.properties = None
        self.additionalProperties = None
        self.patternProperties = None
        self.dependencies = None
        self.additionalItems = None
        self.definitions = None

    @staticmethod
    def from_dict(d: dict) -> 'Schema':
        ci = Schema()
        if d is None:
            return ci
        ci.id = getting = d.get('id', None)
        if getting is not None:
            ci.id = str(getting)
        ci.ref = getting = d.get('ref', None)
        if getting is not None:
            ci.ref = str(getting)
        ci.schemaUrl = getting = d.get('schemaUrl', None)
        if getting is not None:
            ci.schemaUrl = str(getting)
        ci.description = getting = d.get('description', None)
        if getting is not None:
            ci.description = str(getting)
        ci.type = getting = d.get('type', None)
        if getting is not None:
            g_a_cpc_x = None
            g_a_cpc_x_6497c00b_handling = getting
            if isinstance(g_a_cpc_x_6497c00b_handling, list):
                g_a_cpc_x_3ae62090 = list()
                for g_a_cpc_x_3ae62090_value in g_a_cpc_x_6497c00b_handling:
                    g_a_cpc_x = str(g_a_cpc_x_3ae62090_value)
                    g_a_cpc_x_3ae62090.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_3ae62090
            if g_a_cpc_x is None:
                g_a_cpc_x = str(g_a_cpc_x_6497c00b_handling)
            ci.type = g_a_cpc_x
        ci.nullable = getting = d.get('nullable', None)
        if getting is not None:
            ci.nullable = bool(getting)
        ci.format = getting = d.get('format', None)
        if getting is not None:
            ci.format = str(getting)
        ci.title = getting = d.get('title', None)
        if getting is not None:
            ci.title = str(getting)
        ci.default = getting = d.get('default', None)
        if getting is not None:
            ci.default = object.from_dict(getting)
        ci.maximum = getting = d.get('maximum', None)
        if getting is not None:
            ci.maximum = float(getting)
        ci.exclusiveMaximum = getting = d.get('exclusiveMaximum', None)
        if getting is not None:
            ci.exclusiveMaximum = bool(getting)
        ci.minimum = getting = d.get('minimum', None)
        if getting is not None:
            ci.minimum = float(getting)
        ci.exclusiveMinimum = getting = d.get('exclusiveMinimum', None)
        if getting is not None:
            ci.exclusiveMinimum = bool(getting)
        ci.maxLength = getting = d.get('maxLength', None)
        if getting is not None:
            ci.maxLength = int(getting)
        ci.minLength = getting = d.get('minLength', None)
        if getting is not None:
            ci.minLength = int(getting)
        ci.pattern = getting = d.get('pattern', None)
        if getting is not None:
            ci.pattern = str(getting)
        ci.maxItems = getting = d.get('maxItems', None)
        if getting is not None:
            ci.maxItems = int(getting)
        ci.minItems = getting = d.get('minItems', None)
        if getting is not None:
            ci.minItems = int(getting)
        ci.uniqueItems = getting = d.get('uniqueItems', None)
        if getting is not None:
            ci.uniqueItems = bool(getting)
        ci.multipleOf = getting = d.get('multipleOf', None)
        if getting is not None:
            ci.multipleOf = float(getting)
        ci.enum = getting = d.get('enum', None)
        if getting is not None:
            if isinstance(getting, list):
                g_a_cpc_x_5f7abdb1 = list()
                for g_a_cpc_x_5f7abdb1_value in getting:
                    g_a_cpc_x = object.from_dict(g_a_cpc_x_5f7abdb1_value)
                    g_a_cpc_x_5f7abdb1.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_5f7abdb1
            ci.enum = g_a_cpc_x
        ci.maxProperties = getting = d.get('maxProperties', None)
        if getting is not None:
            ci.maxProperties = int(getting)
        ci.minProperties = getting = d.get('minProperties', None)
        if getting is not None:
            ci.minProperties = int(getting)
        ci.required = getting = d.get('required', None)
        if getting is not None:
            if isinstance(getting, list):
                g_a_cpc_x_163f4527 = list()
                for g_a_cpc_x_163f4527_value in getting:
                    g_a_cpc_x = str(g_a_cpc_x_163f4527_value)
                    g_a_cpc_x_163f4527.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_163f4527
            ci.required = g_a_cpc_x
        ci.items = getting = d.get('items', None)
        if getting is not None:
            g_a_cpc_x = None
            g_a_cpc_x_4ca3191a_handling = getting
            if isinstance(g_a_cpc_x_4ca3191a_handling, list):
                g_a_cpc_x_1751679e = list()
                for g_a_cpc_x_1751679e_value in g_a_cpc_x_4ca3191a_handling:
                    g_a_cpc_x = Schema.from_dict(g_a_cpc_x_1751679e_value)
                    g_a_cpc_x_1751679e.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_1751679e
            if g_a_cpc_x is None:
                g_a_cpc_x = Schema.from_dict(g_a_cpc_x_4ca3191a_handling)
            ci.items = g_a_cpc_x
        ci.allOf = getting = d.get('allOf', None)
        if getting is not None:
            if isinstance(getting, list):
                g_a_cpc_x_6f578f34 = list()
                for g_a_cpc_x_6f578f34_value in getting:
                    g_a_cpc_x = Schema.from_dict(g_a_cpc_x_6f578f34_value)
                    g_a_cpc_x_6f578f34.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_6f578f34
            ci.allOf = g_a_cpc_x
        ci.oneOf = getting = d.get('oneOf', None)
        if getting is not None:
            if isinstance(getting, list):
                g_a_cpc_x_5c9005fd = list()
                for g_a_cpc_x_5c9005fd_value in getting:
                    g_a_cpc_x = Schema.from_dict(g_a_cpc_x_5c9005fd_value)
                    g_a_cpc_x_5c9005fd.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_5c9005fd
            ci.oneOf = g_a_cpc_x
        ci.anyOf = getting = d.get('anyOf', None)
        if getting is not None:
            if isinstance(getting, list):
                g_a_cpc_x_72f830d4 = list()
                for g_a_cpc_x_72f830d4_value in getting:
                    g_a_cpc_x = Schema.from_dict(g_a_cpc_x_72f830d4_value)
                    g_a_cpc_x_72f830d4.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_72f830d4
            ci.anyOf = g_a_cpc_x
        ci.not_ = getting = d.get('not', None)
        if getting is not None:
            ci.not_ = Schema.from_dict(getting)
        ci.properties = getting = d.get('properties', None)
        if getting is not None:
            if isinstance(getting, dict):
                g_a_cpc_x_2136d831 = dict()
                for g_a_cpc_x_2136d831_value in getting.items():
                    g_a_cpc_x = Schema.from_dict(g_a_cpc_x_2136d831_value[1])
                    g_a_cpc_x_2136d831[g_a_cpc_x_2136d831_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_2136d831
            ci.properties = g_a_cpc_x
        ci.additionalProperties = getting = d.get('additionalProperties', None)
        if getting is not None:
            g_a_cpc_x = None
            g_a_cpc_x_73b05e91_handling = getting
            g_a_cpc_x = Schema.from_dict(g_a_cpc_x_73b05e91_handling)
            if g_a_cpc_x is None:
                g_a_cpc_x = bool(g_a_cpc_x_73b05e91_handling)
            ci.additionalProperties = g_a_cpc_x
        ci.patternProperties = getting = d.get('patternProperties', None)
        if getting is not None:
            if isinstance(getting, dict):
                g_a_cpc_x_b0c69cc = dict()
                for g_a_cpc_x_b0c69cc_value in getting.items():
                    g_a_cpc_x = Schema.from_dict(g_a_cpc_x_b0c69cc_value[1])
                    g_a_cpc_x_b0c69cc[g_a_cpc_x_b0c69cc_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_b0c69cc
            ci.patternProperties = g_a_cpc_x
        ci.dependencies = getting = d.get('dependencies', None)
        if getting is not None:
            if isinstance(getting, dict):
                g_a_cpc_x_2e7f0311 = dict()
                for g_a_cpc_x_2e7f0311_value in getting.items():
                    g_a_cpc_x = None
                    g_a_cpc_x_4446bac_handling = g_a_cpc_x_2e7f0311_value[1]
                    if isinstance(g_a_cpc_x_4446bac_handling, list):
                        g_a_cpc_x_39abcfc3 = list()
                        for g_a_cpc_x_39abcfc3_value in g_a_cpc_x_4446bac_handling:
                            g_a_cpc_x = str(g_a_cpc_x_39abcfc3_value)
                            g_a_cpc_x_39abcfc3.append(g_a_cpc_x)
                        g_a_cpc_x = g_a_cpc_x_39abcfc3
                    if g_a_cpc_x is None:
                        g_a_cpc_x = Schema.from_dict(g_a_cpc_x_4446bac_handling)
                    g_a_cpc_x_2e7f0311[g_a_cpc_x_2e7f0311_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_2e7f0311
            ci.dependencies = g_a_cpc_x
        ci.additionalItems = getting = d.get('additionalItems', None)
        if getting is not None:
            g_a_cpc_x = None
            g_a_cpc_x_1015246f_handling = getting
            g_a_cpc_x = Schema.from_dict(g_a_cpc_x_1015246f_handling)
            if g_a_cpc_x is None:
                g_a_cpc_x = bool(g_a_cpc_x_1015246f_handling)
            ci.additionalItems = g_a_cpc_x
        ci.definitions = getting = d.get('definitions', None)
        if getting is not None:
            if isinstance(getting, dict):
                g_a_cpc_x_62b0073a = dict()
                for g_a_cpc_x_62b0073a_value in getting.items():
                    g_a_cpc_x = Schema.from_dict(g_a_cpc_x_62b0073a_value[1])
                    g_a_cpc_x_62b0073a[g_a_cpc_x_62b0073a_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_62b0073a
            ci.definitions = g_a_cpc_x
        return ci


class ResponseDesc(object):
    description: Union[None, 'str']
    schema: Union[None, 'Schema']

    def __init__(self):
        self.description = None
        self.schema = None

    @staticmethod
    def from_dict(d: dict) -> 'ResponseDesc':
        ci = ResponseDesc()
        if d is None:
            return ci
        ci.description = getting = d.get('description', None)
        if getting is not None:
            ci.description = str(getting)
        ci.schema = getting = d.get('schema', None)
        if getting is not None:
            ci.schema = Schema.from_dict(getting)
        return ci


class Parameter(object):
    description: Union[None, 'str']
    name: Union[None, 'str']
    in_: Union[None, 'str']
    required: Union[None, 'bool']
    schema: Union[None, 'Schema']
    allowEmptyValue: Union[None, 'bool']

    def __init__(self):
        self.description = None
        self.name = None
        self.in_ = None
        self.required = None
        self.schema = None
        self.allowEmptyValue = None

    @staticmethod
    def from_dict(d: dict) -> 'Parameter':
        ci = Parameter()
        if d is None:
            return ci
        ci.description = getting = d.get('description', None)
        if getting is not None:
            ci.description = str(getting)
        ci.name = getting = d.get('name', None)
        if getting is not None:
            ci.name = str(getting)
        ci.in_ = getting = d.get('in', None)
        if getting is not None:
            ci.in_ = str(getting)
        ci.required = getting = d.get('required', None)
        if getting is not None:
            ci.required = bool(getting)
        ci.schema = getting = d.get('schema', None)
        if getting is not None:
            ci.schema = Schema.from_dict(getting)
        ci.allowEmptyValue = getting = d.get('allowEmptyValue', None)
        if getting is not None:
            ci.allowEmptyValue = bool(getting)
        return ci


class PathItem(object):
    consumes: Union[None, 'List[str]']
    produces: Union[None, 'List[str]']
    tags: Union[None, 'List[str]']
    operationId: Union[None, 'str']
    parameters: Union[None, 'List[Parameter]']
    responses: Union[None, 'Dict[str, ResponseDesc]']

    def __init__(self):
        self.consumes = None
        self.produces = None
        self.tags = None
        self.operationId = None
        self.parameters = None
        self.responses = None

    @staticmethod
    def from_dict(d: dict) -> 'PathItem':
        ci = PathItem()
        if d is None:
            return ci
        ci.consumes = getting = d.get('consumes', None)
        if getting is not None:
            if isinstance(getting, list):
                g_a_cpc_x_20d01ad7 = list()
                for g_a_cpc_x_20d01ad7_value in getting:
                    g_a_cpc_x = str(g_a_cpc_x_20d01ad7_value)
                    g_a_cpc_x_20d01ad7.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_20d01ad7
            ci.consumes = g_a_cpc_x
        ci.produces = getting = d.get('produces', None)
        if getting is not None:
            if isinstance(getting, list):
                g_a_cpc_x_48c2ae92 = list()
                for g_a_cpc_x_48c2ae92_value in getting:
                    g_a_cpc_x = str(g_a_cpc_x_48c2ae92_value)
                    g_a_cpc_x_48c2ae92.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_48c2ae92
            ci.produces = g_a_cpc_x
        ci.tags = getting = d.get('tags', None)
        if getting is not None:
            if isinstance(getting, list):
                g_a_cpc_x_1c5e7af8 = list()
                for g_a_cpc_x_1c5e7af8_value in getting:
                    g_a_cpc_x = str(g_a_cpc_x_1c5e7af8_value)
                    g_a_cpc_x_1c5e7af8.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_1c5e7af8
            ci.tags = g_a_cpc_x
        ci.operationId = getting = d.get('operationId', None)
        if getting is not None:
            ci.operationId = str(getting)
        ci.parameters = getting = d.get('parameters', None)
        if getting is not None:
            if isinstance(getting, list):
                g_a_cpc_x_4481dabf = list()
                for g_a_cpc_x_4481dabf_value in getting:
                    g_a_cpc_x = Parameter.from_dict(g_a_cpc_x_4481dabf_value)
                    g_a_cpc_x_4481dabf.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_4481dabf
            ci.parameters = g_a_cpc_x
        ci.responses = getting = d.get('responses', None)
        if getting is not None:
            if isinstance(getting, dict):
                g_a_cpc_x_78fb5595 = dict()
                for g_a_cpc_x_78fb5595_value in getting.items():
                    g_a_cpc_x = ResponseDesc.from_dict(g_a_cpc_x_78fb5595_value[1])
                    g_a_cpc_x_78fb5595[g_a_cpc_x_78fb5595_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_78fb5595
            ci.responses = g_a_cpc_x
        return ci


class Info(object):
    description: Union[None, 'str']
    title: Union[None, 'str']
    version: Union[None, 'str']

    def __init__(self):
        self.description = None
        self.title = None
        self.version = None

    @staticmethod
    def from_dict(d: dict) -> 'Info':
        ci = Info()
        if d is None:
            return ci
        ci.description = getting = d.get('description', None)
        if getting is not None:
            ci.description = str(getting)
        ci.title = getting = d.get('title', None)
        if getting is not None:
            ci.title = str(getting)
        ci.version = getting = d.get('version', None)
        if getting is not None:
            ci.version = str(getting)
        return ci


class Swagger(object):
    swagger: Union[None, 'str']
    info: Union[None, 'Info']
    basePath: Union[None, 'str']
    paths: Union[None, 'Dict[str, PathItem]']

    def __init__(self):
        self.swagger = None
        self.info = None
        self.basePath = None
        self.paths = None

    @staticmethod
    def from_dict(d: dict) -> 'Swagger':
        ci = Swagger()
        if d is None:
            return ci
        ci.swagger = getting = d.get('swagger', None)
        if getting is not None:
            ci.swagger = str(getting)
        ci.info = getting = d.get('info', None)
        if getting is not None:
            ci.info = Info.from_dict(getting)
        ci.basePath = getting = d.get('basePath', None)
        if getting is not None:
            ci.basePath = str(getting)
        ci.paths = getting = d.get('paths', None)
        if getting is not None:
            if isinstance(getting, dict):
                g_a_cpc_x_7f220d3f = dict()
                for g_a_cpc_x_7f220d3f_value in getting.items():
                    g_a_cpc_x = PathItem.from_dict(g_a_cpc_x_7f220d3f_value[1])
                    g_a_cpc_x_7f220d3f[g_a_cpc_x_7f220d3f_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_7f220d3f
            ci.paths = g_a_cpc_x
        return ci
