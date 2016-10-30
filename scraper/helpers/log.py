import sys


def log(msg, force=False):
    if '--verbose' in sys.argv or force:
        print(msg)
