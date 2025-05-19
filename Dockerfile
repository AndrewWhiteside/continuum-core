FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go build -o continuum ./main.go

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/continuum /continuum
ENTRYPOINT ["/continuum"]