# stage1 ubuntu base image
FROM ubuntu as build

# updating and install golang runtime
RUN apt update && \
    apt install -y golang-go

# app is the work directory
WORKDIR /app

# copying the source code from local pc
COPY . .

# build the golang
RUN go build -o calc app.go

# --------------------------------------------------

# stage2 use the golang runtime 
FROM scratch
# app is the work directory
WORKDIR /app

# copy from app dir from as build
COPY --from=build /app /app


ENTRYPOINT [ "./calc" ]

CMD ["10", "+", "5"]