# Bigtable bttest in docker
# golang version: 1.5.1
FROM golang:1.5.1-wheezy
MAINTAINER hsinho yeh<yhh92u@gmail.com>

ENV BTTESTROOT $GOPATH/src/bttest

RUN go get -u google.golang.org/grpc
RUN go get -u google.golang.org/cloud

# (hack) run commid id: 3af561783031fd92deac27de022a52dcd81dab34 to avoid build break
# TODO: remove me if the master is fixed
RUN cd $GOPATH/src/google.golang.org/grpc && git checkout 3af561783031fd92deac27de022a52dcd81dab34
RUN go get -u github.com/bradfitz/http2

#(hack) replace 127.0.0.1:0 to :9001, so we can listen on 9001
RUN sed -i 's/127.0.0.1:0/:9001/g' $GOPATH/src/google.golang.org/cloud/bigtable/bttest/inmem.go

ADD ./bttest $BTTESTROOT
RUN chmod +x $BTTESTROOT/runlocal.sh
RUN cd $BTTESTROOT && ./build.sh

# 9001 for bigtable test server port
EXPOSE 9001

WORKDIR $BTTESTROOT
CMD "./runlocal.sh"
