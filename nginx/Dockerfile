FROM nginx:1.15.8

RUN rm /etc/nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/
COPY webproject.conf /etc/nginx/conf.d/
