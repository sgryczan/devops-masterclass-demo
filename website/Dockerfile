FROM golang:1.12 as builder

WORKDIR /go/src/gitlab.com/sgryczan/website
COPY . .

ENV PORT 8080

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN CGO_ENABLED=0 GOOS=linux go build -v -o website

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/gitlab.com/sgryczan/website/website .
COPY --from=builder /go/src/gitlab.com/sgryczan/website/models ./models
COPY --from=builder /go/src/gitlab.com/sgryczan/website/static ./static
COPY --from=builder /go/src/gitlab.com/sgryczan/website/templates ./templates

# Run the web service on container startup.
CMD ["./website"]