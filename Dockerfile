FROM golang AS builder

# Download and install the latest release of dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/mediadepot/cli
COPY Gopkg.toml Gopkg.lock ./
COPY . ./
RUN rm -rf vendor
RUN dep ensure --vendor-only

RUN go build -o /mediadepot cmd/mediadepot/main.go
RUN chmod +x /mediadepot



#FROM scratch
#COPY --from=builder /mediadepot ./

#ENTRYPOINT ["./mediadepot"]
