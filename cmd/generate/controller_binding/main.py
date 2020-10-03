import subprocess
from typing import Union, List, Optional, Tuple, Dict
import logging
import os

current_path = os.getcwd()
def simplify_path(path):
    return os.path.relpath(path, current_path).replace('\\', '/')


class GolangToolInvokeError(Exception):
    code: int
    msg: str

    def __init__(self, *args, code, error_dump):
        self.code = code
        self.msg = error_dump


class GolangToolsConfiguration(object):
    dump_tool_package: str = 'github.com/Myriad-Dreamin/boj-v6/cmd/generate/controller_binding/ast-dump'
    dump_cache_path: str = '.cache/ast_dump'


class GolangToolsImpl(object):

    class DefaultRunner(object):

        @staticmethod
        def run_command(cmd: Union[List[str], str], timeout: Optional[float] = None) -> Tuple[int, str, str]:
            if not isinstance(cmd, str):
                cmd = ' '.join(cmd)
            print(cmd)
            process = subprocess.Popen(cmd)
            process.wait(timeout=timeout)
            code = process.returncode
            stdout, stderr = process.communicate()
            return code, stdout, stderr

    runner: DefaultRunner
    config: GolangToolsConfiguration

    def __init__(self, runner=DefaultRunner, config=None, **kwargs):
        self.runner = runner
        self.config = config or GolangToolsConfiguration()

    def run_command(self, cmd: Union[List[str], str], timeout: Optional[float] = None) -> str:
        c, o, e = self.runner.run_command(cmd, timeout)
        if c != 0:
            raise GolangToolInvokeError(code=c, error_dump=e)
        if e is not None:
            logging.warning(e)
        return o

    def go_run(self, cmd: Union[List[str], str], timeout: Optional[float] = None) -> str:
        return self.run_command(['go run'] + cmd, timeout)

    def dump_ast_raw(self, dumping_package):
        self.go_run([self.config.dump_tool_package,
                     dumping_package, self.config.dump_cache_path])
        return


def use_file_mapping(maybe_wrapping_cls, *args):
    if isinstance(maybe_wrapping_cls, str):
        args = (maybe_wrapping_cls, ) + args
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
                        else:
                            p.use_file_mapping(mp)
        return WrappedClass
    if isinstance(maybe_wrapping_cls, str):
        return use_file_mapping_2step
    return use_file_mapping_2step(maybe_wrapping_cls)


def use_file_mapping_not_report_error(target, mp):
    target.use_file_mapping(mp)


class ASTInfo(object):

    class Object(object):
        def __init__(self, doc_tree):
            self.name = doc_tree['n']
            self.type = doc_tree['t']
        
        def __repr__(self):
            return f'{self.name}(T:{self.type})'

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
    
    class Stmt(object):
        def __init__(self, doc_tree):
            self.pos = ASTInfo.FilePos(doc_tree['p'])

    @use_file_mapping('items')
    class FuncBody(Stmt):
        def __init__(self, doc_tree):
            super().__init__(doc_tree)
            self.items = list(map(ASTInfo.Stmt, doc_tree['b']))
        
        def __repr__(self):
            return ':'.join(map(str, [self.pos.line, self.pos.column, self.pos.length]))

    @use_file_mapping('items')
    class FuncDescs(object):
        
        @use_file_mapping
        class Item(object):
            def __init__(self, doc_tree):
                self.pos = ASTInfo.FilePos(doc_tree['p'])
                self.recv = ASTInfo.Object(doc_tree['r'])
                self.name = doc_tree['n']
                self.ins = list(map(ASTInfo.Object, doc_tree['in']))
                self.outs = list(map(ASTInfo.Object, doc_tree['out']))
                self.body = ASTInfo.FuncBody(doc_tree['body'])
            
            def __repr__(self):
                return f"<{self.recv.type}.{self.name} at {repr(self.pos)}, body={repr(self.body)}>"

        def __init__(self, doc_tree):
            self.items = list(map(ASTInfo.FuncDescs.Item, doc_tree))
    
        def __repr__(self):
            items = self.items
            if len(items) > 20:
                return '\n'.join(map(repr, items[:5])) + '\n...'
            return '\n'.join(map(repr, items))

    def __init__(self, doc_tree):
        self.file_mapping = ASTInfo.FilesMapping(doc_tree['file_mapping'])
        self.imports = ASTInfo.ImportStmts(doc_tree['imports'])
        self.functions = ASTInfo.FuncDescs(doc_tree['functions'])
        use_file_mapping_not_report_error(self.imports, self.file_mapping)
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
        if not os.path.exists(cache_path):
            self.toolset.dump_ast_raw(dumping_package)
        return self.deserializer.load_ast(cache_path)


if __name__ == '__main__':
    AstDumper = AstDumperImpl()
    ast_info = AstDumper.dump_ast('github.com/Myriad-Dreamin/boj-v6/cmd/generate/controller_binding/inner/model')
    print(ast_info)
    # for desc in ast_info.functions.items:
    #     f = open(desc.pos.file).read().replace('\n', '\n\r')
    #     print(f[desc.pos.offset:desc.pos.offset+desc.pos.length])
