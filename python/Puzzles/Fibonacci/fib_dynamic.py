#!/usr/bin/env python3
import sys
import argparse

FibArray = [0, 1]

def fib(n):
    if n < 0:
        print("nope! not possible!")
    elif n == 0:
        return FibArray[0]
    elif n == 1:
        return FibArray[1]
    elif n <= len(FibArray):
        return FibArray[n-1]
    else:
        tempFib = fib(n-1) + fib(n-2)
        FibArray.append(tempFib)
        return tempFib

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("-l", "--limit", default=0, type=int)
    args = parser.parse_args()
    print(fib(args.limit))

if __name__ == '__main__':
    sys.exit(main())
