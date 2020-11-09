import os
import pathlib

from scripts import config, render
from scripts.data import MethodDesc


class ApiGroupDocRenderer(object):
    def __init__(self):
        self.template = pathlib.Path(config.script_path, 'api_group_doc.md').read_text(encoding='utf8')

        def doc_splitter(doc_name):
            return f'<!-- {doc_name}_doc -->'

        self.group_doc = doc_splitter('group')
        self.group_body_doc = doc_splitter('group_body')
        self.group_params_body_doc = doc_splitter('group_params_body')

        self.group_doc, self.group_body_doc, self.group_params_body_doc = render.get_doc(
            self.template,
            [self.group_doc, self.group_body_doc, self.group_params_body_doc]
        )

    def render(self, methods):
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

                        params_body_parts.append(render.replace_meta_data_all(
                            self.group_params_body_doc, [
                                ['name', raw_param.name],
                                ['type', raw_param.type or 'any'],
                                ['required', ' (required)' if raw_param.required else ''],
                                ['desc', raw_param.description or ''],
                            ]))
                body_parts.append(render.replace_meta_data_all(
                    self.group_body_doc, [
                        ['params_body', ''.join(params_body_parts)],
                        ['operationId', path_item.item.operationId],
                        ['key', path_item.key],
                    ]))

            render.generate_documentation(
                self.group_doc, os.path.join(config.docs_path, 'api', f'{tag}.md'),
                [
                    ['body', ''.join(body_parts)],
                    ['api_tag', tag],
                ])
