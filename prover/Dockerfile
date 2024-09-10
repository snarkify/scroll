# The build step is skipped, one should build on the host first
# Use the NVIDIA CUDA runtime image for the runtime stage
FROM nvidia/cuda:12.6.0-base-ubuntu22.04

WORKDIR /prover

# Copy the build artifact from the build stage
COPY target/release/prover prover
# Set environment variables
ENV LD_LIBRARY_PATH=/prover:/usr/local/cuda/lib64 \
    SCROLL_PROVER_ASSETS_DIR=/prover/volume/assets_low,/prover/volume/assets_high \
    RUST_MIN_STACK=100000000 \
    CHAIN_ID=534352
# Set the entrypoint for the container
ENTRYPOINT ["./prover"]