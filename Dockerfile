FROM golang:1.18
WORKDIR /src
COPY . .

RUN apt update && apt install -y unoconv
RUN go build -mod=vendor -o pdfcnv ./cmd/main.go

CMD [ "./pdfcnv" ]
