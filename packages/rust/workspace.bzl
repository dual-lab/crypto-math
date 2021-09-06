load("@rules_rust//crate_universe:defs.bzl", "crate", "crate_universe")

def deps():
  crate_universe(
      name = "crypto-math-deps",
      cargo_toml_files = ["//packages/rust:Cargo.toml"],
      lockfile = "//packages/rust:bazel.lock",
      packages = [],
      resolver = "@rules_rust_crate_universe_bootstrap//:crate_universe_resolver",
      )
