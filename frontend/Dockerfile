FROM node:20-alpine3.17 AS builder
RUN mkdir /app
COPY . /app 
WORKDIR /app
RUN npm install && npm run build


FROM nginx:latest
COPY --from=builder /app/dist/frontend /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
ENTRYPOINT ["nginx", "-g", "daemon off;"]
