load("@crypto_math_deps//:defs.bzl", "crates_from")

_commons_deps = crates_from("//packages/rust:Cargo.toml")

def lib_deps():
    return _commons_deps

def bin_deps():
    _deps = [":crypto_math"]
    return _commons_deps + _deps
