import subprocess


def diff(f1, f2):
    try:
        return subprocess.call(['git', 'diff', '--no-index', f1, f2])
    except subprocess.CalledProcessError as ex:
        return ex.output
