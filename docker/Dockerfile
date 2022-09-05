FROM golang:1.18
WORKDIR /src
COPY . .

RUN apt update && apt install -y --no-install-recommends unoconv
RUN go build -mod=vendor -o pdfcnv ./cmd

CMD [ "./pdfcnv" ]
