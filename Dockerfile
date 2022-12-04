FROM golang:1.19

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./
COPY controllers/ controllers/
COPY static/ static/

RUN go build -o /docker-aoc2022web

EXPOSE 80

CMD [ "/docker-aoc2022web" ]