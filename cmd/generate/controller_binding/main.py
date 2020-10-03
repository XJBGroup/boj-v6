import hashlib
import logging
import os
import re
import subprocess
from abc import ABC, abstractmethod
from collections import deque
from dataclasses import dataclass
from typing import Union, List, Optional, Tuple, Dict, Iterable

current_path = os.getcwd()


def simplify_path(path):
    return os.path.relpath(path, current_path).replace('\\', '/')


class GolangToolInvokeError(Exception):
    code: int
    msg: str

    def __init__(self, *args, code, error_dump):
        self.code = code
        self.msg = error_dump


@dataclass
class GolangToolsConfiguration(object):
    dump_tool_package: str = 'github.com/Myriad-Dreamin/boj-v6/cmd/generate/controller_binding/ast-dump'
    dump_cache_path: str = '.cache/ast_dump'
    force_update: bool = False
    verbose: bool = False


class GolangToolsImpl(object):
    class DefaultRunner(object):
        def __init__(self, verbose=False):
            self.verbose = verbose

        def run_command(self, cmd: Union[List[str], str], timeout: Optional[float] = None) -> Tuple[int, str, str]:
            if not isinstance(cmd, str):
                cmd = ' '.join(cmd)
            if self.verbose:
                print(cmd)
            process = subprocess.Popen(cmd)
            process.wait(timeout=timeout)
            code = process.returncode
            stdout, stderr = process.communicate()

            if isinstance(stdout, bytes):
                stdout = stdout.decode()

            if isinstance(stderr, bytes):
                stderr = stderr.decode()

            return code, stdout, stderr

    runner: DefaultRunner
    config: GolangToolsConfiguration

    def __init__(self, runner=None, config=None, **kwargs):
        self.config = config or GolangToolsConfiguration()
        self.runner = runner or GolangToolsImpl.DefaultRunner(self.config.verbose)

    def run_command(self, cmd: Union[List[str], str], timeout: Optional[float] = None) -> str:
        c, o, e = self.runner.run_command(cmd, timeout)
        if c != 0:
            raise GolangToolInvokeError(code=c, error_dump=e)
        if e is not None:
            logging.warning(e)
        return o

    def go_run(self, cmd: Union[List[str], str], timeout: Optional[float] = None) -> str:
        return self.run_command(['go run'] + cmd, timeout)

    def go_fmt(self, cmd: Union[List[str], str], timeout: Optional[float] = None) -> str:
        return self.run_command(['go fmt'] + cmd, timeout)

    def dump_ast_raw(self, dumping_package):
        self.go_run([self.config.dump_tool_package,
                     dumping_package, self.config.dump_cache_path])
        return


def use_file_mapping(maybe_wrapping_cls: Union[type, str], *args):
    if isinstance(maybe_wrapping_cls, str):
        args = (maybe_wrapping_cls,) + args

    def use_file_mapping_2step(wrapping_cls):
        class WrappedClass(wrapping_cls, object):

            __file_mapping_attrs = args

            def use_file_mapping(self, mp):
                pos_prop = getattr(self, 'pos', None)
                if pos_prop:
                    pos_prop.use_file_mapping(mp)
                if self.__file_mapping_attrs:
                    for prop_name in self.__file_mapping_attrs:
                        p = getattr(self, prop_name)
                        if isinstance(p, list):
                            for sp in p:
                                sp.use_file_mapping(mp)
                        elif p:
                            p.use_file_mapping(mp)

        WrappedClass.__name__ = wrapping_cls.__name__
        return WrappedClass

    if isinstance(maybe_wrapping_cls, str):
        return use_file_mapping_2step
    return use_file_mapping_2step(maybe_wrapping_cls)


def use_file_mapping_not_report_error(target, mp):
    target.use_file_mapping(mp)


class CachedIO:
    cached: Dict[str, str] = dict()

    @staticmethod
    def open_read(file_name):
        file_name = os.path.realpath(file_name)

        if not isinstance(file_name, str):
            raise TypeError(f"want file name is str type, got {type(file_name)}")

        if file_name in CachedIO.cached:
            return CachedIO.cached[file_name]

        CachedIO.cached[file_name] = content = open(file_name, 'rb').read().decode('utf8')
        return content

    @staticmethod
    def open_write(file_name, content, *args, **kwargs):
        file_name = os.path.realpath(file_name)

        if file_name in CachedIO.cached:
            CachedIO.cached[file_name] = content

        open(file_name, *args, **kwargs).write(content)

    @staticmethod
    def open_digest(file_name):
        file_name = os.path.realpath(file_name)

        m = hashlib.md5()
        m.update(CachedIO.open_read(file_name).encode())
        return m.digest().hex()


# only cached in same thread
cached_io = CachedIO()


class ASTInfo(object):
    class Object(object):
        def __init__(self, doc_tree=None):
            if doc_tree:
                self.name = doc_tree['n']
                self.type = doc_tree['t']
            else:
                self.name = self.type = ''

        def __repr__(self):
            return f'{self.name}(T:{self.type})'

        @staticmethod
        def create(name, t):
            x = ASTInfo.Object()
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
                self.pos = ASTInfo.FilePos(doc_tree['p'])
            else:
                self.pos = None

        @property
        def body_content(self):
            if self.pos is None:
                raise KeyError("in memory stmt has no file position")
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

        def __getitem__(self, k):
            return simplify_path(self.mapping[k])

    @use_file_mapping('items')
    class ImportStmts(object):

        @use_file_mapping
        class Item(object):
            def __init__(self, doc_tree):
                self.pos = ASTInfo.FilePos(doc_tree['p'])
                self.alias = doc_tree['alias']
                self.path = doc_tree['path']

            def __repr__(self):
                if self.alias is not None and len(self.alias):
                    return f'<{self.alias} from {self.path} at {repr(self.pos)}>'
                return f'<{self.path} at {repr(self.pos)}>'

        def __init__(self, doc_tree):
            self.items = list(map(ASTInfo.ImportStmts.Item, doc_tree))

        def __repr__(self):
            items = self.items
            if len(items) > 20:
                return '\n'.join(map(repr, items[:5])) + '\n...'
            return '\n'.join(map(repr, items))

    class Stmt(WithFilePos):
        polymorphic = dict()

        @staticmethod
        def create_stmt(doc_tree):
            if doc_tree['t'] not in ASTInfo.Stmt.polymorphic:
                raise KeyError(f"stmt type {doc_tree['t']} not registered")

            return ASTInfo.Stmt.polymorphic[doc_tree['t']](doc_tree)

    @use_file_mapping
    class BinaryExp(Stmt):
        def __init__(self, doc_tree):
            super().__init__(doc_tree)

    @use_file_mapping
    class UnaryExp(Stmt):
        def __init__(self, doc_tree):
            super().__init__(doc_tree)

    @use_file_mapping('callee', 'ins')
    class CallExp(Stmt):
        def __init__(self, doc_tree):
            super().__init__(doc_tree)
            self.callee = ASTInfo.Stmt.create_stmt(doc_tree['c'])  # type: ASTInfo.Stmt
            self.variadic = doc_tree['v']
            self.ins = list(map(ASTInfo.Stmt.create_stmt, doc_tree['i']))  # type: List[ASTInfo.Stmt]

    @use_file_mapping
    class OpaqueExp(Stmt):
        def __init__(self, doc_tree):
            super().__init__(doc_tree)
            self.opaque = "" if doc_tree is None else doc_tree["o"]

        @staticmethod
        def create(opaque):
            o = ASTInfo.OpaqueExp(None)
            o.opaque = opaque
            return o

        def __str__(self):
            return self.opaque

    @use_file_mapping('lhs', 'rhs')
    class AssignExp(Stmt):
        def __init__(self, doc_tree):
            super().__init__(doc_tree)
            self.lhs = list(map(ASTInfo.Stmt.create_stmt, doc_tree['l']))  # type: List[ASTInfo.Stmt]
            self.rhs = list(map(ASTInfo.Stmt.create_stmt, doc_tree['r']))  # type: List[ASTInfo.Stmt]

        def __str__(self):
            return (','.join(map(str, self.lhs))) + " := " + (','.join(map(str, self.rhs)))

    @use_file_mapping('items')
    class BlockExp(Stmt):

        def __init__(self, doc_tree):
            super().__init__(doc_tree)
            self.items = list(map(ASTInfo.Stmt.create_stmt, doc_tree['b']))  # type: List[ASTInfo.Stmt]

        def __str__(self):
            return '\n'.join(map(str, self.items))

    @use_file_mapping
    class SelectExp(BlockExp):
        def __init__(self, doc_tree):
            super().__init__(doc_tree)

    @use_file_mapping
    class IfExp(Stmt):
        def __init__(self, doc_tree):
            super().__init__(doc_tree)

    @use_file_mapping
    class IdentExp(Stmt):
        def __init__(self, doc_tree):
            super().__init__(doc_tree)
            self.ident = ASTInfo.Object(doc_tree['o'])

    @use_file_mapping('decls')
    class VarDeclExp(Stmt):
        def __init__(self, doc_tree):
            super().__init__(doc_tree)
            self.decls = list(map(ASTInfo.Stmt.create_stmt, doc_tree['s']))  # type: List[ASTInfo.VarSpecExp]

    @use_file_mapping('values', 'type_spec')
    class VarSpecExp(Stmt):
        def __init__(self, doc_tree):
            super().__init__(doc_tree)
            self.names = doc_tree['l']
            if 'r' in doc_tree:
                self.values = list(map(ASTInfo.Stmt.create_stmt, doc_tree['r']))  # type: List[ASTInfo.Stmt]
            else:
                self.values = []
            if doc_tree.get('ts'):
                self.type_spec = ASTInfo.Stmt.create_stmt(doc_tree['ts'])
            else:
                self.type_spec = None

    @use_file_mapping('type_spec')
    class TypeSpecExp(Stmt):
        def __init__(self, doc_tree):
            super().__init__(doc_tree)
            self.names = doc_tree['n']
            if doc_tree.get('ts'):
                self.type_spec = ASTInfo.Stmt.create_stmt(doc_tree['ts'])
            else:
                self.type_spec = None

    @use_file_mapping('x')
    class SelectorExp(Stmt):
        def __init__(self, doc_tree):
            super().__init__(doc_tree)
            self.x = ASTInfo.Stmt.create_stmt(doc_tree['x'])  # type: ASTInfo.Stmt
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

    @use_file_mapping('items')
    class FuncBody(BlockExp):
        def __init__(self, doc_tree):
            super().__init__(doc_tree)

    @use_file_mapping('body')
    class FuncDesc(WithFilePos):
        def __init__(self, doc_tree):
            super().__init__(doc_tree)
            self.recv = ASTInfo.Object(doc_tree['r'])
            self.name = doc_tree['n']
            self.ins = list(map(ASTInfo.Object, doc_tree['in']))
            self.outs = list(map(ASTInfo.Object, doc_tree['out']))
            self.body = ASTInfo.FuncBody(doc_tree['body'])

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

    @use_file_mapping('items')
    class FuncDescs(object):

        def __init__(self, doc_tree):
            self.items = list(map(ASTInfo.FuncDesc, doc_tree))

        def __repr__(self):
            items = self.items
            if len(items) > 20:
                return '\n'.join(map(repr, items[:5])) + '\n...'
            return '\n'.join(map(repr, items))

    def __init__(self, doc_tree):
        self.file_mapping = ASTInfo.FilesMapping(doc_tree['file_mapping'])
        self.imports = ASTInfo.ImportStmts(doc_tree['imports'])
        self.functions = ASTInfo.FuncDescs(doc_tree['functions'])
        # noinspection PyTypeChecker
        use_file_mapping_not_report_error(self.imports, self.file_mapping)
        # noinspection PyTypeChecker
        use_file_mapping_not_report_error(self.functions, self.file_mapping)

    def __repr__(self):
        return f"""imports:
{repr(self.imports)}
functions:
{repr(self.functions)}"""


class YAMLAstDeserializer(object):

    def __init__(self):
        import yaml
        self.yaml = yaml

    def load_ast(self, file_path):
        return self.load_asts(open(file_path).read())

    def load_asts(self, loading):
        return ASTInfo(self.yaml.safe_load(loading))


class AstDumperImpl(object):
    toolset: GolangToolsImpl
    config: GolangToolsConfiguration

    def __init__(self, **kwargs):
        self.config = kwargs.get('config') or GolangToolsConfiguration()
        self.toolset = kwargs.get('toolset') or GolangToolsImpl(**kwargs)

        self.deserializer = YAMLAstDeserializer()

    def dump_ast(self, dumping_package):
        cache_path = os.path.join(self.config.dump_cache_path, dumping_package)
        if self.config.force_update or not os.path.exists(cache_path):
            self.toolset.dump_ast_raw(dumping_package)
        return self.deserializer.load_ast(cache_path)


class Loader(ABC):
    def __init__(self):
        pass

    @abstractmethod
    def handle_function(self, func: ASTInfo.FuncDesc):
        return func


@dataclass
class LoaderConfig(object):
    test: Union[re.Pattern, str]
    target: str
    use: Union[Loader, type, str]


@dataclass
class ModuleConfig(object):
    loaders: List[LoaderConfig] = None


@dataclass
class ParseConfig(object):
    force_update: bool = False


@dataclass
class GolangPackConfig(object):
    name: Optional[str] = None
    version: Optional[str] = None
    description: Optional[str] = None
    package: Optional[str] = None
    src: Optional[str] = None
    output: Optional[str] = None
    module: Optional[ModuleConfig] = None
    parse: Optional[ParseConfig] = None


class GolangPack(object):
    dialect_loader_mapping = dict()

    def __init__(self, config):
        self.config = config  # type: GolangPackConfig

        force_update = False
        if self.config.parse is not None and self.config.parse.force_update:
            force_update = True

        self.go_tool_config = GolangToolsConfiguration(force_update=force_update)
        self.toolset = GolangToolsImpl(config=self.go_tool_config)

    def import_config(self):
        pass

    def once(self):

        if self.config.src is not None:
            self.once_package(self.config.src)

    def hot_update(self):
        pass

    def once_package(self, pkg):
        loaders = self.config.module.loaders
        if loaders is None:
            return
        self.compile_loader()

        ast_info = AstDumperImpl(config=self.go_tool_config, toolset=self.toolset).dump_ast(pkg)

        file_function_mapping = dict()

        for desc in ast_info.functions.items:
            file_function_mapping[desc.pos.file] = desc_list = file_function_mapping.get(desc.pos.file, [])
            desc_list.append(desc)

        for k, funcs in file_function_mapping.items():
            for loader in loaders:
                if loader.test.match(k):
                    self.invoke_loader(loader, k, funcs)
                    break

    def invoke_loader(self, loader, file, funcs: List[ASTInfo.FuncDesc]):
        _ = self
        target_file = self.eval_file(file, loader.target)
        offsets = []

        for func in funcs:
            offsets.append((func.pos.offset, func))
        offsets.sort(key=lambda t: t[0])

        source = cached_io.open_read(file)
        source_pieces = []
        last_index = 0
        for offset, func in offsets:
            source_pieces.append(source[last_index:offset])
            last_index = func.pos.length + func.pos.offset
            stmts = loader.use.handle_function(func)
            source_pieces.append('\n'.join(map(str, stmts)))
        source_pieces.append(source[last_index:])
        dir_name = os.path.dirname(target_file)
        os.makedirs(dir_name, exist_ok=True)
        cached_io.open_write(target_file, ''.join(source_pieces), 'w+')
        self.toolset.go_fmt([target_file])

    def compile_loader(self):
        for loader in self.config.module.loaders:
            if isinstance(loader.test, str):
                loader.test = re.compile(loader.test)
            if isinstance(loader.use, str):
                loader.use = self.dialect_loader_mapping[loader.use]

            if isinstance(loader.use, type):
                loader.use = loader.use()

    def eval_file(self, file, target):
        target_pieces = filter(lambda s: len(s) > 0, re.split(r'((?:\[[^]]*])|@)', target))
        compiled_target_pieces = []
        for target_piece in target_pieces:
            if target_piece == '@':
                if len(compiled_target_pieces) != 0:
                    raise AssertionError("@ must be at first")
                target_piece = current_path
            elif len(target_piece) > 1 and target_piece[0] == '[':
                target_piece = target_piece[1:-1]
                if target_piece == "file-name":
                    target_piece = os.path.basename(file)
                    target_piece, _ = target_piece.split('.', maxsplit=1)
                elif target_piece.startswith('hash:'):
                    target_piece = cached_io.open_digest(file)[:int(target_piece[5:])]
            compiled_target_pieces.append(target_piece)
        file_name = ''.join(compiled_target_pieces)
        assert len(file_name) > 0
        if file_name[0].isalnum():
            target_dir = os.path.abspath(os.path.dirname(file))
            if self.config.output is not None:
                rel_pkg = os.path.relpath(self.config.src, self.config.package)
                maybe_root = target_dir[:len(target_dir) - len(rel_pkg)]
                target_dir = os.path.join(maybe_root, os.path.relpath(self.config.output, self.config.package))
            file_name = os.path.join(target_dir, file_name)
        file_name = os.path.realpath(file_name)
        return file_name

    @staticmethod
    def register_loader(dialect_loader, loader_factory):
        GolangPack.dialect_loader_mapping[dialect_loader] = loader_factory


class MiddleSnake(object):

    @staticmethod
    def From(text):
        return text.replace('_', '-')

    @staticmethod
    def To(text):
        return text.replace('-', '_')


class BigCamel(object):

    @staticmethod
    def From(text):
        return ''.join(x.title() for x in text.split('_'))

    _underscorer1 = re.compile(r'(.)([A-Z][a-z]+)')
    _underscorer2 = re.compile('([a-z0-9])([A-Z])')

    @staticmethod
    def To(text):
        return BigCamel._underscorer2.sub(r'\1_\2', BigCamel._underscorer1.sub(r'\1_\2', text)).lower()


class ConvertStyle(object):

    def __init__(self, value_container=None):
        self.value_container = value_container
        self.fr = None
        self.to = None

    def Values(self, value_container):
        self.value_container = value_container
        return self

    def From(self, fr):
        self.fr = fr
        return self

    def To(self, to):
        self.to = to
        return self

    def Do(self):
        if isinstance(self.value_container, Iterable):
            return [self.to.From(self.fr.To(x)) for x in self.value_container]
        else:
            return self.to.From(self.fr.To(self.value_container))


class StubLoader(Loader):
    @dataclass
    class Context(object):
        fn: ASTInfo.FuncDesc
        context_vars: set
        local_vars: dict
        created_context: bool = False
        created_ok: bool = False
        created_err: bool = False

    def __init__(self):
        super().__init__()

        self.fn_sub_handlers = {
            ASTInfo.AssignExp: self.check_assign_exp,
            ASTInfo.VarDeclExp: self.check_var_decl_exp,
            ASTInfo.CallExp: self.check_call_exp,
        }  # type: Dict[type, lambda _:[_]]

        invoking_stub_handlers = {
            'Context': self.invoking_stub_context,
            'Serve': self.invoking_stub_serve,
            'ServeKeyed': self.invoking_stub_serve_keyed,
        }

        stub_handlers = {
            'GetID': self.stub_get_id,
            'GetIDKeyed': self.stub_get_id_keyed,
            'AbortIf': self.stub_abort_if,
            'Bind': self.stub_bind,
            'Next': self.stub_next,
            'Emit': self.stub_emit,
            'EmitSelf': self.stub_emit_self,
        }

        stub_handlers.update(invoking_stub_handlers)

        promise_handlers = {
            'Then': self.promise_then,
            'Catch': self.promise_catch,
            'Finally': self.promise_finally,
            'ThenDo': self.promise_then_do,
            'CatchDo': self.promise_catch_do,
        }

        self.callee_fn_handlers = {
            'Binder': stub_handlers,
            'Stub': stub_handlers,
            'InvokingStub': invoking_stub_handlers,
            'Promise': promise_handlers,
        }  # type: Dict[str, Dict[str, lambda _:[_]]]

    def handle_function(self, func: ASTInfo.FuncDesc):
        _ = self
        if func.name != "PostSubmission":
            return [func]

        items = []
        context = StubLoader.Context(fn=func, context_vars=set(), local_vars=dict())

        for item in func.body.items:
            t = type(item)
            if t in self.fn_sub_handlers:
                items += self.fn_sub_handlers[t](context, item)
            else:
                items.append(item)
        func.body.items = items
        res = []
        if context.created_context:
            fields = []
            for k in context.context_vars:
                fields.append(f'{k.name} {k.type}')
            fields = '\n'.join(fields)
            res.append(ASTInfo.OpaqueExp.create(f"""type {func.name}Context struct {{\n{fields}\n}}"""))
        res.append(func)
        return res

    def check_assign_exp(self, context: Context, a: ASTInfo.AssignExp):
        if len(a.rhs) != 1:
            return [a]
        rhs = a.rhs[0]
        if not isinstance(rhs, ASTInfo.CallExp):
            return [a]

        return self.handle_stub_call(context, a.lhs, rhs, a)

    def check_call_exp(self, context: Context, a: ASTInfo.CallExp):
        return self.handle_stub_call(context, [], a, a)

    def handle_stub_call(self, context: Context, lhs: List[ASTInfo.Stmt], a: ASTInfo.CallExp, raw):
        # callee = a.callee.body_content
        # if callee.startswith(context.fn.recv.name):
        #     callee = callee[len(context.fn.recv.name) + 1:].split('(', maxsplit=1)[0]
        #     if callee in self.callee_fn_handlers:
        #         return self.callee_fn_handlers[callee](context, lhs, a)

        q = deque()
        callee = a
        while callee:
            if isinstance(callee, ASTInfo.SelectorExp):
                q.append(tuple([callee.name]))
                callee = callee.x
            elif isinstance(callee, ASTInfo.CallExp):
                assert isinstance(callee.callee, ASTInfo.SelectorExp)
                q.append((callee.callee.name, callee.ins))
                callee = callee.callee.x
            else:
                callee = None

        start_stub, current_type = False, None

        res = []
        while len(q):
            p = q.pop()
            if len(p) == 1:
                # todo: search type
                current_type = p[0]
            elif len(p) == 2:
                if current_type is None:
                    raise AssertionError("unknown type of invoker")
                if current_type not in self.callee_fn_handlers:
                    if not start_stub:
                        return [raw]
                    raise AssertionError(f"maybe a bug, want type dict {current_type}")
                start_stub = True
                handlers = self.callee_fn_handlers[current_type]
                current_type, res_sub = handlers[p[0]](context, [] if len(q) else lhs, p[1])
                res.extend(res_sub)
            else:
                raise AssertionError(f"maybe a bug, got {p}")

        return res

    def check_var_decl_exp(self, context: Context, a: ASTInfo.VarDeclExp):
        _ = self
        if len(a.decls) != 1:
            return [a]

        for decl in a.decls:
            # detect new
            if decl.values and len(decl.values) == 1:
                rhs = decl.values[0]
                if not isinstance(rhs, ASTInfo.CallExp):
                    return [a]
                if rhs.callee.body_content != "new":
                    return [a]
                t = rhs.ins[0].body_content
                for name in decl.names:
                    context.local_vars[name] = ASTInfo.Object.create(name, "*" + t)

            # detect null declare
            if decl.type_spec:
                t = decl.type_spec
                for name in decl.names:
                    context.local_vars[name] = ASTInfo.Object.create(name, t)

        return [a]

    # noinspection PyUnresolvedReferences
    def stub_get_id(self, context: Context, lhs: List[ASTInfo.Stmt], rhs: List[ASTInfo.Stmt]):
        assert len(rhs) == 0
        assert len(lhs) == 1

        idName = lhs[0].ident.name.title()

        res = self.must_create_context(context, [])
        res = self.must_create_ok_decl(context, res)
        context.context_vars.add(ASTInfo.Object.create(idName, 'uint'))

        res.append(ASTInfo.OpaqueExp.create(f"context.{idName}, ok = snippet.ParseUint(c, {context.fn.recv.name}.key)"))
        res.append(ASTInfo.OpaqueExp.create("if !ok {\nreturn\n}"))
        return None, res

    # noinspection PyUnresolvedReferences
    def stub_get_id_keyed(self, context: Context, lhs: List[ASTInfo.Stmt], rhs: List[ASTInfo.Stmt]):
        _ = self
        res = []

        return None, res

    # noinspection PyUnresolvedReferences
    def stub_abort_if(self, context: Context, lhs: List[ASTInfo.Stmt], rhs: List[ASTInfo.Stmt]):
        _ = self
        res = []

        return None, res

    # noinspection PyUnresolvedReferences
    def stub_next(self, context: Context, lhs: List[ASTInfo.Stmt], rhs: List[ASTInfo.Stmt]):
        _ = self
        res = []

        return None, res

    # noinspection PyUnresolvedReferences
    def stub_bind(self, context: Context, lhs: List[ASTInfo.Stmt], rhs: List[ASTInfo.Stmt]):
        assert len(lhs) == 0

        res = []
        for binding in rhs:
            bc = binding.body_content
            bc.strip('&').strip()
            if bc not in context.local_vars:
                raise KeyError(f'can not bind {repr(binding)}')
            res = self.must_create_context(context, res)
            local_var = context.local_vars[bc]
            cc = bc.title()
            context.context_vars.add(ASTInfo.Object.create(cc, local_var.type))
            res.append(ASTInfo.OpaqueExp.create(f"context.{cc} = {bc}"))
            res.append(ASTInfo.OpaqueExp.create(f"if !snippet.BindRequest(c, {bc}) {{\nreturn\n}}"))
        return None, res

    def invoking_stub_context(self, context: Context, lhs: List[ASTInfo.Stmt], rhs: List[ASTInfo.Stmt]):
        _ = self
        res = []

        return 'InvokingStub', res

    def invoking_stub_serve(self, context: Context, lhs: List[ASTInfo.Stmt], rhs: List[ASTInfo.Stmt]):
        _ = self
        res = []

        return 'Promise', res

    def invoking_stub_serve_keyed(self, context: Context, lhs: List[ASTInfo.Stmt], rhs: List[ASTInfo.Stmt]):
        _ = self
        res = []

        return 'Promise', res

    def stub_emit_self(self, context: Context, lhs: List[ASTInfo.Stmt], rhs: List[ASTInfo.Stmt]):
        _ = self
        assert len(lhs) == 0
        res = []
        return 'Promise', res

    def stub_emit(self, context: Context, lhs: List[ASTInfo.Stmt], rhs: List[ASTInfo.Stmt]):
        _ = self
        assert len(lhs) == 0
        res = []
        return 'Promise', res

    def promise_then(self, context: Context, lhs: List[ASTInfo.Stmt], rhs: List[ASTInfo.Stmt]):
        _ = self
        assert len(lhs) == 0
        res = []
        return 'Promise', res

    def promise_catch(self, context: Context, lhs: List[ASTInfo.Stmt], rhs: List[ASTInfo.Stmt]):
        _ = self
        assert len(lhs) == 0
        res = []
        return 'Promise', res

    def promise_finally(self, context: Context, lhs: List[ASTInfo.Stmt], rhs: List[ASTInfo.Stmt]):
        _ = self
        assert len(lhs) == 0
        res = []
        return 'Promise', res

    def promise_then_do(self, context: Context, lhs: List[ASTInfo.Stmt], rhs: List[ASTInfo.Stmt]):
        _ = self
        assert len(lhs) == 0
        res = []
        return 'Promise', res

    def promise_catch_do(self, context: Context, lhs: List[ASTInfo.Stmt], rhs: List[ASTInfo.Stmt]):
        _ = self
        assert len(lhs) == 0
        res = []
        return 'Promise', res

    def must_create_context(self, context: Context, res):
        _ = self
        if not context.created_context:
            context.created_context = True
            res.append(
                ASTInfo.OpaqueExp.create(f"var context {context.fn.name}Context")
            )
        return res

    def must_create_ok_decl(self, context: Context, res):
        _ = self
        if not context.created_ok:
            context.created_ok = True
            res.append(ASTInfo.OpaqueExp.create(f"var ok bool"))
        return res

    def must_create_err_decl(self, context: Context, res):
        _ = self
        if not context.created_err:
            context.created_err = True
            res.append(ASTInfo.OpaqueExp.create(f"var err error"))
        return res


if __name__ == '__main__':
    GolangPack.register_loader('stub-loader', StubLoader)

    golang_pack = GolangPack(GolangPackConfig(
        name='boj-v6',
        version='v0.5.0',
        description='golang pack test config',
        package='github.com/Myriad-Dreamin/boj-v6',
        src='github.com/Myriad-Dreamin/boj-v6/cmd/generate/controller_binding/inner/model',
        output='github.com/Myriad-Dreamin/boj-v6/app/generated_controller',
        module=ModuleConfig(
            loaders=[
                LoaderConfig(test=re.compile(r'.*?.stub.go$'), target='[file-name].gen.go', use='stub-loader'),
            ]
        ),
        parse=ParseConfig(
            force_update=True,
        ),
    ))

    golang_pack.once()
