FROM golang:1.20-bullseye as builder

# Install minimum necessary dependencies
ENV PACKAGES curl make git libc-dev bash gcc
RUN apt-get update \
    && apt-get upgrade -y \
    && apt-get install -y $PACKAGES

WORKDIR /root

# Download and compile the ottod binaries
RUN git clone --depth 1 -b feat/otto-upgrade-ccv https://github.com/octopus-appchains/otto.git otto-v1 \
    && cd otto-v1 \
    && make build

# Install the cosmovisor binary
ENV COSMOS_VERSION=v0.47.5
RUN git clone --depth 1 -b ${COSMOS_VERSION} https://github.com/cosmos/cosmos-sdk.git \
    && cd cosmos-sdk/tools/cosmovisor \
    && make cosmovisor

# Final image
FROM ubuntu:22.04

RUN apt-get update \
    && apt-get upgrade -y \
    && apt-get install -y curl jq

WORKDIR /root

# Set the environmental variables
ENV DAEMON_HOME=/data
ENV DAEMON_NAME=ottod

# Create the folder structure required for using cosmovisor
RUN mkdir -p cosmovisor/genesis/bin

# Copy the cosmovisor binary
COPY --from=builder /root/cosmos-sdk/tools/cosmovisor/cosmovisor /usr/local/bin/cosmovisor

# Copy the ottod binaries to the appropriate directories
COPY --from=builder /root/otto-v1/build/$DAEMON_NAME /usr/local/bin/$DAEMON_NAME
COPY --from=builder /root/otto-v1/build/$DAEMON_NAME /root/cosmovisor/genesis/bin/$DAEMON_NAME

# Copy the script to the appropriate directories
COPY --from=builder /root/otto-v1/scripts/init_and_start.sh /root/init_and_start.sh
RUN chmod +x /root/init_and_start.sh

EXPOSE 26656 26657 1317 9090 8545 8546
USER 0

ENTRYPOINT ["/bin/bash", "-c"]