# syntax=docker/dockerfile:1
FROM ubuntu:22.04
WORKDIR /usr/local/app
SHELL ["/bin/bash", "-c"]

# install app dependencies
RUN apt-get update && apt-get install -y wget make && wget -c https://go.dev/dl/go1.22.5.linux-amd64.tar.gz && tar -C /usr/local/ -xzf go1.22.5.linux-amd64.tar.gz 
ENV PATH="$PATH:/usr/local/go/bin"

# install app
COPY  ./ ./
RUN go get && make build

# final configuration
EXPOSE 8080
RUN useradd app
USER app
CMD ["./bin/home_mgmt"]




