
import os
import re


def get_go_files(p):
    for r, _, files in os.walk(p):
        for file in files:
            if file.endswith(".go"):
                yield os.path.join(r, file)


encoding = 'utf8'

def work(p):
    service_re = re.compile(r'control\.([^\sS]*)Service')
    service_re2 = re.compile(r'type ([^\sS]*)Service = ')
    for f in get_go_files(p):
        text = open(f, encoding=encoding).read()
        searched = service_re.search(text)
        if searched:
            text = service_re.sub(r'control.\1Controller', text)
            text = service_re2.sub(r'type \1Controller = ', text)

            open(f, 'w', encoding=encoding).write(text)


# def work2(p):
#     for f in get_go_files(p):
#         if f.endswith('service.go'):
#             print(f)
#             os.rename(f, f.replace('service.go', 'controller.go'))


if __name__ == '__main__':
    # work2(".")
    pass