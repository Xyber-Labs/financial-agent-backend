FROM golang:1.23.3-bookworm AS builder

RUN apt update && apt install -y ca-certificates

WORKDIR /app
COPY . .

ENV APP_NAME=financial-agent
ENV APP_BIN=/bin/${APP_NAME}

RUN go mod download

RUN BUILD_COMMIT=$(git rev-parse HEAD) CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -ldflags '-s -w' -o ${APP_BIN} -v ./cmd

FROM golang:1.23.3-bookworm AS runtime

RUN apt update && apt install -y ca-certificates

WORKDIR /app

ENV APP_NAME=financial-agent
ENV APP_BIN=/bin/${APP_NAME}

COPY --from=builder ${APP_BIN} ${APP_BIN}

CMD [ "/bin/financial-agent" ]