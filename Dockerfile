# syntax=docker/dockerfile:1.2

# Copyright 2020-2021 Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

ARG CILIUM_BUILDER_IMAGE=quay.io/cilium/cilium-builder:ad254d9e565fa17f3e7708d520e490a92249a5e5@sha256:a2fd024f924a606b07722a8cf56ef026a4476ad14a573c28fca9b0b5484302fd

FROM ${CILIUM_BUILDER_IMAGE} as builder
WORKDIR /go/src/github.com/cilium/cilium-cli
COPY . .
RUN make

FROM docker.io/library/busybox:stable-glibc@sha256:898fba644d1783e8de96c9a7f82255d5af2c88af53000802a1f24854aad61de7
LABEL maintainer="maintainer@cilium.io"
COPY --from=builder /go/src/github.com/cilium/cilium-cli/cilium /usr/local/bin/cilium
RUN ["wget", "-P", "/usr/local/bin", "https://dl.k8s.io/release/v1.21.0/bin/linux/amd64/kubectl"]
RUN ["chmod", "+x", "/usr/local/bin/kubectl"]
ENTRYPOINT []
