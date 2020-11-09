import json

from scripts.api_doc import ApiDoc
from scripts.api_group_doc import ApiGroupDocRenderer
from scripts.api_spec_doc import ApiSpecDoc
from scripts.data import MethodDesc
from scripts.root_doc import RootDocRenderer
from scripts.swagger_data_classes import Swagger

if __name__ == '__main__':
    docs_path = 'docs'
    spec_path = f'{docs_path}/main_spec.json'
    with open(spec_path, 'r', encoding='utf8') as f:
        swagger_j = json.load(f)
    swagger = Swagger.from_dict(swagger_j)

    methods = dict()

    for k, v in swagger.paths.items():
        for m, v_item in v.items():
            if len(v_item.tags) != 1:
                raise IndexError(f"item.tags should only contains only one tag, pos: <{k}@{m}>")
            methods[v_item.tags[0]] = list_ref = methods.get(v_item.tags[0], [])
            list_ref.append(MethodDesc(k, m, v_item))

    # generate api documentation
    ApiGroupDocRenderer().render(methods)

    # generate api documentation
    ApiSpecDoc().render(methods)

    # generate api documentation
    ApiDoc().render(methods)

    # generate root documentation
    RootDocRenderer().render(swagger)
