load("@crypto-math-deps//:defs.bzl", "crate")

_common_deps = []

def libs_deps():
    return _common_deps

def bin_deps():
    _deps = [":crypto-math"]
    return _common_deps
