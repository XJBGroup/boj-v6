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
            g_a_cpc_x = None
            ci.id = str(getting)
        ci.ref = getting = d.get('$ref', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.ref = str(getting)
        ci.schemaUrl = getting = d.get('schemaUrl', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.schemaUrl = str(getting)
        ci.description = getting = d.get('description', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.description = str(getting)
        ci.type = getting = d.get('type', None)
        if getting is not None:
            g_a_cpc_x = None
            g_a_cpc_x_1a6e3267_handling = getting
            if isinstance(g_a_cpc_x_1a6e3267_handling, list):
                g_a_cpc_x_1660d541 = list()
                for g_a_cpc_x_1660d541_value in g_a_cpc_x_1a6e3267_handling:
                    g_a_cpc_x = str(g_a_cpc_x_1660d541_value)
                    g_a_cpc_x_1660d541.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_1660d541
            if g_a_cpc_x is None:
                g_a_cpc_x = str(g_a_cpc_x_1a6e3267_handling)
            ci.type = g_a_cpc_x
        ci.nullable = getting = d.get('nullable', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.nullable = bool(getting)
        ci.format = getting = d.get('format', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.format = str(getting)
        ci.title = getting = d.get('title', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.title = str(getting)
        ci.default = getting = d.get('default', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.default = getting
        ci.maximum = getting = d.get('maximum', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.maximum = float(getting)
        ci.exclusiveMaximum = getting = d.get('exclusiveMaximum', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.exclusiveMaximum = bool(getting)
        ci.minimum = getting = d.get('minimum', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.minimum = float(getting)
        ci.exclusiveMinimum = getting = d.get('exclusiveMinimum', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.exclusiveMinimum = bool(getting)
        ci.maxLength = getting = d.get('maxLength', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.maxLength = int(getting)
        ci.minLength = getting = d.get('minLength', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.minLength = int(getting)
        ci.pattern = getting = d.get('pattern', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.pattern = str(getting)
        ci.maxItems = getting = d.get('maxItems', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.maxItems = int(getting)
        ci.minItems = getting = d.get('minItems', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.minItems = int(getting)
        ci.uniqueItems = getting = d.get('uniqueItems', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.uniqueItems = bool(getting)
        ci.multipleOf = getting = d.get('multipleOf', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.multipleOf = float(getting)
        ci.enum = getting = d.get('enum', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, list):
                g_a_cpc_x_ed7be07 = list()
                for g_a_cpc_x_ed7be07_value in getting:
                    g_a_cpc_x = g_a_cpc_x_ed7be07_value
                    g_a_cpc_x_ed7be07.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_ed7be07
            ci.enum = g_a_cpc_x
        ci.maxProperties = getting = d.get('maxProperties', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.maxProperties = int(getting)
        ci.minProperties = getting = d.get('minProperties', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.minProperties = int(getting)
        ci.required = getting = d.get('required', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, list):
                g_a_cpc_x_128857da = list()
                for g_a_cpc_x_128857da_value in getting:
                    g_a_cpc_x = str(g_a_cpc_x_128857da_value)
                    g_a_cpc_x_128857da.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_128857da
            ci.required = g_a_cpc_x
        ci.items = getting = d.get('items', None)
        if getting is not None:
            g_a_cpc_x = None
            g_a_cpc_x_2433c064_handling = getting
            if isinstance(g_a_cpc_x_2433c064_handling, list):
                g_a_cpc_x_46e82519 = list()
                for g_a_cpc_x_46e82519_value in g_a_cpc_x_2433c064_handling:
                    g_a_cpc_x = Schema.from_dict(g_a_cpc_x_46e82519_value)
                    g_a_cpc_x_46e82519.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_46e82519
            if g_a_cpc_x is None:
                g_a_cpc_x = Schema.from_dict(g_a_cpc_x_2433c064_handling)
            ci.items = g_a_cpc_x
        ci.allOf = getting = d.get('allOf', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, list):
                g_a_cpc_x_54d2f617 = list()
                for g_a_cpc_x_54d2f617_value in getting:
                    g_a_cpc_x = Schema.from_dict(g_a_cpc_x_54d2f617_value)
                    g_a_cpc_x_54d2f617.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_54d2f617
            ci.allOf = g_a_cpc_x
        ci.oneOf = getting = d.get('oneOf', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, list):
                g_a_cpc_x_1b02cd25 = list()
                for g_a_cpc_x_1b02cd25_value in getting:
                    g_a_cpc_x = Schema.from_dict(g_a_cpc_x_1b02cd25_value)
                    g_a_cpc_x_1b02cd25.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_1b02cd25
            ci.oneOf = g_a_cpc_x
        ci.anyOf = getting = d.get('anyOf', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, list):
                g_a_cpc_x_3ba42b6b = list()
                for g_a_cpc_x_3ba42b6b_value in getting:
                    g_a_cpc_x = Schema.from_dict(g_a_cpc_x_3ba42b6b_value)
                    g_a_cpc_x_3ba42b6b.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_3ba42b6b
            ci.anyOf = g_a_cpc_x
        ci.not_ = getting = d.get('not', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.not_ = Schema.from_dict(getting)
        ci.properties = getting = d.get('properties', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, dict):
                g_a_cpc_x_429ac16 = dict()
                for g_a_cpc_x_429ac16_value in getting.items():
                    g_a_cpc_x = Schema.from_dict(g_a_cpc_x_429ac16_value[1])
                    g_a_cpc_x_429ac16[g_a_cpc_x_429ac16_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_429ac16
            ci.properties = g_a_cpc_x
        ci.additionalProperties = getting = d.get('additionalProperties', None)
        if getting is not None:
            g_a_cpc_x = None
            g_a_cpc_x_5330b8d0_handling = getting
            g_a_cpc_x = Schema.from_dict(g_a_cpc_x_5330b8d0_handling)
            if g_a_cpc_x is None:
                g_a_cpc_x = bool(g_a_cpc_x_5330b8d0_handling)
            ci.additionalProperties = g_a_cpc_x
        ci.patternProperties = getting = d.get('patternProperties', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, dict):
                g_a_cpc_x_7d8018f8 = dict()
                for g_a_cpc_x_7d8018f8_value in getting.items():
                    g_a_cpc_x = Schema.from_dict(g_a_cpc_x_7d8018f8_value[1])
                    g_a_cpc_x_7d8018f8[g_a_cpc_x_7d8018f8_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_7d8018f8
            ci.patternProperties = g_a_cpc_x
        ci.dependencies = getting = d.get('dependencies', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, dict):
                g_a_cpc_x_2b45a48d = dict()
                for g_a_cpc_x_2b45a48d_value in getting.items():
                    g_a_cpc_x = None
                    g_a_cpc_x_67320ab8_handling = g_a_cpc_x_2b45a48d_value[1]
                    if isinstance(g_a_cpc_x_67320ab8_handling, list):
                        g_a_cpc_x_7ff5f70b = list()
                        for g_a_cpc_x_7ff5f70b_value in g_a_cpc_x_67320ab8_handling:
                            g_a_cpc_x = str(g_a_cpc_x_7ff5f70b_value)
                            g_a_cpc_x_7ff5f70b.append(g_a_cpc_x)
                        g_a_cpc_x = g_a_cpc_x_7ff5f70b
                    if g_a_cpc_x is None:
                        g_a_cpc_x = Schema.from_dict(g_a_cpc_x_67320ab8_handling)
                    g_a_cpc_x_2b45a48d[g_a_cpc_x_2b45a48d_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_2b45a48d
            ci.dependencies = g_a_cpc_x
        ci.additionalItems = getting = d.get('additionalItems', None)
        if getting is not None:
            g_a_cpc_x = None
            g_a_cpc_x_79ec6e27_handling = getting
            g_a_cpc_x = Schema.from_dict(g_a_cpc_x_79ec6e27_handling)
            if g_a_cpc_x is None:
                g_a_cpc_x = bool(g_a_cpc_x_79ec6e27_handling)
            ci.additionalItems = g_a_cpc_x
        ci.definitions = getting = d.get('definitions', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, dict):
                g_a_cpc_x_3ce16d44 = dict()
                for g_a_cpc_x_3ce16d44_value in getting.items():
                    g_a_cpc_x = Schema.from_dict(g_a_cpc_x_3ce16d44_value[1])
                    g_a_cpc_x_3ce16d44[g_a_cpc_x_3ce16d44_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_3ce16d44
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
            g_a_cpc_x = None
            ci.description = str(getting)
        ci.schema = getting = d.get('schema', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.schema = Schema.from_dict(getting)
        return ci


class Parameter(object):
    type: Union[None, 'str']
    description: Union[None, 'str']
    name: Union[None, 'str']
    in_: Union[None, 'str']
    required: Union[None, 'bool']
    schema: Union[None, 'Schema']
    allowEmptyValue: Union[None, 'bool']

    def __init__(self):
        self.type = None
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
        ci.type = getting = d.get('type', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.type = str(getting)
        ci.description = getting = d.get('description', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.description = str(getting)
        ci.name = getting = d.get('name', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.name = str(getting)
        ci.in_ = getting = d.get('in', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.in_ = str(getting)
        ci.required = getting = d.get('required', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.required = bool(getting)
        ci.schema = getting = d.get('schema', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.schema = Schema.from_dict(getting)
        ci.allowEmptyValue = getting = d.get('allowEmptyValue', None)
        if getting is not None:
            g_a_cpc_x = None
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
            g_a_cpc_x = None
            if isinstance(getting, list):
                g_a_cpc_x_1f0dc9a5 = list()
                for g_a_cpc_x_1f0dc9a5_value in getting:
                    g_a_cpc_x = str(g_a_cpc_x_1f0dc9a5_value)
                    g_a_cpc_x_1f0dc9a5.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_1f0dc9a5
            ci.consumes = g_a_cpc_x
        ci.produces = getting = d.get('produces', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, list):
                g_a_cpc_x_658c7592 = list()
                for g_a_cpc_x_658c7592_value in getting:
                    g_a_cpc_x = str(g_a_cpc_x_658c7592_value)
                    g_a_cpc_x_658c7592.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_658c7592
            ci.produces = g_a_cpc_x
        ci.tags = getting = d.get('tags', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, list):
                g_a_cpc_x_2645af72 = list()
                for g_a_cpc_x_2645af72_value in getting:
                    g_a_cpc_x = str(g_a_cpc_x_2645af72_value)
                    g_a_cpc_x_2645af72.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_2645af72
            ci.tags = g_a_cpc_x
        ci.operationId = getting = d.get('operationId', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.operationId = str(getting)
        ci.parameters = getting = d.get('parameters', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, list):
                g_a_cpc_x_73cbaa67 = list()
                for g_a_cpc_x_73cbaa67_value in getting:
                    g_a_cpc_x = Parameter.from_dict(g_a_cpc_x_73cbaa67_value)
                    g_a_cpc_x_73cbaa67.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_73cbaa67
            ci.parameters = g_a_cpc_x
        ci.responses = getting = d.get('responses', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, dict):
                g_a_cpc_x_5393a26c = dict()
                for g_a_cpc_x_5393a26c_value in getting.items():
                    g_a_cpc_x = ResponseDesc.from_dict(g_a_cpc_x_5393a26c_value[1])
                    g_a_cpc_x_5393a26c[g_a_cpc_x_5393a26c_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_5393a26c
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
            g_a_cpc_x = None
            ci.description = str(getting)
        ci.title = getting = d.get('title', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.title = str(getting)
        ci.version = getting = d.get('version', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.version = str(getting)
        return ci


class Swagger(object):
    swagger: Union[None, 'str']
    info: Union[None, 'Info']
    basePath: Union[None, 'str']
    paths: Union[None, 'Dict[str, Dict[str, PathItem]]']
    definitions: Union[None, 'Dict[str, Schema]']

    def __init__(self):
        self.swagger = None
        self.info = None
        self.basePath = None
        self.paths = None
        self.definitions = None

    @staticmethod
    def from_dict(d: dict) -> 'Swagger':
        ci = Swagger()
        if d is None:
            return ci
        ci.swagger = getting = d.get('swagger', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.swagger = str(getting)
        ci.info = getting = d.get('info', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.info = Info.from_dict(getting)
        ci.basePath = getting = d.get('basePath', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.basePath = str(getting)
        ci.paths = getting = d.get('paths', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, dict):
                g_a_cpc_x_1a249a34 = dict()
                for g_a_cpc_x_1a249a34_value in getting.items():
                    if isinstance(g_a_cpc_x_1a249a34_value[1], dict):
                        g_a_cpc_x_3bf7ea50 = dict()
                        for g_a_cpc_x_3bf7ea50_value in g_a_cpc_x_1a249a34_value[1].items():
                            g_a_cpc_x = PathItem.from_dict(g_a_cpc_x_3bf7ea50_value[1])
                            g_a_cpc_x_3bf7ea50[g_a_cpc_x_3bf7ea50_value[0]] = (g_a_cpc_x)
                        g_a_cpc_x = g_a_cpc_x_3bf7ea50
                    g_a_cpc_x_1a249a34[g_a_cpc_x_1a249a34_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_1a249a34
            ci.paths = g_a_cpc_x
        ci.definitions = getting = d.get('definitions', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, dict):
                g_a_cpc_x_2cedbf = dict()
                for g_a_cpc_x_2cedbf_value in getting.items():
                    g_a_cpc_x = Schema.from_dict(g_a_cpc_x_2cedbf_value[1])
                    g_a_cpc_x_2cedbf[g_a_cpc_x_2cedbf_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_2cedbf
            ci.definitions = g_a_cpc_x
        return ci
