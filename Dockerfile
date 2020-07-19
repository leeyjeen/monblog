FROM golang:1.12.1-stretch as packages
LABEL maintainer="leeyjeen@naver.com"

# set go module mode without GOPATH
ENV GO111MODULE=on

WORKDIR /usr/src/app
COPY . .
RUN make build

FROM golang:1.12.1-stretch as stage

RUN apt-get update && \
    apt-get -y install netcat && \
    apt-get -y install jq && \
    apt-get -y install uuid-runtime && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app

COPY --from=packages /usr/src/app/scripts/entrypoint.sh /usr/src/app/scripts/entrypoint.sh
COPY --from=packages /usr/src/app/monblog_api /usr/src/app/monblog_api

RUN chmod +x ./scripts/entrypoint.sh

CMD ["sh", "./scripts/entrypoint.sh"]