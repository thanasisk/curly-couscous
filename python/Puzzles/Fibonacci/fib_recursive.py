#!/usr/bin/env python3
import sys
import argparse

def fib(n):
    if n == 0:
        return 0
    if n < 2 and n > 0:
        return 1
    else:
        return fib(n-1) + fib(n-2)

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("-l", "--limit", default=0, type=int)
    args = parser.parse_args()
    print(fib(args.limit))

if __name__ == '__main__':
    sys.exit(main())
