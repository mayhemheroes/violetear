FROM golang:1.18 as builder

COPY . /violetear
WORKDIR /violetear/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzVioletear.a && \
    clang -fsanitize=fuzzer fuzzVioletear.a -o fuzzVioletear.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /violetear/mayhem/fuzzVioletear.libfuzzer /