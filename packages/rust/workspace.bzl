load("@rules_rust//crate_universe:defs.bzl", "crate_universe")
load(":crate_universe_default.bzl", "DEFAULT_URL_TEMPLATE", "DEFAULT_SHA256_CHECKSUMS")

def deps():
    crate_universe(
        name = "crypto_math_deps",
        cargo_toml_files = ["//packages/rust:Cargo.toml"],
        lockfile = "//packages/rust:bazel.lock",
        resolver_download_url_template = DEFAULT_URL_TEMPLATE,
        resolver_sha256s = DEFAULT_SHA256_CHECKSUMS,
    )
