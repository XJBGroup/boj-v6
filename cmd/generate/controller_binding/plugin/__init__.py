import os

from config import Plugin, GolangPackConfig


class CopySourcePlugin(Plugin):

    def __init__(self):
        pass

    def before_loader(self, config: GolangPackConfig):
        if not config.local_package:
            raise KeyError("CopyGolangSourcePlugin not support config without local_package set")

        src = os.path.join(config.local_package, os.path.relpath(config.src, config.package))
        output = os.path.join(config.local_package, os.path.relpath(config.output, config.package))

        for r, _, files in os.walk(src):
            for f in files:
                if f.endswith('.go') and f.count('.') > 1:
                    continue

                x_f = os.path.join(r, f)
                y_f = os.path.join(output, f)
                if os.path.exists(y_f):
                    os.unlink(y_f)
                os.link(x_f, y_f)
            break

    def after_loader(self, config: GolangPackConfig):
        pass
