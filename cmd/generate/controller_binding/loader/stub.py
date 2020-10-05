from collections import deque
from dataclasses import dataclass
from typing import List, Dict, Optional

from config import Loader
from go_ast import FuncDesc, AssignExp, VarDeclExp, CallExp, OpaqueExp, Stmt, SelectorExp, Object, IdentExp, BinaryExp, \
    UnaryExp


@dataclass
class Context(object):
    fn: FuncDesc
    extended_methods: dict
    context_vars: dict
    local_vars: dict
    service_methods: dict
    insert_points: dict
    insert_items_list: List[Dict]
    stmt_index: int = 0
    exception_ttl: int = 1
    last_exception_index: int = Optional[None]
    last_exception_ok_index: int = Optional[None]
    created_context: bool = False


def escape(s):
    return s.replace('"', '\\"')


def do_capitalize(name):
    return name[0].upper() + name[1:]


class StubLoader(Loader):

    def __init__(self):
        super().__init__()

        # noinspection PyUnusedLocal
        fn_sub_handler_type = type(self.check_assign_exp)

        self.fn_sub_handlers = {
            AssignExp: self.check_assign_exp,
            VarDeclExp: self.check_var_decl_exp,
            CallExp: self.check_call_exp,

        }  # type: Dict[type, fn_sub_handler_type]

        # noinspection PyUnusedLocal
        chained_stub_handler_type = type(self.invoking_stub_context)

        invoking_stub_handlers = {
            'Context': self.invoking_stub_context,
            'Serve': self.invoking_stub_serve,
            'ServeKeyed': self.invoking_stub_serve_keyed,
        }  # type: Dict[str, chained_stub_handler_type]

        stub_handlers = {
            'GetID': self.stub_get_id,
            'GetIDKeyed': self.stub_get_id_keyed,
            'AbortIf': self.stub_abort_if,
            'AbortIfHint': self.stub_abort_if_hint,
            'Bind': self.stub_bind,
            'OnErr': self.stub_on_error,
        }  # type: Dict[str, chained_stub_handler_type]

        stub_handlers.update(invoking_stub_handlers)

        promise_handlers = {
            'Then': self.promise_then,
            'Catch': self.promise_catch,
            'Finally': self.promise_finally,
            'ThenRef': self.promise_then_ref,
            'CatchRef': self.promise_catch_ref,
            'FinallyRef': self.promise_finally_ref,
        }  # type: Dict[str, chained_stub_handler_type]

        self.callee_fn_handlers = {
            'Binder': stub_handlers,
            'Stub': stub_handlers,
            'InvokingStub': invoking_stub_handlers,
            'Promise': promise_handlers,
        }  # type: Dict[str, Dict[str, chained_stub_handler_type]]

        self.stub_variables = {
            'Ok': Object.create('ok', 'bool'),
            'Err': Object.create('err', 'error'),
            'Int64': Object.create('stubInt64', 'int64'),
            'Int32': Object.create('stubInt32', 'int32'),
            'Int16': Object.create('stubInt16', 'int16'),
            'Int8': Object.create('stubInt8', 'int8'),
            'Int': Object.create('stubInt', 'int'),
            'Uint64': Object.create('stubUint64', 'uint64'),
            'Uint32': Object.create('stubUint32', 'uint32'),
            'Uint16': Object.create('stubUint16', 'uint16'),
            'Uint8': Object.create('stubUint8', 'uint8'),
            'Uint': Object.create('stubUint', 'uint'),
        }

    def handle_function(self, func: FuncDesc):
        items = []
        opaque_items_list = []
        context = Context(
            fn=func,
            context_vars=dict(), local_vars=dict(), insert_items_list=list(),
            extended_methods=dict(), service_methods=dict(), insert_points=dict())
        for i, item in enumerate(func.body.items):
            context.insert_items_list.append(context.insert_points)
            context.stmt_index = i
            new_items = []
            self.handle_stmt(item, context, new_items)
            opaque_items_list.append(filter(lambda o_item: isinstance(o_item, OpaqueExp), new_items))
            items.extend(new_items)
            context.insert_points = dict()

        for opaque_items, insert_points in zip(opaque_items_list, context.insert_items_list):
            for x_item in opaque_items:
                for k, v in insert_points.items():
                    x_item.opaque = x_item.opaque.replace(k, '\n'.join(map(str, v)))
        func.body.items = items
        res = []
        if context.created_context:
            fields = []
            for k in context.context_vars.values():
                fields.append(f'{k.name} {k.type}')
            fields = '\n'.join(fields)
            res.append(OpaqueExp.create(f"""type {func.name}Context struct {{\n{fields}\n}}"""))

        interface_sub_services = []
        for service_name, value in context.service_methods.items():
            _, _, key = value
            method_name = self.resolve_service_method_handler(context, service_name + 'Service', context.fn.name, key)
            service_name = f"{func.name}{key or 'Self'}Service"
            res.append(OpaqueExp.create(f"""\ntype {service_name} interface {{
{method_name}(context *{func.name}Context) (err error)\n}}"""))
            interface_sub_services.append(service_name)

        interface_sub_services = '\n'.join(interface_sub_services)
        res.append(OpaqueExp.create(f"""\ntype {func.name}Service interface {{\n{interface_sub_services}\n}}"""))
        res.extend(context.extended_methods.values())
        res.append(func)
        return res

    def handle_stmt(self, item, context, items):
        t = type(item)
        if t in self.fn_sub_handlers:
            items += self.fn_sub_handlers[t](context, item)
        else:
            items.append(item)
        return

    def check_replace_stub_variable_item(self, context: Context, lhs, res):
        if isinstance(lhs, SelectorExp):
            if lhs.x.body_content != context.fn.recv.name:
                return lhs

            maybe_stub_variable = lhs.name
            if maybe_stub_variable not in self.stub_variables:
                return lhs

            decl = self.stub_variables[maybe_stub_variable]
            self.must_create_decl(context, decl, res)
            return decl
        return lhs

    def check_replace_stub_variable(self, context: Context, xhs: List[Stmt], res):
        handling = None
        for i, lhs in enumerate(xhs):
            decl = self.check_replace_stub_variable_item(context, lhs, res)
            if lhs == decl:
                continue
            xhs[i] = decl
            if decl.name == 'err':
                if handling:
                    raise KeyError(f"already find a catchable object {handling}")
                handling = 'E'
            elif decl.name == 'ok':
                if handling:
                    raise KeyError(f"already find a catchable object {handling}")
                handling = 'O'
        return handling

    def check_replace_stub_variable_exp(self, context: Context, xhs: Stmt, res):
        if isinstance(xhs, BinaryExp):
            xhs.lhs = self.check_replace_stub_variable_exp(context, xhs.lhs, res)
            xhs.rhs = self.check_replace_stub_variable_exp(context, xhs.rhs, res)
            return xhs

        if isinstance(xhs, UnaryExp):
            xhs.lhs = self.check_replace_stub_variable_exp(context, xhs.lhs, res)
            return xhs

        if isinstance(xhs, CallExp):
            self.check_replace_stub_variable(context, xhs.ins, res)
            return xhs

        return self.check_replace_stub_variable_item(context, xhs, res)

    def check_assign_exp(self, context: Context, a: AssignExp):
        res = []
        handling = self.check_replace_stub_variable(context, a.lhs, res)
        self.check_replace_stub_variable(context, a.rhs, res)
        if handling:
            if handling == 'E':
                return self.global_exception_handler(context, a, res)
            elif handling == 'O':
                return self.global_bool_handler(context, a, res)
            else:
                raise AssertionError("catchable object should be either error or bool")

        if len(a.rhs) != 1:
            return res + [a]

        rhs = a.rhs[0]
        if not isinstance(rhs, CallExp):
            return res + [a]

        return res + self.handle_stub_call(context, a.lhs, rhs, [a])

    def check_call_exp(self, context: Context, a: CallExp):
        res = []
        self.check_replace_stub_variable(context, a.ins, res)
        return res + self.handle_stub_call(context, [], a, [a])

    def resolve_service_method_handler(self, _: Context, _service, method, key: Optional[str]):
        _ = self
        return key or method

    def resolve_service_handler(self, context: Context, service, method, key: Optional[str]):
        method = self.resolve_service_method_handler(context, service, method, key)
        return f'{context.fn.recv.name}.resolver.Require(' \
               f'{context.fn.recv.name}.serviceName).({service}).{method}'

    def handle_stub_call(self, context: Context, lhs: List[Stmt], a: CallExp, raw):
        q = deque()
        callee = a
        while callee:
            if isinstance(callee, SelectorExp):
                q.append(tuple([callee.name]))
                callee = callee.x
            elif isinstance(callee, CallExp):
                if not isinstance(callee.callee, SelectorExp):
                    return raw
                q.append((callee.callee.name, callee.ins))
                callee = callee.callee.x
            elif isinstance(callee, IdentExp):
                if callee.body_content != context.fn.recv.name:
                    return raw
                callee = None
            else:
                raise TypeError(f'unknown callee ast, content: {callee.body_content} type: {type(callee)}')

        start_stub, current_type = False, 'Stub'

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
                        return raw
                    raise AssertionError(f"maybe a bug, want type dict {current_type}")
                handlers = self.callee_fn_handlers[current_type]
                if not start_stub and p[0] not in handlers:
                    return raw
                start_stub = True
                if p[0] not in handlers:
                    raise KeyError(f'not registered stub type in scope {current_type}, want handler name is {p[0]}')
                current_type, res_sub = handlers[p[0]](context, [] if len(q) else lhs, p[1])
                res.extend(res_sub)
            else:
                raise AssertionError(f"maybe a bug, got {p}")

        return res

    def check_var_decl_exp(self, context: Context, a: VarDeclExp):
        _ = self
        if len(a.decls) != 1:
            return [a]

        for decl in a.decls:
            # detect new
            if decl.values and len(decl.values) == 1:
                rhs = decl.values[0]
                if not isinstance(rhs, CallExp):
                    return [a]
                if rhs.callee.body_content != "new":
                    return [a]
                t = rhs.ins[0].body_content
                for name in decl.names:
                    context.local_vars[name] = Object.create(name, "*" + t)

            # detect null declare
            if decl.type_spec:
                t = decl.type_spec
                for name in decl.names:
                    context.local_vars[name] = Object.create(name, t)

        return [a]

    # noinspection PyUnresolvedReferences
    def stub_do_get_id(self, context: Context, lhs: List[Stmt], k: str):
        raw = lhs[0].ident.name
        id_name = do_capitalize(raw)

        res = self.must_create_context(context, [])
        res = self.must_create_ok_decl(context, res)
        context.context_vars[id_name] = Object.create(id_name, 'uint')
        context.local_vars[raw] = Object.create(raw, 'uint')

        res.append(OpaqueExp.create(f"var {raw} uint"))
        res.append(OpaqueExp.create(f"{raw}, ok = snippet.ParseUint(c, {k})"))
        res.append(OpaqueExp.create("if !ok {\nreturn\n}"))
        res.append(OpaqueExp.create(f"stubContext.{id_name} = {raw}"))

        return None, res

    # noinspection PyUnresolvedReferences
    def stub_get_id(self, context: Context, lhs: List[Stmt], rhs: List[Stmt]):
        assert len(rhs) == 0
        assert len(lhs) == 1
        return self.stub_do_get_id(context, lhs, f'{context.fn.recv.name}.key')

    # noinspection PyUnresolvedReferences
    def stub_get_id_keyed(self, context: Context, lhs: List[Stmt], rhs: List[Stmt]):
        assert len(rhs) == 1
        assert len(lhs) == 1
        return self.stub_do_get_id(context, lhs, rhs[0].body_content)

    # noinspection PyUnresolvedReferences
    def stub_do_abort_if_hint(self, context: Context, lhs: List[Stmt], assertion: Stmt, rhs: List[Stmt], hint: str):
        assert len(lhs) == 0
        _ = self

        rest_args = ','.join(map(str, rhs))

        res = []
        assertion = str(self.check_replace_stub_variable_exp(context, assertion, res))
        if len(res) != 0:
            raise LookupError(f"input of the assertion is not defined: {res}")

        return None, [
            OpaqueExp.create(
                f"""if {assertion} {{
snippet.DoReportHintRaw(c, "Assertion Failed: want {escape(assertion)}", {hint}, {rest_args})\nreturn\n}}""")]

    # noinspection PyUnresolvedReferences
    def stub_abort_if(self, context: Context, lhs: List[Stmt], rhs: List[Stmt]):
        assert len(rhs) > 0
        return self.stub_do_abort_if_hint(context, lhs, rhs[0], rhs[1:], '0xbad')

    # noinspection PyUnresolvedReferences
    def stub_abort_if_hint(self, context: Context, lhs: List[Stmt], rhs: List[Stmt]):
        assert len(rhs) > 1
        return self.stub_do_abort_if_hint(context, lhs, rhs[0], rhs[2:], rhs[1].body_content)

    # noinspection PyUnresolvedReferences
    def stub_bind(self, context: Context, lhs: List[Stmt], rhs: List[Stmt]):
        assert len(lhs) == 0

        res = []
        for binding in rhs:
            bc = binding.body_content
            bc.strip('&').strip()
            if bc not in context.local_vars:
                raise KeyError(f'can not bind {repr(binding)}')
            res = self.must_create_context(context, res)
            local_var = context.local_vars[bc]
            cc = do_capitalize(bc)
            context.context_vars[cc] = Object.create(cc, local_var.type)
            res.append(OpaqueExp.create(f"stubContext.{cc} = {bc}"))
            res.append(OpaqueExp.create(f"if !snippet.BindRequest(c, {bc}) {{\nreturn\n}}"))
        return None, res

    def stub_on_error(self, context: Context, lhs: List[Stmt], rhs: List[Stmt]):
        assert len(lhs) == 0
        assert len(rhs) >= 2
        if context.last_exception_index + context.exception_ttl < context.stmt_index:
            raise LookupError(
                f"exception could not handle, it is too far, "
                f"current index {context.stmt_index}, last {context.last_exception_index}, "
                f"ttl: {context.exception_ttl}")

        error_handler = f'Stub{do_capitalize(context.fn.name)}ErrorHandler{context.stmt_index}'
        if error_handler in context.extended_methods:
            raise KeyError(f"method {error_handler} is already exists in current struct scope")
        capturing = rhs[2:]

        _, func_sign_body = rhs[1].body_content.split('{', maxsplit=1)

        capturing_args = ', '.join(map(
            lambda st: f'{st.ident.name} {self.find_local_type(context, st.ident)}', capturing))
        context.extended_methods[error_handler] = OpaqueExp.create(
            f'func ({context.fn.recv.name} {context.fn.recv.type}) {error_handler}(err error, '
            f'{capturing_args}) error {{{func_sign_body}')
        capturing_input = ','.join(map(lambda st: st.ident.name, capturing))
        self.get_promise_handler(context, 'catch', context.last_exception_index).append(
            OpaqueExp.create(f'err = {context.fn.recv.name}.{error_handler}(err, {capturing_input})'))
        return None, []

    def invoking_stub_context(self, context: Context, _: List[Stmt], rhs: List[Stmt]):
        _ = self
        res = []

        for binding in rhs:
            bc = binding.body_content
            if bc not in context.local_vars:
                raise KeyError(f'can not serve {repr(binding)}')
            res = self.must_create_context(context, res)
            local_var = context.local_vars[bc]
            tbc = do_capitalize(bc)
            context.context_vars[tbc] = Object.create(tbc, local_var.type)
            res.append(OpaqueExp.create(f"stubContext.{tbc} = {bc}"))

        return 'InvokingStub', res

    def invoking_stub_do_serve(self, context: Context, _: List[Stmt], rhs: List[Stmt], key: Optional[str]):
        res = []
        res = self.must_create_err_decl(context, res)
        res = self.must_create_context(context, res)

        last_one = rhs[-1]

        service_name = context.fn.name + (key or 'Self')
        serve_handler = self.resolve_service_handler(context, service_name + 'Service', context.fn.name, key)

        if service_name not in context.service_methods:
            context.service_methods[service_name] = method_desc = (list(), list(), key)

            for binding in rhs:
                bc = binding.body_content
                tbc = do_capitalize(bc)
                if tbc in context.context_vars:
                    method_desc[0].append(context.context_vars[tbc])
                    continue

                if bc in context.local_vars:
                    local_var = context.local_vars[bc]
                    method_desc[0].append(local_var)
                    context.context_vars[tbc] = Object.create(tbc, local_var.type)
                    if binding != last_one:
                        res.append(OpaqueExp.create(f"stubContext.{tbc} = {bc}"))
                    continue
                raise KeyError(f'can not bind {repr(binding)} to service method')

            method_desc[1].append(Object.create('err', 'error'))
        else:
            raise NotImplementedError("todo check signature")
            # method_desc = context.service_methods[service_name]
            #
            # if len(method_desc[0]) != len(rhs):
            #     raise TypeError(f"signature not consistent want {method_desc[0]} got {rhs}")
            #
            # for i, binding in enumerate(rhs):
            #     bc = binding.body_content
            #     tbc = do_capitalize(bc)
            #
            #     if tbc in context.context_vars:
            #         method_desc[0][i].append(context.context_vars[tbc])
            #
            #     if bc in context.local_vars:
            #         method_desc[0][i].append(context.local_vars[bc])
            #     raise KeyError(f'can not bind {repr(binding)} to service method')
            #
            # method_desc[1].append(Object.create('err', 'error'))
        if_stmt = [
            f"if err = {serve_handler}(&stubContext, ); err != nil {{"
            f'{self.create_promise_handler(context, "catch")}\n'
            f'{self.create_promise_handler(context, "finally")}\nsnippet.DoReport(c, err)\nreturn\n}} else {{']

        for binding in rhs:
            bc = binding.body_content
            tbc = do_capitalize(bc)
            if bc in context.local_vars:
                if_stmt.append(f"{bc} = stubContext.{tbc}")

        if_stmt.append(f'{self.create_promise_handler(context, "then")}\n'
                       f'{self.create_promise_handler(context, "finally")}}}')

        res.append(OpaqueExp.create('\n'.join(if_stmt)))

        return 'Promise', res

    def invoking_stub_serve(self, context: Context, lhs: List[Stmt], rhs: List[Stmt]):
        assert len(rhs) > 0
        return self.invoking_stub_do_serve(context, lhs, rhs, None)

    def invoking_stub_serve_keyed(self, context: Context, lhs: List[Stmt], rhs: List[Stmt]):
        assert len(rhs) > 1
        # todo: key type
        return self.invoking_stub_do_serve(context, lhs, rhs[1:], rhs[0].body_content[1:-1])

    def promise_insert_stmts(self, context: Context, _: List[Stmt], rhs: List[Stmt], point_name):
        _ = self
        assert len(rhs) == 1
        res = []

        # todo rhs type
        rhs = rhs[0].body_content

        _, rhs = rhs.split('{', maxsplit=1)
        rhs = rhs.strip()[:-1]

        self.get_promise_handler(context, point_name).append(OpaqueExp.create(rhs))
        return 'Promise', res

    def promise_then(self, context: Context, lhs: List[Stmt], rhs: List[Stmt]):
        return self.promise_insert_stmts(context, lhs, rhs, 'then')

    def promise_then_ref(self, context: Context, lhs: List[Stmt], rhs: List[Stmt]):
        return self.promise_then(context, lhs, rhs)

    def promise_catch(self, context: Context, lhs: List[Stmt], rhs: List[Stmt]):
        return self.promise_insert_stmts(context, lhs, rhs, 'catch')

    def promise_catch_ref(self, context: Context, lhs: List[Stmt], rhs: List[Stmt]):
        return self.promise_catch(context, lhs, rhs)

    def promise_finally(self, context: Context, lhs: List[Stmt], rhs: List[Stmt]):
        return self.promise_insert_stmts(context, lhs, rhs, 'finally')

    def promise_finally_ref(self, context: Context, lhs: List[Stmt], rhs: List[Stmt]):
        return self.promise_finally(context, lhs, rhs)

    def must_create_context(self, context: Context, res):
        _ = self
        if not context.created_context:
            context.created_context = True
            res.append(
                OpaqueExp.create(f"var stubContext {context.fn.name}Context")
            )
        return res

    def must_create_decl(self, context: Context, decl: Object, res):
        _ = self
        if decl.name not in context.local_vars:
            context.local_vars[decl.name] = decl
            res.append(OpaqueExp.create(f"var {decl.name} {decl.type}"))
        return res

    ok_decl = Object.create('ok', 'bool')
    err_decl = Object.create('err', 'error')

    def must_create_ok_decl(self, context: Context, res):
        return self.must_create_decl(context, StubLoader.ok_decl, res)

    def must_create_err_decl(self, context: Context, res):
        return self.must_create_decl(context, StubLoader.err_decl, res)

    # noinspection PyMethodMayBeStatic
    def create_promise_handler(self, context, handler_type, index=None):
        index = index or context.stmt_index
        insert_point = f'_ = "serve_promise_{handler_type}_handler{index}"'
        context.insert_items_list[index][insert_point] = []
        return insert_point

    # noinspection PyMethodMayBeStatic
    def get_promise_handler(self, context, handler_type, index=None):
        index = index or context.stmt_index
        return context.insert_items_list[index].get(f'_ = "serve_promise_{handler_type}_handler{index}"', [])

    def global_exception_handler(self, context: Context, a: AssignExp, res):
        print(a)
        # todo
        context.last_exception_index = context.stmt_index
        if_stmt = OpaqueExp.create(
            f'{self.create_promise_handler(context, "catch")}\n'
            f"if err != nil {{\nreturn\n}}")
        return res + [a, if_stmt]

    def global_bool_handler(self, context: Context, a: AssignExp, res):
        print(a)
        # todo
        context.last_excetion_ok_index = context.stmt_index
        if_stmt = OpaqueExp.create(
            f'{self.create_promise_handler(context, "catch_ok")}\n'
            f"if !ok {{\nreturn\n}}")
        return res + [a, if_stmt]

    def find_local_type(self, context, ident):
        _ = self
        if len(ident.type) != 0:
            return ident.type
        if ident.name in context.local_vars:
            return context.local_vars[ident.name].type
