FROM golang:1.21-alpine AS builder

WORKDIR /build

# Copy shared jobs module
COPY jobs ./jobs

# Copy the importer
COPY importers/railway-importer ./importers/railyway-second-importer

WORKDIR /build/importers/railway-importer

# Download dependencies and build
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /railway-importer .

# Final stage
FROM alpine:3.19

# Copy the compiled binary
COPY --from=builder /railway-importer /railway-importer

CMD ["/railway-importer"]
