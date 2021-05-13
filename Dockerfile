FROM golang:1.15.11-alpine

WORKDIR /home/ica

COPY go.mod . 
COPY go.sum .

RUN go mod download

COPY . .

RUN for bin in cmd/*; do \ 
        go install ./cmd/$(basename $bin); \ 
    done

RUN addgroup ica && \
    adduser -S -G ica ica -h /home/ica

USER ica:ica

ENTRYPOINT [ "./scripts/docker-entrypoint.sh" ]
