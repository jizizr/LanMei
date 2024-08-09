FROM golang:1.22-alpine
RUN apk add --no-cache libgcc gcc musl-dev
WORKDIR /app
CMD ["sh"]
