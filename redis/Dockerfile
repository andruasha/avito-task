FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y redis-server

COPY redis.conf /etc/redis/redis.conf

CMD ["redis-server", "/etc/redis/redis.conf"]
