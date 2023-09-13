FROM golang:1.19.4-alpine3.17

EXPOSE 9090

RUN apk update \
    && apk add --no-cache \ 
    mysql-client \
    build-base

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
COPY ./entrypoint.sh /usr/local/bin/entrypoint.sh
RUN /bin/chmod +x /usr/local/bin/entrypoint.sh

RUN go build -o bin/evs
RUN mv bin/evs /usr/local/bin/

CMD ["evs"]
ENTRYPOINT ["entrypoint.sh"]