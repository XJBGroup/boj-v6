from go_ast import ASTInfo


class YAMLAstDeserializer(object):

    def __init__(self):
        import yaml
        self.yaml = yaml

    def load_ast(self, file_path):
        return self.load_asts(open(file_path).read())

    def load_asts(self, loading):
        return ASTInfo(self.yaml.safe_load(loading))


class JSONAstDeserializer(object):

    def __init__(self):
        import json
        self.json = json

    def load_ast(self, file_path):
        return ASTInfo(self.json.load(open(file_path)))

    def load_asts(self, loading):
        return ASTInfo(self.json.loads(loading))
