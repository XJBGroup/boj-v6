import os
import pathlib

from scripts import config, render
from scripts.data import MethodDesc


class ApiSpecDoc(object):
    def __init__(self):
        self.template = pathlib.Path(config.script_path, 'api_spec_doc.md').read_text(encoding='utf8')

        def doc_splitter(doc_name):
            return f'<!-- {doc_name}_doc -->'

        self.api_doc = doc_splitter('api')
        self.api_table_doc = doc_splitter('api_table')

        self.api_doc, self.api_table_doc = render.get_doc(
            self.template, [self.api_doc, self.api_table_doc])

    def render(self, methods):
        body_parts = []
        for tag, items in methods.items():
            tag_methods = []
            for path_item in sorted(items, key=lambda x_item: x_item.key):  # type: MethodDesc

                tag_methods.append(
                    f"+ [{path_item.item.operationId}](./{tag}.md#{path_item.item.operationId}): "
                    f"The uri/restful key of this method is `{path_item.key}`"
                )
                if path_item.item.parameters:
                    params_body_short_desc = []
                    for raw_param in path_item.item.parameters:
                        # print(raw_param.in_, raw_param.allowEmptyValue)

                        # if raw_param.schema:
                        #     print(raw_param.schema.ref)
                        params_body_short_desc.append(
                            f'{raw_param.name} ([{"!" if raw_param.required else ""}{raw_param.type or "any"}]())'
                        )
                    tag_methods.append('')
                    tag_methods.append(
                        '    params: ' + (', '.join(params_body_short_desc))
                    )
            tag_methods.append('')
            body_parts.append(render.replace_meta_data_all(
                self.api_table_doc, [
                    ['tag_methods', '\n    '.join(tag_methods)],
                    ['api_tag_link', f'./{tag}.md'],
                    ['api_tag', tag],
                ]))

        render.generate_documentation(
            self.api_doc, os.path.join(config.docs_path, 'api', 'ApiSpec.md'),
            [
                ['references', ''.join(body_parts)],
            ])
