import os

from scripts import gen_class

class_defs = \
    """
Swagger:
  swagger: str
  info: Info
  basePath: str
  paths: Dict[str, Dict[str, PathItem]]
  definitions: Dict[str, Schema]
------
Info:
  description: str
  title: str
  version: str
------
PathItem:
  consumes: List[str]
  produces: List[str]
  tags: List[str]
  operationId: str
  parameters: List[Parameter]
  responses: Dict[str, ResponseDesc]
------
Parameter:
  type: str
  description: str
  name: str
  in: str
  required: bool
  schema: Schema
  allowEmptyValue: bool
------
ResponseDesc:
  description: str
  schema: Schema
------
Schema:
  id: str
  ref, $ref: str
  schemaUrl: str
  description: str
  type: Union[str, List[str]]
  nullable: bool
  format: str
  title: str
  default: object
  maximum: float
  exclusiveMaximum: bool
  minimum: float
  exclusiveMinimum: bool
  maxLength: int
  minLength: int
  pattern: str
  maxItems: int
  minItems: int
  uniqueItems: bool
  multipleOf: float
  enum: List[object]
  maxProperties: int
  minProperties: int
  required: List[str]
  items: Union[Schema, List[Schema]]
  allOf: List[Schema]
  oneOf: List[Schema]
  anyOf: List[Schema]
  not: Schema
  properties: Dict[str, Schema]
  additionalProperties: Union[Schema, bool]
  patternProperties: Dict[str, Schema]
  dependencies: Dict[str, Union[Schema, List[str]]]
  additionalItems: Union[Schema, bool]
  definitions: Dict[str, Schema]
"""

if __name__ == '__main__':
    current_path = os.path.dirname(__file__)
    gen_class.main(f'{current_path}/swagger_data_classes.py', class_defs)
