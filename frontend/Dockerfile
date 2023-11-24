FROM --platform=linux/amd64 node:18 as builder

WORKDIR /app

COPY . .

RUN npm install

RUN npm run build

FROM --platform=linux/amd64 thistine/simple-http-server as runner

WORKDIR /app

COPY --from=builder /app/dist ./static


ENTRYPOINT ["/app/serviceapifrontend"]

EXPOSE 4000