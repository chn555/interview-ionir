FROM debian:bullseye-slim

RUN apt-get update

RUN apt-get install ca-certificates -y

ADD main /usr/local/bin

RUN chmod 755 /usr/local/bin/main

RUN mkdir -p /etc/main

ENTRYPOINT ["/bin/bash", "-c"]

CMD ["/usr/local/bin/main"]