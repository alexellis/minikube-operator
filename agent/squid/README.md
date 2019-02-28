Build

```
docker build -t alexellis2/squid-proxy:0.2 . \
&& docker push alexellis2/squid-proxy:0.2
```

Test

```
docker run --rm --name squid -p 3129:3129 -ti alexellis2/squid-proxy:0.2
```

In another terminal:

```
curl google.com -x 127.0.0.1:3129
```