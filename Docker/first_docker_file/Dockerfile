FROM ubuntu:latest

WORKDIR /app

COPY . /app/

RUN  apt update && apt install python3 -y 

ENV NAME world

CMD [ "python3", "app.py" ]
