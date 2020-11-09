import pathlib
import re


def replace_meta_data(the_tmpl, key, value):
    return the_tmpl.replace(f'{{{{{key}}}}}', value)


def replace_meta_data_all(the_tmpl, metas):
    for meta in metas:
        the_tmpl = replace_meta_data(the_tmpl, meta[0], meta[1])
    return the_tmpl


def replace_insert_fields(matched, the_tmpl):
    bx = []
    last_xx = 0
    for x in re.finditer(r'([ ]*)!!insert_field:([^!]*)!!', the_tmpl):
        bx.append(the_tmpl[last_xx:x.span()[0]])
        indent, field_id = x[1], x[2]
        if field_id not in matched:
            bx.append(f'{indent}<!--beg l {field_id} -->\n{indent}\n{indent}<!--end l-->')
        else:
            bx.append(f'{indent}<!--beg l {field_id} -->{matched[field_id]}<!--end l-->')
        last_xx = x.span()[1]
    bx.append(the_tmpl[last_xx:])
    return ''.join(bx)


def prepare_path(raw_path):
    if isinstance(raw_path, str):
        raw_path = pathlib.Path(raw_path)
        raw_path.parent.mkdir(parents=True, exist_ok=True)
    return raw_path


def generate_documentation(tmpl, path, metas, in_memory_doc=None, wait_after_insert=False, match_ref=None):
    path = prepare_path(path)
    if not match_ref:
        match_ref = dict()
    matched = match_ref
    if in_memory_doc is not None:
        path_x = in_memory_doc
    elif path.exists():
        path_x = path.read_text(encoding='utf8')
    else:
        path_x = ''

    for x in re.finditer(r'<!--beg l (\S*) -->([\s\S]*)<!--end l-->', path_x):
        matched[x[2]] = x[1]

    # step2: replace insert field
    tmpl = replace_insert_fields(
        matched,
        # step1: replace meta data
        replace_meta_data_all(tmpl, metas))

    if not wait_after_insert:
        # step3: write text
        path.write_text(tmpl, encoding='utf8')

    return tmpl


import bisect


def get_doc(tmpl, finding):
    finding_l = [0] * len(finding)
    for i, f in enumerate(finding):
        finding_l[i] = tmpl.find(f)

    search_list = list(sorted(finding_l + [len(tmpl)]))

    for i, f in enumerate(finding_l):
        finding_r = search_list[bisect.bisect_right(search_list, f)]
        finding[i] = tmpl[finding_l[i] + len(finding[i]): finding_r]
    return finding
