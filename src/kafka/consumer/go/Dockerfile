FROM golang:latest
WORKDIR /go/src/app
RUN go get github.com/Shopify/sarama
RUN go get gopkg.in/alecthomas/kingpin.v2
COPY main.go ./
RUN go build -o main .
CMD ["/go/src/app/main","--brokerList=kafka-kafka:9092", "--topic=test"]

