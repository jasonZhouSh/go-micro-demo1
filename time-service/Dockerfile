# 暂未将 Golang 集成到 docker 中
FROM alpine:latest
RUN mkdir /app
WORKDIR /app
ADD time-service /app/time-service
ADD conf /app/conf 
CMD ["./time-service"]
