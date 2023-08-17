## Homework 8 

### Nginx Fine Tuning

Configure nginx that will cache only images, that were requested at least twice

Add ability to drop nginx cache by request.

You should drop cache for specific file only ( not all cache )

Implementation:

I used two servers: content server and proxy server

For proxy server I used https://github.com/nginx-modules/ngx_cache_purge plugin with custom Docker image to enable purge cache content. 
Had similar implementation with lua code, but diceded to practice custom nginx build

