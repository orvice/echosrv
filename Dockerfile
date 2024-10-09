FROM golang:1.23 as builder

ARG ARG_GOPROXY
ENV GOPROXY $ARG_GOPROXY

WORKDIR /home/app
COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN make build


FROM orvice/go-runtime

ENV PROJECT_NAME echosrv

COPY --from=builder /home/app/bin/${PROJECT_NAME} /home/app/bin/${PROJECT_NAME}
ENTRYPOINT /home/app/bin/${PROJECT_NAME}
