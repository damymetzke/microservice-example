# Use the official Rust image as the base image
FROM rust:1.71 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Cargo.toml and Cargo.lock files to the container
COPY Cargo.toml Cargo.lock ./

# Build the dependencies in a separate layer to take advantage of Docker layer caching
RUN cargo fetch

# Copy the source code to the container
COPY src ./src

# Build the Rust project with optimizations
RUN cargo build --release

# Create a new Docker image with a minimal runtime environment
FROM debian:buster-slim

# Set the working directory inside the container
WORKDIR /app

# Install system dependencies required by the Rust project
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Copy the built binary from the builder stage to the container
COPY --from=builder /app/target/release/user_management .

# Set the default command to run your binary
CMD ["./user_management"]
