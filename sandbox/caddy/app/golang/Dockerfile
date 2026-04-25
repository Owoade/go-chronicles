# -------------------------
# 1. Build stage
# -------------------------
    FROM golang:1.22-alpine AS builder

    WORKDIR /app
    
    # Copy source code
    COPY . .
    
    # Build static binary
    RUN go build -o server main.go
    
    
    # -------------------------
    # 2. Runtime stage
    # -------------------------
    FROM alpine:latest
    
    WORKDIR /app
    
    # Copy binary from builder
    COPY --from=builder /app/server .
    
    # Expose port
    EXPOSE 3000
    
    # Run the binary
    CMD ["./server"]