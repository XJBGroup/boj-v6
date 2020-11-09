from typing import *


class TestCaseScriptStatement(object):
    func_name: Union[None, 'str']
    args: Union[None, 'List[object]']

    def __init__(self):
        self.func_name = None
        self.args = None

    @staticmethod
    def from_dict(d: dict) -> 'TestCaseScriptStatement':
        ci = TestCaseScriptStatement()
        if d is None:
            return ci
        ci.func_name = getting = d.get('func_name', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.func_name = str(getting)
        ci.args = getting = d.get('args', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, list):
                g_a_cpc_x_375c7701 = list()
                for g_a_cpc_x_375c7701_value in getting:
                    g_a_cpc_x = g_a_cpc_x_375c7701_value
                    g_a_cpc_x_375c7701.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_375c7701
            ci.args = g_a_cpc_x
        return ci


class TestCase(object):
    abstract: Union[None, 'bool']
    path: Union[None, 'str']
    name: Union[None, 'str']
    meta: Union[None, 'Dict[str, object]']
    scripts: Union[None, 'List[TestCaseScriptStatement]']

    def __init__(self):
        self.abstract = None
        self.path = None
        self.name = None
        self.meta = None
        self.scripts = None

    @staticmethod
    def from_dict(d: dict) -> 'TestCase':
        ci = TestCase()
        if d is None:
            return ci
        ci.abstract = getting = d.get('abstract', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.abstract = bool(getting)
        ci.path = getting = d.get('path', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.path = str(getting)
        ci.name = getting = d.get('name', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.name = str(getting)
        ci.meta = getting = d.get('meta', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, dict):
                g_a_cpc_x_5585a183 = dict()
                for g_a_cpc_x_5585a183_value in getting.items():
                    g_a_cpc_x = g_a_cpc_x_5585a183_value[1]
                    g_a_cpc_x_5585a183[g_a_cpc_x_5585a183_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_5585a183
            ci.meta = g_a_cpc_x
        ci.scripts = getting = d.get('scripts', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, list):
                g_a_cpc_x_2d8c579c = list()
                for g_a_cpc_x_2d8c579c_value in getting:
                    g_a_cpc_x = TestCaseScriptStatement.from_dict(g_a_cpc_x_2d8c579c_value)
                    g_a_cpc_x_2d8c579c.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_2d8c579c
            ci.scripts = g_a_cpc_x
        return ci


class DataRecord(object):
    comment: Union[None, 'str']
    request_body: Union[None, 'str']
    request_header: Union[None, 'Dict[str, List[str]]']
    response_code: Union[None, 'int']
    response_body: Union[None, 'str']
    response_header: Union[None, 'Dict[str, List[str]]']

    def __init__(self):
        self.comment = None
        self.request_body = None
        self.request_header = None
        self.response_code = None
        self.response_body = None
        self.response_header = None

    @staticmethod
    def from_dict(d: dict) -> 'DataRecord':
        ci = DataRecord()
        if d is None:
            return ci
        ci.comment = getting = d.get('comment', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.comment = str(getting)
        ci.request_body = getting = d.get('request_body', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.request_body = str(getting)
        ci.request_header = getting = d.get('request_header', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, dict):
                g_a_cpc_x_1f7a1009 = dict()
                for g_a_cpc_x_1f7a1009_value in getting.items():
                    if isinstance(g_a_cpc_x_1f7a1009_value[1], list):
                        g_a_cpc_x_5d8dc527 = list()
                        for g_a_cpc_x_5d8dc527_value in g_a_cpc_x_1f7a1009_value[1]:
                            g_a_cpc_x = str(g_a_cpc_x_5d8dc527_value)
                            g_a_cpc_x_5d8dc527.append(g_a_cpc_x)
                        g_a_cpc_x = g_a_cpc_x_5d8dc527
                    g_a_cpc_x_1f7a1009[g_a_cpc_x_1f7a1009_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_1f7a1009
            ci.request_header = g_a_cpc_x
        ci.response_code = getting = d.get('response_code', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.response_code = int(getting)
        ci.response_body = getting = d.get('response_body', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.response_body = str(getting)
        ci.response_header = getting = d.get('response_header', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, dict):
                g_a_cpc_x_458c4e2c = dict()
                for g_a_cpc_x_458c4e2c_value in getting.items():
                    if isinstance(g_a_cpc_x_458c4e2c_value[1], list):
                        g_a_cpc_x_3660457d = list()
                        for g_a_cpc_x_3660457d_value in g_a_cpc_x_458c4e2c_value[1]:
                            g_a_cpc_x = str(g_a_cpc_x_3660457d_value)
                            g_a_cpc_x_3660457d.append(g_a_cpc_x)
                        g_a_cpc_x = g_a_cpc_x_3660457d
                    g_a_cpc_x_458c4e2c[g_a_cpc_x_458c4e2c_value[0]] = (g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_458c4e2c
            ci.response_header = g_a_cpc_x
        return ci


class DataResult(object):
    handler: Union[None, 'str']
    method: Union[None, 'str']
    path: Union[None, 'str']
    data_records: Union[None, 'List[DataRecord]']

    def __init__(self):
        self.handler = None
        self.method = None
        self.path = None
        self.data_records = None

    @staticmethod
    def from_dict(d: dict) -> 'DataResult':
        ci = DataResult()
        if d is None:
            return ci
        ci.handler = getting = d.get('handler', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.handler = str(getting)
        ci.method = getting = d.get('method', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.method = str(getting)
        ci.path = getting = d.get('path', None)
        if getting is not None:
            g_a_cpc_x = None
            ci.path = str(getting)
        ci.data_records = getting = d.get('data_records', None)
        if getting is not None:
            g_a_cpc_x = None
            if isinstance(getting, list):
                g_a_cpc_x_130b1d98 = list()
                for g_a_cpc_x_130b1d98_value in getting:
                    g_a_cpc_x = DataRecord.from_dict(g_a_cpc_x_130b1d98_value)
                    g_a_cpc_x_130b1d98.append(g_a_cpc_x)
                g_a_cpc_x = g_a_cpc_x_130b1d98
            ci.data_records = g_a_cpc_x
        return ci
