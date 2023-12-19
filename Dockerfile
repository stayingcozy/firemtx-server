FROM bluenviron/mediamtx:latest

WORKDIR /app

COPY /brokenpipe/brokenpipe /brokenpipe

COPY firemtx-server .

COPY *.json .

CMD ["./firemtx-server"]

