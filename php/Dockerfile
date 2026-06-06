FROM php:8.3-fpm
WORKDIR /var/www/html
COPY . .
RUN apt-get install -y libzip-dev && docker-php-ext-install zip
EXPOSE 8080
CMD ["php", "-S", "0.0.0.0:8080", "-t", "/var/www/html"]
