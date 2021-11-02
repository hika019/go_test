load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar", "pkg_deb")

go_library(
    name = "defa",
    srcs = ["main.go"],
    deps = ["//lib"],
    visibility = [":helloworld"],
)

go_binary(
    name = "helloworld",
    embed = [":defa"],
)

