import abc
import atexit
import hashlib
import json
import os
import pickle
import tempfile

from go_ast import ASTInfo

caching = dict()


def get_object_digests(cache_path):
    global caching
    if cache_path in caching:
        return caching[cache_path]
    cache_file = os.path.join(cache_path, 'ast-persistence')
    caching[cache_path] = json.load(open(cache_file)) if os.path.exists(cache_file) else [set(), {}]
    caching[cache_path][0] = set(caching[cache_path][0])
    atexit.register(
        lambda: json.dump([list(caching[cache_path][0]), {}], open(cache_file, 'w+')))
    return caching[cache_path]


def get_cached_object(self, f, serialized):
    object_digests = get_object_digests(self.cache_path)
    m = hashlib.md5()
    m.update(serialized.encode())
    d = m.digest().hex()
    if d in object_digests[1]:
        return object_digests[1][d]
    if d in object_digests[0]:
        object_digests[1][d] = o = pickle.load(open(os.path.join(self.cache_path, d), 'rb'))
        return o
    object_digests[0].add(d)
    o = f(self, serialized)
    pickle.dump(o, open(os.path.join(self.cache_path, d), 'wb'))
    return o


def pickle_file_helper(f):
    def wrap(self, file):
        return get_cached_object(self, f, open(file).read())

    return wrap


def pickle_str_helper(f):
    def wrap(self, serialized):
        return get_cached_object(self, f, serialized)

    return wrap


class AstDeserializer(abc.ABC):

    def __init__(self, cache_path=None):
        self.cache_path = cache_path
        if not self.cache_path:
            self.cache_path = os.path.join(tempfile.gettempdir(), 'golang-pack-ast')
            os.makedirs(self.cache_path, exist_ok=True)

    @abc.abstractmethod
    def load_asts_(self, loading):
        return loading

    @pickle_file_helper
    def load_ast(self, loading):
        return self.load_asts_(loading)

    # def load_ast(self, loading):
    #     return self.load_asts_(open(loading).read())

    # @pickle_str_helper
    def load_asts(self, loading):
        return self.load_asts_(loading)


class YAMLAstDeserializer(AstDeserializer):

    def __init__(self, cache_path=None):
        super().__init__(cache_path)
        import yaml
        self.yaml = yaml

    def load_asts_(self, loading):
        return ASTInfo(self.yaml.safe_load(loading))


class JSONAstDeserializer(AstDeserializer):

    def __init__(self, cache_path=None):
        super().__init__(cache_path)
        import json
        self.json = json

    def load_asts_(self, loading):
        return ASTInfo(self.json.loads(loading))
