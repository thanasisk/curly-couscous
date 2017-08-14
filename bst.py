#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import argparse
import random
import sys

# shamelessly stolen from SO :-P
class TreeNode:

    def __init__(self, value):
        self.left = None
        self.right = None
        self.data = value

class Tree:

    def __init__(self):
        self.root = None

    def addNode(self, node, value):
        if node is None:
            self.root = TreeNode(value)
        else:
            if (value < node.data):
                if node.left is None:
                    node.left = TreeNode(value)
                else:
                    self.addNode(node.left, value)
            else:
                if node.right is None:
                    node.right = TreeNode(value)
                else:
                    self.addNode(node.right, value)

    def printInOrder(self, node):
        if node is not None:
            self.printInOrder(node.left)
            print(node.data)
            self.printInOrder(node.right)

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("nodes", type=int,
                    help="how many nodes to generate", default=100)
    args = parser.parse_args()
    bst = Tree()
    for i in range(0, args.nodes):
        bst.addNode(bst.root, random.randint(0,2048))
    bst.printInOrder(bst.root)

if __name__ == '__main__':
    sys.exit(main())
