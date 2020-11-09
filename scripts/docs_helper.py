import json
import os
import pathlib
import re

from scripts.swagger_data_classes import Swagger

ReadmeTemplate = """
# {{docs}} Root Documentation
this is the root documentation of {{docs}}.

{{description}}

<!-- you can complete description in the following field -->

!!insert_field:desc!!

## table of contents
"""


def replace_insert_fields(path, the_tmpl):
    matched = dict()
    if path.exists():
        path_x = path.read_text(encoding='utf8')
        for x in re.finditer(r'-------\n<!--(\S*) -->([\s\S]*)-------', path_x):
            matched[x[1]] = x[2]

    bx = []
    last_xx = 0
    for x in re.finditer(r'!!insert_field:([^!]*)!!', the_tmpl):
        bx.append(the_tmpl[last_xx:x.span()[0]])
        if x[1] not in matched:
            bx.append(f'-------\n<!--{x[1]} -->\n\n-------')
        else:
            bx.append(f'-------\n<!--{x[1]} -->{matched[x[1]]}-------')
        last_xx = x.span()[1]
    bx.append(the_tmpl[last_xx:])
    return ''.join(bx)


if __name__ == '__main__':
    docs_path = 'docs'
    spec_path = f'{docs_path}/main_spec.json'
    with open(spec_path, 'r', encoding='utf8') as f:
        swagger_j = json.load(f)
    swagger = Swagger.from_dict(swagger_j)

    tmpl = ReadmeTemplate
    tmpl = tmpl.replace('{{docs}}', swagger.info.title)
    tmpl = tmpl.replace('{{description}}', swagger.info.description)

    root_doc = pathlib.Path(os.path.join(docs_path, 'README.md'))
    tmpl = replace_insert_fields(root_doc, tmpl)
    root_doc.write_text(tmpl, encoding='utf8')
