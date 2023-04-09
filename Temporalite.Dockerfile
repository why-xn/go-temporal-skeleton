FROM golang:1.19

RUN apt-get -y update && apt-get install -y wget zip unzip

WORKDIR /app

RUN wget https://github.com/temporalio/temporalite/releases/download/v0.3.0/temporalite_0.3.0_linux_amd64.tar.gz -O /app/temporalite.tar.gz

RUN tar -xf /app/temporalite.tar.gz

RUN mkdir /app/db

CMD /app/temporalite start --namespace default --ip 0.0.0.0 --ui-ip 0.0.0.0 --ui-port 8233 -f /app/db/temporal.db