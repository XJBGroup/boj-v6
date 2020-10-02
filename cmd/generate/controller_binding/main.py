import subprocess

if __name__ == '__main__':
    process = subprocess.Popen(
        ' '.join(['go run github.com/Myriad-Dreamin/boj-v6/cmd/generate/controller_binding/ast-dump',
                  'github.com/Myriad-Dreamin/boj-v6/cmd/generate/controller_binding/inner/model',
                  '.cache/ast_dump']))
    process.wait()

    o, e = process.communicate()

    print(process.returncode, o, e)
