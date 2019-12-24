FROM golang:1.13

ENV APP_NAME api
ENV PORT 8080

COPY . /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN go build -o ${APP_NAME}


CMD ./${APP_NAME}

EXPOSE 8080