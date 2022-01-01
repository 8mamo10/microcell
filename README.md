# microcell

```
$ sudo mkdir /var/www/html
$ sudo ln -s microcell/streaming/nginx/html/index.html /var/www/html/index.html
$ sudo mkdir /var/www/bin
$ sudo ln -s microcell/streaming/start-streaming.sh /var/www/bin/start-streaming.sh
$ sudo ln -s microcell/streaming/stop-streaming.sh /var/www/bin/stop-streaming.sh
$ sudo ln -s microcell/feeding/start-feeding.sh /var/www/bin/start-feeding.sh
$ sudo ln -s microcell/feeding/stop-feeding.sh /var/www/bin/stop-feeding.sh
```

```
$ sudo nginx -t
$ sudo systemctl restart nginx
```

```
$ go run microcell/web/server.go
$ go build -o server server.go
```
