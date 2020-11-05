# Build Stage
FROM golang:1.15-alpine AS build-stage
ADD ./ /build
RUN cd /build && go build -o easi-trm-gateway ./cmd/easi-trm-gateway/...

# Final Stage
FROM alpine
MAINTAINER Raphael Reyna <raphaelreyna@protonmail.com>
COPY --from=build-stage /build/easi-trm-gateway /bin/easi-trg-gateway
CMD [\"easi-trg-gateway\"]
