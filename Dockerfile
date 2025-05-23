# Build stage
FROM golang:1.21 AS builder

WORKDIR /app
COPY . .

# Disable CGO and enable static build
ENV CGO_ENABLED=0
RUN go build -o continuum ./main.go

# Final stage using a secure, minimal base
FROM gcr.io/distroless/static-debian11

COPY --from=builder /app/continuum /continuum
ENTRYPOINT ["/continuum"]