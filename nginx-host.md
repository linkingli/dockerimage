```
　　server { 
　　　　...
　　　　server_name www.abc.com; 
　　　　proxy_set_header Host $host;
　　　　proxy_set_header X-Real-IP $remote_addr;
　　　　proxy_set_header REMOTE-HOST $remote_addr;
　　　　proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
　　　　...
　　}
```
```
server {
  server_name  192.168.0.171;
  listen       8888;
  if ($http_Host !~*^192.168.0.171:8888$){
    return 403;
  }
  include /etc/nginx/default.d/*.conf;
  location / {
    root /www/dvwa;
    index index.php index.html index.htm;
  }
}
```
