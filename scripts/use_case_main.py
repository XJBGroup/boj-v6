import json
import os

from scripts.data import MethodDesc
from scripts.swagger_data_classes import Swagger
from scripts.unittest_data_classes import DataResult

if __name__ == '__main__':
    docs_path = 'docs'
    spec_path = f'{docs_path}/main_spec.json'
    with open(spec_path, 'r', encoding='utf8') as f:
        swagger_j = json.load(f)
    swagger = Swagger.from_dict(swagger_j)

    methods = dict()
    fast_ref = dict()

    bad_ref = object()

    for k, v in swagger.paths.items():
        for m, v_item in v.items():
            if len(v_item.tags) != 1:
                raise IndexError(f"item.tags should only contains only one tag, pos: <{k}@{m}>")
            methods[v_item.tags[0]] = dict_ref = methods.get(v_item.tags[0], dict())
            md = MethodDesc(k, m, v_item)
            dict_ref[md.item.operationId] = md
            if md.item.operationId in fast_ref:
                fast_ref[md.item.operationId] = bad_ref
            else:
                fast_ref[md.item.operationId] = md

    datasets_fn = os.listdir(f'{docs_path}/test_cases')
    for dataset_fn in datasets_fn:
        if dataset_fn.endswith('.json'):
            with open(os.path.join('docs/test_cases', dataset_fn), 'r', encoding='utf8') as f:
                dataset_j = json.load(f)
            dataset = list(map(DataResult.from_dict, dataset_j))
            for testcase in dataset:
                golang_func_handler = testcase.handler
                if golang_func_handler.endswith('-fm'):
                    golang_func_handler = golang_func_handler[:-3]
                golang_func_handler = os.path.basename(golang_func_handler)
                select_path = golang_func_handler.split('.')
                method_desc = None
                if len(select_path) > 2:
                    select_path = select_path[-2:]
                if len(select_path) <= 0:
                    raise IndexError("select path length should not be zero")
                if select_path[-1] in fast_ref:
                    md = fast_ref[select_path[-1]]
                    if md is not bad_ref:
                        method_desc = md
                if len(select_path) == 2:
                    if select_path[0] in methods:
                        dict_ref = methods[select_path[0]]
                        if select_path[1] in dict_ref:
                            method_desc = dict_ref[select_path[1]]
                print(testcase.handler, '=>', method_desc)
