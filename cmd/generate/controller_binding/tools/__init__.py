import json
import logging
import os
import re
import subprocess
from dataclasses import dataclass
from datetime import datetime
from typing import Union, List, Optional, Tuple

from exception import GolangToolInvokeError
from go_ast.persistence import JSONAstDeserializer


@dataclass
class GolangToolsConfig(object):
    dump_tool_package: str = 'github.com/Myriad-Dreamin/golang-pack/ast-dump'
    dump_cache_path: str = '.cache/ast_dump'
    force_update: bool = False
    verbose: bool = False


def use_cache_json(cache_path):
    def wrap_func(f):
        def wrap(self, *args, **kwargs):
            cached, cached_path = None, os.path.join(self.config.dump_cache_path, cache_path)
            if os.path.exists(cached_path):
                cached = json.load(open(cached_path, 'r+'))
                if cached['golang_pack_cache_expired'] < datetime.now().timestamp():
                    cached = None

            if not cached:
                cached = {
                    'payload': f(self, *args, **kwargs),
                    'golang_pack_cache_expired': datetime.now().timestamp() + 86400000,
                }
                json.dump(cached, open(cached_path, 'w'))
            return cached['payload']

        return wrap

    return wrap_func


class GolangToolsImpl(object):
    class DefaultRunner(object):
        def __init__(self, verbose=False):
            self.verbose = verbose

        def run_command(self, cmd: Union[List[str], str], timeout: Optional[float] = None) -> Tuple[int, str, str]:
            if not isinstance(cmd, str):
                cmd = ' '.join(cmd)
            if self.verbose:
                print(cmd)
            process = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            process.wait(timeout=timeout)
            code = process.returncode
            stdout, stderr = process.communicate()

            if isinstance(stdout, bytes):
                stdout = stdout.decode()

            if isinstance(stderr, bytes):
                stderr = stderr.decode()

            return code, stdout, stderr

    runner: DefaultRunner
    config: GolangToolsConfig

    def __init__(self, runner=None, config=None, **_):
        self.config = config or GolangToolsConfig()
        self.runner = runner or GolangToolsImpl.DefaultRunner(self.config.verbose)
        self.env = self.go_env_obj()

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

    def go_env(self, cmd: Union[List[str], str], timeout: Optional[float] = None) -> str:
        return self.run_command(['go env'] + cmd, timeout)

    @use_cache_json('.go-env.json')
    def go_env_obj(self, timeout: Optional[float] = None) -> dict:
        return json.loads(self.go_env(['-json'], timeout=timeout))

    def go_version(self, timeout: Optional[float] = None) -> str:
        return self.run_command(['go version'], timeout)

    @use_cache_json('.go-version.json')
    def go_version_repr(self, timeout: Optional[float] = None) -> List[int]:
        version = self.go_version(timeout=timeout)
        version = list(map(int, version.split()[2][2:].split('.')))
        if len(version) == 3:
            main_version, major, minor = version
        else:
            main_version, major = version
            minor = None
        return [main_version, major, minor]

    def is_golang_using_go_module(self):
        version = self.go_version_repr()
        go_module_env = self.env.get('GO111MODULE')
        if version > [1, 13, -1]:
            if go_module_env == 'off':
                return False
            return True
        else:
            if go_module_env == 'on':
                return True
        return False

    def read_golang_module_name(self, module_path):
        _ = self
        module_file = f'{module_path}/go.mod'
        if not os.path.exists(module_file):
            return None
        module_name = re.search(r'module\S*([^\n]*)', open(module_file, 'r').read()).group(1)
        return module_name and module_name.strip()

    def dump_ast_raw(self, dumping_package):
        self.go_run([self.config.dump_tool_package,
                     dumping_package, self.config.dump_cache_path])
        return


class AstDumperImpl(object):
    toolset: GolangToolsImpl
    config: GolangToolsConfig

    def __init__(self, **kwargs):
        self.config = kwargs.get('config') or GolangToolsConfig()
        self.toolset = kwargs.get('toolset') or GolangToolsImpl(**kwargs)

        self.deserializer = JSONAstDeserializer()

    def dump_ast(self, dumping_package):
        cache_path = os.path.join(self.config.dump_cache_path, dumping_package)
        if self.config.force_update or not os.path.exists(cache_path):
            self.toolset.dump_ast_raw(dumping_package)
        return self.deserializer.load_ast(cache_path)
