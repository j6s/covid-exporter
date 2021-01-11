FROM golang:1.15

COPY ./ /build
RUN cd /build \
    && go build -o /covid-exporter /build/main.go \
    && rm -Rf /build

ENTRYPOINT /covid-exporter
