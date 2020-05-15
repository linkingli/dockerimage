```
FROM centos:centos7
RUN wget https://www.openssl.org/source/openssl-1.0.2j.tar.gz
RUN wget ftp://ftp.csx.cam.ac.uk/pub/software/programming/pcre/pcre-8.39.tar.gz
RUN wget http://zlib.net/zlib-1.2.8.tar.gz
RUN wget http://nginx.org/download/nginx-1.10.2.tar.gz
RUN git clone -b 2.2 https://github.com/vkholodkov/nginx-upload-module
RUN tar zxf openssl-1.0.2j.tar.gz
RUN tar zxf pcre-8.39.tar.gz
RUN tar zxf zlib-1.2.8.tar.gz
RUN tar zxf nginx-1.10.2.tar.gz
RUN cd nginx-1.10.2

RUN ./configure --add-module=../nginx-upload-module --with-openssl=../openssl-OpenSSL_1_0_2j --with-http_gzip_static_module --with-http_ssl_module --with-pcre=../pcre-8.39/ --with-zlib=../zlib-1.2.8
RUN make && sudo make install

RUN /usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf
RUN ln -s /usr/local/nginx/sbin/* /usr/local/sbin/
#EXPOSE 映射端口
EXPOSE 80 443
 
#CMD 运行以下命令
#CMD ["nginx"]
CMD ["/usr/local/nginx/sbin/nginx","-g","daemon off;"]
```
