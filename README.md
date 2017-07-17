# POC Caddy OAuth

## Components

- backend
  it's only example test component, backend http service, which will print current time, hostname and request headers

- caddy
  plugins: 
  - https://github.com/BTBurke/caddy-jwt - verification of JWT tokens
  - https://github.com/tarent/loginsrv - login and creation of JWT tokens

## Pros and cons

Cons:
- not perfect for web app, problem with 301/302 redirect, resource path issues, ...
- can't force logout action - token must expiry, or verification secret must be changed (all user will be logged out as well)

Pros:
- caddy (golang) - lightweight, image size 18MB, initial memory consumption ~5MB
- horizontal scalability
- JWT - fast local authentification
- perfect for API microservice calls

## Stress test

50MB RAM + 3x vCPU => ~2000 req/sec with median 100ms latency:

```
[root@dockerhost caddy]# ab -n 40000 -c 300 -C "jwt_token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJqYW5nYXJhaiIsInBpY3R1cmUiOiJodHRwczovL2F2YXRhcnM1LmdpdGh1YnVzZXJjb250ZW50LmNvbS91LzExODI5MzI_dj00IiwibmFtZSI6IkphbiBHYXJhaiIsImVtYWlsIjoiamFuLmdhcmFqQGdtYWlsLmNvbSIsIm9yaWdpbiI6ImdpdGh1YiIsImV4cCI6MTUwMDM3NTEyNX0.MEahrE5-PIcxrL6qrpHUGFV3qKYHj660nmlKo2hzOHMtNA_5dkKRM1EEjv3CMWGQCEQvjQB5OWUeghCkP7f1Cw" http://192.168.32.128:2015/
This is ApacheBench, Version 2.3 <$Revision: 1430300 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 192.168.32.128 (be patient)
Completed 4000 requests
Completed 8000 requests
Completed 12000 requests
Completed 16000 requests
Completed 20000 requests
Completed 24000 requests
Completed 28000 requests
Completed 32000 requests
Completed 36000 requests
Completed 40000 requests
Finished 40000 requests


Server Software:        Caddy
Server Hostname:        192.168.32.128
Server Port:            2015

Document Path:          /
Document Length:        869 bytes

Concurrency Level:      300
Time taken for tests:   19.287 seconds
Complete requests:      40000
Failed requests:        0
Write errors:           0
Total transferred:      45360000 bytes
HTML transferred:       34760000 bytes
Requests per second:    2073.95 [#/sec] (mean)
Time per request:       144.652 [ms] (mean)
Time per request:       0.482 [ms] (mean, across all concurrent requests)
Transfer rate:          2296.74 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   11  94.4      1    1006
Processing:     2  133 252.4     94    2179
Waiting:        2  132 252.2     93    2178
Total:          3  144 277.4     96    2891

Percentage of the requests served within a certain time (ms)
  50%     96
  66%    110
  75%    121
  80%    127
  90%    147
  95%    169
  98%   1128
  99%   2043
 100%   2891 (longest request)


$ docker stats
CONTAINER           CPU %               MEM USAGE / LIMIT     MEM %               NET I/O             BLOCK I/O           PIDS
caddy               247.21%             41.95MiB / 3.686GiB   1.11%               0B / 0B             0B / 0B             12
```

## Alternatives

- https://github.com/gambol99/keycloak-proxy (Golang) 
- https://github.com/bitly/oauth2_proxy (Golang)
- https://github.com/jaimejorge/rig (Golang)
