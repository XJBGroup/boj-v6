import json
import os
import pathlib
import re

from scripts.swagger_data_classes import Swagger, PathItem


def replace_meta_data(the_tmpl, key, value):
    return the_tmpl.replace(f'{{{{{key}}}}}', value)


def replace_meta_data_all(the_tmpl, metas):
    for meta in metas:
        the_tmpl = replace_meta_data(the_tmpl, meta[0], meta[1])
    return the_tmpl


def replace_insert_fields(matched, the_tmpl):
    bx = []
    last_xx = 0
    for x in re.finditer(r'([ ]*)!!insert_field:([^!]*)!!', the_tmpl):
        bx.append(the_tmpl[last_xx:x.span()[0]])
        indent, field_id = x[1], x[2]
        if field_id not in matched:
            bx.append(f'{indent}<!--beg l {field_id} -->\n{indent}\n{indent}<!--end l-->')
        else:
            bx.append(f'{indent}<!--beg l {field_id} -->{matched[field_id]}<!--end l-->')
        last_xx = x.span()[1]
    bx.append(the_tmpl[last_xx:])
    return ''.join(bx)


def prepare_path(raw_path):
    if isinstance(raw_path, str):
        raw_path = pathlib.Path(raw_path)
        raw_path.parent.mkdir(parents=True, exist_ok=True)
    return raw_path


def generate_documentation(tmpl, path, metas, in_memory_doc=None, wait_after_insert=False):
    path = prepare_path(path)
    matched = dict()
    if in_memory_doc is not None:
        path_x = in_memory_doc
    elif path.exists():
        path_x = path.read_text(encoding='utf8')
    else:
        path_x = ''

    for x in re.finditer(r'<!--beg l (\S*) -->([\s\S]*)<!--end l-->', path_x):
        matched[x[2]] = x[1]

    # step2: replace insert field
    tmpl = replace_insert_fields(
        matched,
        # step1: replace meta data
        replace_meta_data_all(tmpl, metas))

    if not wait_after_insert:
        # step3: write text
        path.write_text(tmpl, encoding='utf8')

    return tmpl


if __name__ == '__main__':
    docs_path = 'docs'
    spec_path = f'{docs_path}/main_spec.json'
    with open(spec_path, 'r', encoding='utf8') as f:
        swagger_j = json.load(f)
    swagger = Swagger.from_dict(swagger_j)


    class MethodDesc:
        def __init__(self, path, method, item):
            self.path = path
            self.method = method
            self.item = item  # type: PathItem
            self.key = self.path + '@' + self.method.upper()


    methods = dict()

    for k, v in swagger.paths.items():
        for m, v_item in v.items():
            if len(v_item.tags) != 1:
                raise IndexError(f"item.tags should only contains only one tag, pos: <{k}@{m}>")
            methods[v_item.tags[0]] = list_ref = methods.get(v_item.tags[0], [])
            list_ref.append(MethodDesc(k, m, v_item))

    ApiGroupDocTemplate = """
# Api Group Documentation ({{api_tag}})

{{body}}
"""
    ApiGroupDocBodyTemplate = """
## {{operationId}}

The uri/restful key of this method is `{{key}}`

!!insert_field:desc_{{operationId}}!!

{{params_body}}

"""
    ApiGroupDocParamTemplate = """
+ `{{name}}`: `{{type}}`{{required}}: {{desc}}
    !!insert_field:desc_{{operationId}}_{{name}}!!
"""

    for tag, items in methods.items():
        body_parts = []
        for path_item in sorted(items, key=lambda x_item: x_item.key):  # type: MethodDesc

            params = []
            params_body_parts = []
            if path_item.item.parameters:
                for raw_param in path_item.item.parameters:
                    # print(raw_param.in_, raw_param.allowEmptyValue)

                    if raw_param.schema:
                        print(raw_param.schema.ref)

                    params_body_parts.append(replace_meta_data_all(
                        ApiGroupDocParamTemplate, [
                            ['name', raw_param.name],
                            ['type', raw_param.type or 'any'],
                            ['required', ' (required)' if raw_param.required else ''],
                            ['desc', raw_param.description or ''],
                        ]))
            body_parts.append(replace_meta_data_all(
                ApiGroupDocBodyTemplate, [
                    ['params_body', ''.join(params_body_parts)],
                    ['operationId', path_item.item.operationId],
                    ['key', path_item.key],
                ]))

        generate_documentation(
            ApiGroupDocTemplate, os.path.join(docs_path, 'api', f'{tag}.md'),
            [
                ['body', ''.join(body_parts)],
                ['api_tag', tag],
            ])

    # generate root documentation

    ReadmeTemplate = """
# {{doc_title}} Root Documentation
this is the root documentation of {{doc_title}}.

{{doc_desc}}

<!-- you can complete description in the following field -->

!!insert_field:desc!!

## table of contents
"""

    generate_documentation(
        ReadmeTemplate, os.path.join(docs_path, 'README.md'),
        [
            ['doc_title', swagger.info.title],
            ['doc_desc', swagger.info.description],
        ])
