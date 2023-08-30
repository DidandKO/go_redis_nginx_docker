docker-compose up -d
curl -i -L -d "foo=bar" 127.0.0.1:8089/set_key
curl -i 127.0.0.1:8089/get_key?key=foo
curl -i -L -d "key=foo" 127.0.0.1:8089/del_key
