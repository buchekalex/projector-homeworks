FROM debian:buster

ENV NGINX_VERSION 1.17.6
ENV NGX_CACHE_PURGE_VERSION=2.5

# Install basic packages and build tools
RUN apt-get update && apt-get install -y \
    wget \
    build-essential \
    libssl-dev \
    libpcre3 \
    libpcre3-dev \
    zlib1g-dev \
    && apt-get clean

# Get sources
RUN cd /tmp && \
    wget "http://nginx.org/download/nginx-${NGINX_VERSION}.tar.gz" && \
    wget "https://github.com/nginx-modules/ngx_cache_purge/archive/${NGX_CACHE_PURGE_VERSION}.tar.gz" -O ngx_cache_purge-$NGX_CACHE_PURGE_VERSION.tar.gz && \
    tar -zxvf nginx-$NGINX_VERSION.tar.gz && mv nginx-$NGINX_VERSION nginx && rm nginx-$NGINX_VERSION.tar.gz && ls -lhtr /tmp && \
    tar -zxvf ngx_cache_purge-$NGX_CACHE_PURGE_VERSION.tar.gz && mv ngx_cache_purge-$NGX_CACHE_PURGE_VERSION ngx_cache_purge && rm ngx_cache_purge-$NGX_CACHE_PURGE_VERSION.tar.gz

# Configure and build software
RUN cd /tmp/nginx && ./configure \
    --sbin-path=/usr/local/sbin \
    --conf-path=/etc/nginx/nginx.conf \
    --error-log-path=/var/log/nginx/error.log \
    --http-log-path=/var/log/nginx/access.log \
    --user=www-data \
    --group=www-data \
    --with-http_gzip_static_module \
    --with-http_ssl_module \
    --without-mail_pop3_module \
    --without-mail_imap_module \
    --without-mail_smtp_module \
    --add-module=/tmp/ngx_cache_purge && \
    make && make install


# Copy configuration files into /etc/nginx and clean up config file to point to
# correct root
RUN mv /usr/local/nginx/html /var/www && \
    chown -R www-data:www-data /var/www && \
    mkdir -p /etc/nginx/conf.d && \
    sed -i 's/root\s*html;.*$/root \/var\/www;/' /etc/nginx/nginx.conf && \
    echo "\ninclude /etc/nginx/conf.d/\*.conf;" >> /etc/nginx/nginx.conf

# Forward request and error logs to docker log collector
RUN ln -sf /dev/stdout /var/log/nginx/access.log && \
    ln -sf /dev/stderr /var/log/nginx/error.log

VOLUME ["/var/cache/nginx"]

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]
