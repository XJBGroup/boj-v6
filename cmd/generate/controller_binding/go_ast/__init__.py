import abc
from typing import List

from utils import simplify_path
from utils.cache_io import cached_io


class HasFileMapping(abc.ABC):

    def use_file_mapping(self, mp):
        pos_prop = getattr(self, 'pos', None)
        if pos_prop:
            pos_prop.use_file_mapping(mp)
        mp_attrs = getattr(self.__class__, 'file_mapping_attrs', None)
        if mp_attrs:
            for prop_name in mp_attrs:
                p = getattr(self, prop_name)
                if isinstance(p, list):
                    for sp in p:
                        sp.use_file_mapping(mp)
                elif p:
                    p.use_file_mapping(mp)


class Object(object):
    def __init__(self, doc_tree=None):
        if doc_tree:
            self.name = doc_tree['n']
            self.type = doc_tree['t']
        else:
            self.name = self.type = ''

    def __repr__(self):
        return f'{self.name}(T:{self.type})'

    def __str__(self):
        return self.name

    @staticmethod
    def create(name, t):
        x = Object()
        x.name = name
        x.type = t
        return x


class FilePos(object):
    def __init__(self, doc_tree):
        self.file = doc_tree['f']
        self.line = doc_tree['l']
        self.column = doc_tree['c']
        self.offset = doc_tree['o']
        self.length = doc_tree['s']

    def __repr__(self):
        return ':'.join(map(str, [self.file, self.line, self.column]))

    def use_file_mapping(self, mp):
        self.file = mp[self.file]


class WithFilePos(object):
    def __init__(self, doc_tree):
        if doc_tree:
            self.pos = FilePos(doc_tree['p'])
        else:
            self.pos = None

    @property
    def body_content(self):
        if self.pos is None:
            raise KeyError(f"in memory stmt({self.__class__.__name__}) has no file position")
        f = cached_io.open_read(self.pos.file)
        return f[self.pos.offset:self.pos.offset + self.pos.length]

    def __str__(self):
        return self.body_content

    def __repr__(self):
        content = self.body_content
        if len(content) > 30:
            content = content[:30] + '...'
        if self.pos:
            pos = ':'.join(map(str, [self.pos.line, self.pos.column, self.pos.length]))
        else:
            pos = ':memory:'
        return f"<{self.__class__.__name__} at {pos}, body='{content}'> "


class FilesMapping(object):
    def __init__(self, doc_tree):
        self.mapping = doc_tree
        self.mapping_key_type = int
        for k in self.mapping:
            if isinstance(k, str):
                self.mapping_key_type = str
            break

    def __getitem__(self, k):
        return simplify_path(self.mapping[self.mapping_key_type(k)])


class ImportStmts(HasFileMapping):
    file_mapping_attrs = ('items',)

    class Item(HasFileMapping):
        def __init__(self, doc_tree):
            self.pos = FilePos(doc_tree['p'])
            self.alias = doc_tree['alias']
            self.path = doc_tree['path']

        def __repr__(self):
            if self.alias is not None and len(self.alias):
                return f'<{self.alias} from {self.path} at {repr(self.pos)}>'
            return f'<{self.path} at {repr(self.pos)}>'

    def __init__(self, doc_tree):
        self.items = list(map(ImportStmts.Item, doc_tree))

    def __repr__(self):
        items = self.items
        if len(items) > 20:
            return '\n'.join(map(repr, items[:5])) + '\n...'
        return '\n'.join(map(repr, items))


class Stmt(WithFilePos):
    polymorphic = dict()

    @staticmethod
    def create_stmt(doc_tree):
        if doc_tree['t'] not in Stmt.polymorphic:
            raise KeyError(f"stmt type {doc_tree['t']} not registered")

        return Stmt.polymorphic[doc_tree['t']](doc_tree)


class BinaryExp(HasFileMapping, Stmt):
    file_mapping_attrs = ('lhs', 'rhs')

    def __init__(self, doc_tree):
        super().__init__(doc_tree)
        self.o = doc_tree['o']
        self.lhs = Stmt.create_stmt(doc_tree['l'])  # type: Stmt
        self.rhs = Stmt.create_stmt(doc_tree['r'])  # type: Stmt

    def __str__(self):
        return str(self.lhs) + self.o + str(self.rhs)


class UnaryExp(HasFileMapping, Stmt):
    file_mapping_attrs = ('lhs',)

    def __init__(self, doc_tree):
        super().__init__(doc_tree)
        self.o = doc_tree['o']
        self.lhs = Stmt.create_stmt(doc_tree['l'])  # type: Stmt

    def __str__(self):
        if len(self.o) == 2:
            return str(self.lhs) + self.o
        return self.o + str(self.lhs)


class CallExp(HasFileMapping, Stmt):
    file_mapping_attrs = ('callee', 'ins')

    def __init__(self, doc_tree):
        super().__init__(doc_tree)
        self.callee = Stmt.create_stmt(doc_tree['c'])  # type: Stmt
        self.variadic = doc_tree['v']
        self.ins = list(map(Stmt.create_stmt, doc_tree.get('i') or []))  # type: List[Stmt]


class OpaqueExp(HasFileMapping, Stmt):
    def __init__(self, doc_tree):
        super().__init__(doc_tree)
        self.opaque = "" if doc_tree is None else doc_tree["o"]

    @staticmethod
    def create(opaque):
        o = OpaqueExp(None)
        o.opaque = opaque
        return o

    def __str__(self):
        return self.opaque


class AssignExp(HasFileMapping, Stmt):
    file_mapping_attrs = ('lhs', 'rhs')

    def __init__(self, doc_tree):
        super().__init__(doc_tree)
        self.lhs = list(map(Stmt.create_stmt, doc_tree['l']))  # type: List[Stmt]
        self.rhs = list(map(Stmt.create_stmt, doc_tree['r']))  # type: List[Stmt]
        self.tok = doc_tree['o']

    def __str__(self):
        return (','.join(map(str, self.lhs))) + f" {self.tok} " + (','.join(map(str, self.rhs)))


class BlockExp(HasFileMapping, Stmt):
    file_mapping_attrs = ('items',)

    def __init__(self, doc_tree):
        super().__init__(doc_tree)
        self.items = list(map(Stmt.create_stmt, doc_tree['b']))  # type: List[Stmt]

    def __str__(self):
        return '\n'.join(map(str, self.items))


class SelectExp(BlockExp):
    def __init__(self, doc_tree):
        super().__init__(doc_tree)


class IfExp(HasFileMapping, Stmt):
    def __init__(self, doc_tree):
        super().__init__(doc_tree)


class IdentExp(HasFileMapping, Stmt):
    def __init__(self, doc_tree):
        super().__init__(doc_tree)
        self.ident = Object(doc_tree['o'])


class VarDeclExp(HasFileMapping, Stmt):
    file_mapping_attrs = ('decls',)

    def __init__(self, doc_tree):
        super().__init__(doc_tree)
        self.decls = list(map(Stmt.create_stmt, doc_tree['s']))  # type: List[VarSpecExp]


class VarSpecExp(HasFileMapping, Stmt):
    file_mapping_attrs = ('values', 'type_spec')

    def __init__(self, doc_tree):
        super().__init__(doc_tree)
        self.names = doc_tree['l']
        if 'r' in doc_tree:
            self.values = list(map(Stmt.create_stmt, doc_tree.get('r') or []))  # type: List[Stmt]
        else:
            self.values = []
        if doc_tree.get('ts'):
            self.type_spec = Stmt.create_stmt(doc_tree['ts'])
        else:
            self.type_spec = None


class TypeSpecExp(HasFileMapping, Stmt):
    file_mapping_attrs = ('type_spec',)

    def __init__(self, doc_tree):
        super().__init__(doc_tree)
        self.names = doc_tree['n']
        if doc_tree.get('ts'):
            self.type_spec = Stmt.create_stmt(doc_tree['ts'])
        else:
            self.type_spec = None


class SelectorExp(HasFileMapping, Stmt):
    file_mapping_attrs = ('x',)

    def __init__(self, doc_tree):
        super().__init__(doc_tree)
        self.x = Stmt.create_stmt(doc_tree['x'])  # type: Stmt
        self.name = doc_tree['n']


Stmt.polymorphic = {
    'b': BinaryExp,
    'u': UnaryExp,
    'c': CallExp,
    'o': OpaqueExp,
    'a': AssignExp,
    'k': BlockExp,
    's': SelectExp,
    'i': IfExp,
    'var': VarDeclExp,
    'id': IdentExp,
    'v': VarSpecExp,
    't': TypeSpecExp,
    'l': SelectorExp,
}


class FuncBody(BlockExp):
    file_mapping_attrs = ('items',)

    def __init__(self, doc_tree):
        super().__init__(doc_tree)


class FuncDesc(HasFileMapping, WithFilePos):
    file_mapping_attrs = ('body',)

    def __init__(self, doc_tree):
        super().__init__(doc_tree)
        self.recv = Object(doc_tree['r'])
        self.name = doc_tree['n']
        self.ins = list(map(Object, doc_tree.get('in') or []))
        self.outs = list(map(Object, doc_tree.get('out') or []))
        self.body = FuncBody(doc_tree['body'])

    def __repr__(self):
        return f"<{self.recv.type}.{self.name} at {repr(self.pos)}, body={repr(self.body)}>"

    def __str__(self):
        ins = ", ".join(
            map(lambda obj: f"{obj.name} {obj.type}", self.ins))
        outs = map(lambda obj: f"{obj.name} {obj.type}", self.outs)
        if len(self.outs) == 0:
            outs = ""
        elif len(self.outs) == 1 and len(self.outs[0].name) == 0:
            outs = next(outs)
        else:
            outs = f"({', '.join(outs)})"
        return f'func ({self.recv.name} {self.recv.type}) {self.name}({ins}) {outs} {{\n{str(self.body)}\n}}'


class FuncDescs(HasFileMapping, object):
    file_mapping_attrs = ('items',)

    def __init__(self, doc_tree):
        self.items = list(map(FuncDesc, doc_tree))

    def __repr__(self):
        items = self.items
        if len(items) > 20:
            return '\n'.join(map(repr, items[:5])) + '\n...'
        return '\n'.join(map(repr, items))


def use_file_mapping_not_report_error(target, mp):
    target.use_file_mapping(mp)


class ASTInfo(object):
    def __init__(self, doc_tree):
        self.file_mapping = FilesMapping(doc_tree['file_mapping'])
        self.imports = ImportStmts(doc_tree['imports'])
        self.functions = FuncDescs(doc_tree['functions'])
        # noinspection PyTypeChecker
        use_file_mapping_not_report_error(self.imports, self.file_mapping)
        # noinspection PyTypeChecker
        use_file_mapping_not_report_error(self.functions, self.file_mapping)

    def __repr__(self):
        return f"""imports:\n{repr(self.imports)}\nfunctions:\n{repr(self.functions)}"""
