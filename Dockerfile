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
COPY --from=backend /backend-build/main.py /usr/local/tools/main.py
COPY --from=backend /backend-build/config.yaml /usr/local/tools/config.yaml

EXPOSE 7789
RUN apk add python3 py3-pip ffmpeg
RUN pip install requests yt_dlp ytdl-nfo
RUN mkdir -p /var/opt/tools
RUN mkdir -p /var/opt/video
COPY --from=backend /backend-build/bilibili.yaml /usr/lib/python3.10/site-packages/ytdl_nfo/configs/bilibili.yaml
VOLUME /var/opt/tools

ENTRYPOINT ["./tools"]