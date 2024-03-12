server {
	listen 80;
  # 這就是對應 linode domain 裡面綁定的 domain
	server_name mytest777.com;
	location / {
    # 這是對應我們的 web server
		proxy_pass http://localhost:3000;

		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header X-Forwarded-Proto $http_x_forwarded_proto;
	}
}