FROM golang:1.14.2 AS build-env

WORKDIR /go/src/app

ADD . /dockerbuild
WORKDIR /dockerbuild

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -i -o ./file_stat ./main.go
RUN chmod +x file_stat

FROM ubuntu:20.04

RUN useradd -ms /bin/bash appuser

USER appuser
WORKDIR /home/appuser

# Creating some files
RUN echo "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer ac." > lorem_ipsum-1.txt
RUN echo "Lorem ipsum dolor sit amet, consectetur adipiscing elit." > lorem_ipsum-2.txt
RUN echo "Lorem ipsum dolor sit amet." > lorem_ipsum-3.txt

RUN echo "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer ac." > /tmp/lorem_ipsum-1.txt
RUN echo "Lorem ipsum dolor sit amet, consectetur adipiscing elit." > /tmp/lorem_ipsum-2.txt
RUN echo "Lorem ipsum dolor sit amet." > /tmp/lorem_ipsum-3.txt

EXPOSE 5150

COPY --from=build-env /dockerbuild/file_stat /usr/local/bin/

CMD ["/usr/local/bin/file_stat", "serve"]