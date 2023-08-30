FROM golang
WORKDIR /source
COPY ./source ./
RUN cd /source
RUN go mod download
RUN go build -o /main main.go 
CMD ["/main"]
