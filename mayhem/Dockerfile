# FROM fuzzers/go-fuzz:1.2.0
FROM golang:1.18 AS builder

RUN apt-get update && apt-get install -y clang && rm -rf /var/lib/apt/lists/*

# Install go-fuzz-build pinned to a Go-1.18-compatible commit, with x/tools pinned
# to v0.1.12 to avoid the x/tools@v0.44.0 requirement for Go 1.22+ stdlib packages.
RUN mkdir /tmp/igf && cd /tmp/igf && go mod init igf && \
    printf "module igf\n\ngo 1.18\n\nrequire github.com/dvyukov/go-fuzz v0.0.0-20220726122315-1d375ef9f9f6\n\nrequire golang.org/x/tools v0.1.12\n" > go.mod && \
    GOFLAGS=-mod=mod go get golang.org/x/tools@v0.1.12 && \
    GOFLAGS=-mod=mod go install github.com/dvyukov/go-fuzz/go-fuzz-build && \
    rm -rf /tmp/igf

COPY . /calcite-avatica-go
WORKDIR /calcite-avatica-go

RUN GOFLAGS=-mod=mod go get github.com/dvyukov/go-fuzz@v0.0.0-20220726122315-1d375ef9f9f6

WORKDIR /calcite-avatica-go/mayhem
RUN go-fuzz-build -libfuzzer -o fuzzOpen.a . && \
    clang -fsanitize=fuzzer fuzzOpen.a -o fuzzOpen.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /calcite-avatica-go/mayhem/fuzzOpen.libfuzzer /
