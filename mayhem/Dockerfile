# FROM fuzzers/go-fuzz:1.2.0
FROM golang:1.18 as builder

COPY . /calcite-avatica-go
WORKDIR /calcite-avatica-go/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzOpen.a && \
    clang -fsanitize=fuzzer fuzzOpen.a -o fuzzOpen.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /calcite-avatica-go/mayhem/fuzzOpen.libfuzzer /