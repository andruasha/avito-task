FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y nginx

COPY default.conf /etc/nginx/conf.d/default.conf

EXPOSE 8089

CMD ["nginx", "-g", "daemon off;"]
