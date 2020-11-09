import os
import pathlib
from typing import Optional

from scripts import config, render
from scripts.data import MethodDesc


class ApiGroupDocRenderer(object):
    def __init__(self):
        self.template = pathlib.Path(config.script_path, 'api_group_doc.md').read_text(encoding='utf8')

        def doc_splitter(doc_name):
            return f'<!-- {doc_name}_doc -->'

        docs = render.get_doc(
            self.template, list(map(doc_splitter, [
                'group', 'group_body', 'group_params_body',
                'local_object_reference', 'object', 'object_field', ]))
        )
        self.group_doc, self.group_body_doc, self.group_params_body_doc = docs[0:3]
        self.local_object_reference_doc, self.object_doc, self.object_field_doc = docs[3:]

    def render(self, swagger, methods):
        for tag, items in methods.items():
            body_parts = []
            object_references = dict()
            for path_item in sorted(items, key=lambda x_item: x_item.key):  # type: MethodDesc
                body_parts.append(self.render_single_group(swagger, path_item, object_references))
            object_parts = []
            for ref, schema in object_references.items():
                object_fields_parts = []
                if schema.properties is not None:
                    for name, prop in schema.properties.items():
                        object_fields_parts.append(render.replace_meta_data_all(
                            self.object_field_doc, [
                                ['name', name],
                                ['type', prop.type or ''],
                                ['format', prop.format or ''],
                                ['required', ' (required)' if prop.required else ''],
                                ['desc', prop.description or ''],
                            ]))

                the_raw_param_type = schema.type or 'any'
                if ref is None:
                    ref = schema.type
                the_raw_param_type = f'[{the_raw_param_type}](#{ref})'
                object_parts.append(render.replace_meta_data_all(
                    self.object_doc, [
                        ['ref', f'./ObjectModelSpec.md#{ref}'],
                        ['type', the_raw_param_type],
                        ['object_name', ref],
                        ['object_fields', ''.join(object_fields_parts)],
                    ]))
                #     object_fields
            render.generate_documentation(
                self.group_doc, os.path.join(config.docs_path, 'api', f'{tag}.md'),
                [
                    ['body', ''.join(body_parts)],
                    ['object_reference', render.replace_meta_data_all(self.local_object_reference_doc, [
                        ['objects', ''.join(object_parts)]
                    ])],
                    ['api_tag', tag],
                ])

    def render_single_group(self, swagger, path_item, object_references):

        params_body_parts = []
        if path_item.item.parameters:
            for raw_param in path_item.item.parameters:
                # print(raw_param.in_, raw_param.allowEmptyValue)
                ref = None  # type: Optional[str]
                if raw_param.schema:
                    assert raw_param.schema.ref.startswith('#/definitions/')
                    ref = raw_param.schema.ref[len('#/definitions/'):]
                    schema = swagger.definitions[ref]
                    object_references[ref] = schema

                the_raw_param_type = raw_param.type or ref or 'any'
                if ref is None:
                    ref = raw_param.type
                the_raw_param_type = f'[{the_raw_param_type}](#{ref})'

                params_body_parts.append(render.replace_meta_data_all(
                    self.group_params_body_doc, [
                        ['name', raw_param.name],
                        ['type', the_raw_param_type],
                        ['required', ' (required)' if raw_param.required else ''],
                        ['desc', raw_param.description or ''],
                    ]))
        return render.replace_meta_data_all(
            self.group_body_doc, [
                ['params_body', ''.join(params_body_parts)],
                ['operationId', path_item.item.operationId],
                ['key', path_item.key],
            ])
