import sys


def log(msg):
    if '--verbose' in sys.argv:
        print(msg)
