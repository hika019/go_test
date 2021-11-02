load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "lib",
    srcs = ["greet.go", "hello.go"],
    visibility = ["//main:helloworld"],
    importpath = "//main"
)