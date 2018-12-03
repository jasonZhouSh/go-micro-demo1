# 暂未将 Golang 集成到 docker 中
FROM alpine:latest
RUN mkdir /app
WORKDIR /app
ADD user-service /app/user-service
ADD conf /app/conf 
CMD ["./user-service"]
