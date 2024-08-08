FROM alpine:latest
RUN apk add --no-cache libgcc
WORKDIR /app
COPY . .
CMD ["sh","bootstrap.sh"]
