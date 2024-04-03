#! /usr/bin/python3
# coding=utf-8

class Program:
    def __init__(self):
        try:
            import src.core.langa
        except KeyboardInterrupt:
            sys.exit(0)

if __name__ == '__main__':
    main = Program()
