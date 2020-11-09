import os
import pathlib

from scripts import config, render


class ApiDoc(object):
    def __init__(self):
        self.template = pathlib.Path(config.script_path, 'api_doc.md').read_text(encoding='utf8')

        def doc_splitter(doc_name):
            return f'<!-- {doc_name}_doc -->'

        self.api_doc = doc_splitter('api')
        self.content_item_doc = doc_splitter('content_item')

        self.api_doc, self.content_item_doc = render.get_doc(
            self.template, [self.api_doc, self.content_item_doc])

    def render(self, methods):
        body_parts = []
        body_parts.append(render.replace_meta_data_all(
            self.content_item_doc, [
                ['content_item', 'Api Specification'],
                ['content_item_link', f'./ApiSpec.md'],
            ]))
        for tag, items in methods.items():
            body_parts.append('\n\n------\n')
            body_parts.append(render.replace_meta_data_all(
                self.content_item_doc, [
                    ['content_item', f'{tag} Api Specification'],
                    ['content_item_link', f'./{tag}.md'],
                ]))

        render.generate_documentation(
            self.api_doc, os.path.join(config.docs_path, 'api', 'README.md'),
            [
                ['table_of_contents', ''.join(body_parts)],
            ])
