FROM nginx:latest

COPY ./dhparam.pem /etc/nginx
COPY ./fido.workshop+4-key.pem /etc/nginx/certs/fido.workshop+4-key.pem
COPY ./fido.workshop+4.pem /etc/nginx/certs/fido.workshop+4.pem
COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./sites-available/default /etc/nginx/sites-available/default
COPY ./tls_config.conf /etc/nginx/snippets/tls_config.conf
