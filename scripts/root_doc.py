import os
import pathlib

from scripts import config, render


class RootDocRenderer(object):
    def __init__(self):
        self.template = pathlib.Path(config.script_path, 'root_doc.md').read_text(encoding='utf8')

    def render(self, swagger):
        render.generate_documentation(
            self.template, os.path.join(config.docs_path, 'README.md'),
            [
                ['doc_title', swagger.info.title],
                ['doc_desc', swagger.info.description],
            ])
