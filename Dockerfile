# NOTE! the dockerfile is not working right now!
# Build frontend dist.
FROM node:18.12.1-alpine3.16 AS frontend
WORKDIR /frontend-build

COPY ./web/ .
RUN corepack enable && pnpm install
RUN pnpm build

# Build backend exec file.
FROM golang:1.19.3-alpine3.16 AS backend
WORKDIR /backend-build

COPY . .
COPY --from=frontend /frontend-build/dist ./backend/server/dist

RUN CGO_ENABLED=0 go build -o tools ./main.go


# Make workspace with above generated files.
FROM alpine:3.16 AS monolithic
WORKDIR /usr/local/tools

RUN apk add --no-cache tzdata
ENV TZ="UTC"
ENV PROFILE="PRODUCTION"
COPY --from=backend /backend-build/tools /usr/local/tools/
COPY --from=backend /backend-build/script /usr/local/tools/script
EXPOSE 8080
RUN apk add python3  py3-pip
RUN pip install requests yt_dlp ytdl-nfo
RUN mkdir -p /var/opt/tools
RUN mkdir -p /var/opt/video
RUN mkdir -p /var/opt/video/temp

VOLUME /var/opt/tools

ENTRYPOINT ["./tools"]