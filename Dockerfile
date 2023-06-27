FROM golang:1.20
COPY . /ascii-art-web
WORKDIR /ascii-art-web/server
RUN go build -v main.go
CMD ./main
EXPOSE 8080
LABEL "version"="1.0"
LABEL "project name"="ascii-art-web"
LABEL "description"="hosting ascii-art project on a web server that uses a simple GUI"