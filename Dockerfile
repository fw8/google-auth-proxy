# https://medium.com/statuscode/golang-docker-for-development-and-production-ce3ad4e69673

FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./app /go/src/github.com/fw8/google-auth-proxy/app
WORKDIR /go/src/github.com/fw8/google-auth-proxy/app

RUN go mod init
RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
  then \
  app; \
  else \
  go get github.com/pilu/fresh && \
  fresh; \
  fi

EXPOSE 8080