FROM registry.fedoraproject.org/fedora:latest

RUN dnf --setopt=install_weak_deps=False install -y golang-bin

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /main

EXPOSE 8080

CMD [ "/main" ]
