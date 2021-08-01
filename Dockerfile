FROM golang:1.16 As base
RUN mkdir /app_dir
COPY . /app_dir
WORKDIR /app_dir
RUN go build -o /app .

FROM golang:1.16 As final
COPY --from=base /app /app
CMD [ "/app "]
