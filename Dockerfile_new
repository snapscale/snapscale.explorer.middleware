FROM ubuntu:18.04

RUN apt-get update -y && \
	apt-get install -y wget && \
	wget https://dl.google.com/go/go1.13.5.linux-amd64.tar.gz && \
	tar zxvf go1.13.5.linux-amd64.tar.gz -C /usr/local

ENV PATH="/usr/local/go/bin:${PATH}"

COPY . /app
WORKDIR /app

RUN go build -o /backend

FROM ubuntu:18.04

COPY --from=0 /backend /opt/xeniro/bin/explorermw

EXPOSE 8089
EXPOSE 8090

ENTRYPOINT  ["/opt/xeniro/bin/explorermw"]

