FROM node:18
WORKDIR /app
COPY . /app
RUN npm install 
EXPOSE 300
COPY wait-for-mysql.sh .
CMD ["sh", "wait-for-mysql.sh"]
