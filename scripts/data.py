from scripts.swagger_data_classes import PathItem


class MethodDesc:
    def __init__(self, path, method, item):
        self.path = path
        self.method = method
        self.item = item  # type: PathItem
        self.key = self.path + '@' + self.method.upper()
