import os

from scripts import gen_class

class_defs = \
    """
DataResult:
  handler: str
  method: str
  path: str
  data_records: List[DataRecord]
------
DataRecord:
  comment: str
  request_body: str
  request_header: Dict[str, List[str]]
  response_code: int
  response_body: str
  response_header: Dict[str, List[str]]
------
TestCase:
  abstract: bool
  path: str
  name: str
  meta: Dict[str, object]
  scripts: List[TestCaseScriptStatement]
------
TestCaseScriptStatement:
  func_name: str
  args: List[object]
"""

if __name__ == '__main__':
    current_path = os.path.dirname(__file__)
    gen_class.main(f'{current_path}/unittest_data_classes.py', class_defs)
