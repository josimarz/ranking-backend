FROM golang:alpine@sha256:7772cb5322baa875edd74705556d08f0eeca7b9c4b5367754ce3f2f00041ccee AS builder

WORKDIR /var/app

COPY . .

RUN go build -ldflags='-s' -o api ./cmd/api

FROM scratch

WORKDIR /var/app

COPY --from=builder /var/app/api .

ARG PORT
ARG AWS_TABLE
ARG AWS_BUCKET
ARG AWS_ENDPOINT_URL
ENV PORT=${PORT}
ENV AWS_TABLE=${AWS_TABLE}
ENV AWS_BUCKET=${AWS_BUCKET}
ENV AWS_ENDPOINT_URL=${AWS_ENDPOINT_URL}

EXPOSE ${PORT}

ENTRYPOINT [ "./api" ]