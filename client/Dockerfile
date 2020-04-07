# 開発環境
#FROM node:10.12-alpine as build-stage
#FROM ubuntu:latest
FROM node:13.7.0 as build-stage
WORKDIR /app
COPY . .

#RUN echo "file_uploads = On\n" \
#         "memory_limit = 500M\n" \
#         "upload_max_filesize = 500M\n" \
#         "post_max_size = 500M\n" \
#         "max_execution_time = 600\n" \
#         > /usr/local/etc/php/conf.d/uploads.ini

RUN yarn install
#RUN npm install sass-loader node-sass flush-promises --save
#RUN npm install vuex-persistedstate  marked  --save
#RUN npm rebuild node-sass
#RUN yarn add element-ui
#RUN vue add element


