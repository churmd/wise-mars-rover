# syntax=docker/dockerfile:1

FROM golang:1.17 as build
 
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o mars-rover . 


From scratch
COPY --from=build /app/mars-rover .
CMD [ "/mars-rover" ]
