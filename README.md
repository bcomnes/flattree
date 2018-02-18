# flattree
[![GoDoc][godoc-img]][godoc]

A series of functions to map a binary tree to a list. A port of [flat-tree][ft] to go. 

## Install

```
dep ensure -add github.com/bcomnes/flattree
```

## Usage

You can represent a binary tree in a simple flat list using the following structure

```
      3
  1       5
0   2   4   6  ...
```

See [Godoc][example] example on godoc.

## API

See [API][api] example on godoc.

## See also

- [mafintosh/flat-tree][ft]: The node module that this was ported from.
- [mafintosh/flat-tree-rs][rs]: A port of the node module to rust.
- [mafintosh/print-flat-tree][print]: A node cli that can pretty print flat-trees.

[ft]: https://github.com/mafintosh/flat-tree
[godoc]: https://godoc.org/github.com/bcomnes/flattree
[godoc-img]: https://godoc.org/github.com/bcomnes/flattree?status.svg
[example]: https://godoc.org/github.com/bcomnes/flattree
[api]: https://godoc.org/github.com/bcomnes/flattree#pkg-index
[print]: https://github.com/mafintosh/print-flat-tree
[rs]: https://github.com/mafintosh/flat-tree-rs

