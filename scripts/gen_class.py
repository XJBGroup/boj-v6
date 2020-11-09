import random
import re
from typing import Union, Callable, Any


def indented(indenting_lines, base_indent=4, indent=0):
    for i, line in enumerate(indenting_lines):
        if isinstance(line, list):
            indenting_lines[i] = indented(line, base_indent, indent + base_indent)
    indents = ('\n' + (' ' * indent))
    if indent:
        return '    ' + indents.join(indenting_lines)
    return indents.join(indenting_lines)


class Stream(object):
    def __init__(self, s):
        self.s = None
        self.ref(s)

    def ref(self, s):
        self.s = s
        return self

    def __getitem__(self, item_key):
        return self.s[item_key]

    def __contains__(self, contains_key):
        return contains_key in self.s

    def forward(self, n):
        self.s = self.s[n:]

    def strip(self):
        self.s = self.s.strip()
        return self

    def split(self, ch=None, max_split=None):
        if max_split is None:
            return self.s.split(ch)
        return self.s.split(ch, max_split)


generic_annotation_factory = None  # type: Union[None, Callable[['Stream'], Any]]


class ListTypeAnnotation(object):
    x_type = 'ListTypeAnnotation'
    python_type = 'list'

    def __init__(self, type_str, inner_operator=None, _can_instance=None):
        self.type_str = type_str
        self.inner_operator = inner_operator
        self.can_instance = False

    @staticmethod
    def from_stream(s: Stream):
        return ListTypeAnnotation(s.s, generic_annotation_factory(s))

    @property
    def as_dict_constructor(self):
        return 'g_a_cpc_x'

    def as_dict_constructor_ph(self, lk):
        cv = f'g_a_cpc_x_{hex(random.randint(0, 1 << 31))[2:]}'
        lph = [f'if isinstance({lk}, list):']
        rv = f'{cv}_value'
        lph.append([
            f'{cv} = list()',
            f'for {rv} in {lk}:',
            self.inner_operator.as_dict_constructor_ph(rv) + [
                f'{cv}.append({self.inner_operator.as_dict_constructor})',
            ],
            f'g_a_cpc_x = {cv}'
        ])
        return lph


class UnionTypeAnnotation(object):
    x_type = 'UnionTypeAnnotation'
    python_type = 'any'

    def __init__(self, inner_operators=None):
        self.inner_operators = inner_operators or []
        self.can_instance = False

    @staticmethod
    def from_stream(s: Stream):
        uu = generic_annotation_factory(s)
        s.strip()
        if s[0] != ',':
            raise ValueError("want a comma")
        s.forward(1)
        s.strip()
        uv = generic_annotation_factory(s)

        if uv.python_type in ['list', 'dict']:
            uu, uv = uv, uu
        return UnionTypeAnnotation([uu, uv])

    @property
    def as_dict_constructor(self):
        return 'g_a_cpc_x'

    def as_dict_constructor_ph(self, uk):
        uph = ['g_a_cpc_x = None']
        if len(self.inner_operators) == 0:
            return uph
        cv = f'g_a_cpc_x_{hex(random.randint(0, 1 << 31))[2:]}'
        hv = f'{cv}_handling'
        uph += [f'{hv} = {uk}']
        i = iter(self.inner_operators)
        uph += next(i).as_dict_constructor_ph(hv)
        for op in i:
            uph += ['if g_a_cpc_x is None:', op.as_dict_constructor_ph(hv)]
        return uph


class DictTypeAnnotation(object):
    x_type = 'DictTypeAnnotation'
    python_type = 'dict'

    def __init__(self, dict_key, dict_v):
        self.dict_key = dict_key
        self.dict_value = dict_v
        self.can_instance = True

    @staticmethod
    def from_stream(s: Stream):
        u = generic_annotation_factory(s)
        s.strip()
        if s[0] != ',':
            raise ValueError("want a comma")
        s.forward(1)
        s.strip()
        gv = generic_annotation_factory(s)
        return DictTypeAnnotation(u, gv)

    @property
    def as_dict_constructor(self):
        return 'g_a_cpc_x'

    def as_dict_constructor_ph(self, dk):
        dph = []
        cv = f'g_a_cpc_x_{hex(random.randint(0, 1 << 31))[2:]}'
        dph.append(f'if isinstance({dk}, dict):')
        rv = f'{cv}_value'
        dph.append([
            f'{cv} = dict()',
            f'for {rv} in {dk}.items():',
            self.dict_value.as_dict_constructor_ph(f'{rv}[1]') + [
                f'{cv}[{rv}[0]] = ({self.dict_value.as_dict_constructor})',
            ],
            f'g_a_cpc_x = {cv}'
        ])
        return dph


basic_type = ['str', 'list', 'int', 'float', 'dict', 'bool', 'object']


class DefaultAnnotationOperation(object):
    x_type = 'DefaultAnnotationOperation'
    python_type = 'any'

    def __init__(self, type_str):
        self.type_str = type_str
        if self.type_str in basic_type:
            self.python_type = self.type_str

    @property
    def as_dict_constructor(self):
        return 'g_a_cpc_x'

    def as_dict_constructor_ph(self, dk):
        if self.python_type == 'any':
            return [f'g_a_cpc_x = {self.type_str}.from_dict({dk})']
        if self.python_type == 'object':
            return [f'g_a_cpc_x = {dk}']
        return [f'g_a_cpc_x = {self.type_str}({dk})']


class TypeAnnotation(object):
    x_type = 'TypeAnnotation'
    python_type = 'any'

    annotation_stream_factories = {
        'List': ListTypeAnnotation.from_stream,
        'Union': UnionTypeAnnotation.from_stream,
        'Dict': DictTypeAnnotation.from_stream,
    }

    def __init__(self, type_str, operator=None, can_instance=False):
        self.type_str = type_str
        self.operator = operator or DefaultAnnotationOperation(self.type_str)
        self.can_instance = can_instance
        self.python_type = self.operator.python_type

    @property
    def as_dict_constructor(self):
        return self.operator.as_dict_constructor

    def as_dict_constructor_ph(self, tk):
        return self.operator.as_dict_constructor_ph(tk)

    @staticmethod
    def from_stream(s: Stream) -> Any:
        if '[' in s:
            g = s.split('[', 1)
            an_type, nc = map(lambda fs: fs.strip(), g)
            if ',' not in an_type:
                return TypeAnnotation.annotation_stream_factories[an_type.strip()](s.ref(nc[:-1]))
        tt = re.split(r'[ ,\[]', s.s, 1)
        s.forward(len(tt[0]))
        return TypeAnnotation(tt[0])

    @staticmethod
    def from_str(s: str):
        return TypeAnnotation.from_stream(Stream(s))


# noinspection PyRedeclaration
generic_annotation_factory = TypeAnnotation.from_stream

python_keys = ['and', 'as', 'assert', 'break', 'class', 'continue', 'def',
               'del', 'elif', 'else', 'except', 'exec', 'finally', 'for',
               'from', 'global', 'if', 'import', 'in', 'is', 'lambda',
               'not', 'or', 'pass', 'print', 'raise', 'return', 'try',
               'while', 'with', 'yield']


class Prop(object):
    def __init__(self, dict_field, field_name, t, type_str):
        self.dict_field = dict_field
        self.field_name = field_name
        self.t = t
        self.type_str = type_str


default_values = {
    'str': "''",
    'int': "0",
    'float': "0.",
    'bool': "False",
    'list': "[]",
    'dict': "dict()",
}

def main(target_path, class_defs):
    f = open(target_path, 'w', encoding='utf8')
    lines = ['from typing import *']

    for class_def in reversed(re.split('---+\n', class_defs)):
        lines.append('')
        class_name, spec = class_def.strip().split('\n', 1)
        class_name, inherit = map(lambda x: x.strip(), class_name.split(':'))
        if len(inherit) == 0:
            inherit = 'object'

        class_props = dict()

        for spec_item in spec.split('\n'):
            spec_item = spec_item.strip()
            if len(spec_item) == 0:
                continue
            k, v = map(lambda x: x.strip(), spec_item.split(':', 1))

            if ',' in k:
                k, fn = map(lambda x: x.strip(), k.split(','))
            else:
                fn = k
            prop = Prop(k, fn, TypeAnnotation.from_str(v), v)
            if k in python_keys:
                prop.dict_field = k + '_'
            class_props[k] = prop

        lines.append(f'class {class_name}({inherit}):')
        class_props_lines = []
        lines.append(class_props_lines)
        for k, v in class_props.items():
            class_props_lines.append(f"{v.dict_field}: Union[None, '{v.type_str}']")

        def create_method(insert_point, signature, decorators=None):
            method, method_body = [], []
            insert_point.append(method)
            if decorators:
                method.extend(decorators)
            method.append(f'def {signature}:')
            method.append(method_body)
            return method_body

        def create_if(insert_point, cond):
            insert_point.append(f'if {cond}:')
            if_block = []
            insert_point.append(if_block)
            return if_block

        body = create_method(lines, '__init__(self)')
        for k, v in class_props.items():
            body.append(f'self.{v.dict_field} = None')

        decs = ['@staticmethod']
        body = create_method(lines, f"from_dict(d: dict) -> '{class_name}'", decs)
        body.append(f'ci = {class_name}()')
        body.extend([
            'if d is None:',
            [
                'return ci'
            ]
        ])
        for k, v in class_props.items():
            # it is Optional[T]
            body.append(f"ci.{v.dict_field} = getting = d.get('{v.field_name}', None)")
            ib = create_if(body, 'getting is not None')
            ib.append('g_a_cpc_x = None')
            ph = v.t.as_dict_constructor_ph('getting')
            if ph:
                ib.extend(ph)
            ib.append(f"ci.{v.dict_field} = {v.t.as_dict_constructor}")
            # dv = default_values.get(v.t.python_type, 'None')
        body.append('return ci')

    lines.append('')

    code = indented(lines, 4)

    def do_optimize(c, optimized_block):
        codes = []
        last_xx = 0
        for o in optimized_block:
            x, y = o
            optimized_l, optimized_r = x
            codes.append(c[last_xx:optimized_l])
            codes.append(y)
            last_xx = optimized_r
        codes.append(c[last_xx:])
        return ''.join(codes)

    optimized = []
    for m in re.finditer(r'([ ]*)g_a_cpc_x\s*=\s*(\S*)[^ ]*([ ]*)(\S*)\s*=\s*g_a_cpc_x(\s+)', code):
        i1, rhs, i2, lhs, i3 = m[1], m[2], m[3], m[4], m[5]
        if len(i1) == len(i2):
            optimized.append((m.span(), f'{i1}{lhs} = {rhs}{i3}'))
    code = do_optimize(code, optimized)

    optimized = []
    for m in re.finditer(r'([ ]*)g_a_cpc_x\s*=\s*\S*[^ ]*([ ]*)(g_a_cpc_x\s*=\s*\S*)(\s+)', code):
        i1, i2, lhs, i3 = m[1], m[2], m[3], m[4]
        if len(i1) == len(i2):
            optimized.append((m.span(), f'{i1}{lhs}{i3}'))
    code = do_optimize(code, optimized)

    f.write(code)
    f.close()
