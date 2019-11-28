FROM golang:1.10

ENV APP_NAME api
ENV PORT 8080
#ENV SRC_DIR=/go/src/github.com/project-management-api

COPY . /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/dgrijalva/jwt-go
#RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"

RUN go get ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o ${APP_NAME}
#RUN chmod +x ${APP_NAME}

CMD ./${APP_NAME}

EXPOSE ${PORT}