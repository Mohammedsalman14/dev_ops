FROM ubuntu:latest

WORKDIR /app

RUN apt update && \
    apt install -y nodejs npm curl

COPY . .


RUN npm install 

EXPOSE 3000

CMD ["node", "app.js"]

