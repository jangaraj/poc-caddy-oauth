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

50MB RAM consumption + 3x vCPU => 4000+ req/sec:

```
[root@dockerhost caddy]# ab -n 100000 -c 600 -C "jwt_token=eyJhbGciOiJIUzUxMiIsI                                                                             nR5cCI6IkpXVCJ9.eyJzdWIiOiJqYW5nYXJhaiIsInBpY3R1cmUiOiJodHRwczovL2F2YXRhcnM1Lmdp                                                                             dGh1YnVzZXJjb250ZW50LmNvbS91LzExODI5MzI_dj00IiwibmFtZSI6IkphbiBHYXJhaiIsImVtYWls                                                                             IjoiamFuLmdhcmFqQGdtYWlsLmNvbSIsIm9yaWdpbiI6ImdpdGh1YiIsImV4cCI6MTUwMDM3NTEyNX0.                                                                             MEahrE5-PIcxrL6qrpHUGFV3qKYHj660nmlKo2hzOHMtNA_5dkKRM1EEjv3CMWGQCEQvjQB5OWUeghCk                                                                             P7f1Cw" http://192.168.32.128:2015/
This is ApacheBench, Version 2.3 <$Revision: 1430300 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 192.168.32.128 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        Caddy
Server Hostname:        192.168.32.128
Server Port:            2015

Document Path:          /
Document Length:        16 bytes

Concurrency Level:      600
Time taken for tests:   24.242 seconds
Complete requests:      100000
Failed requests:        0
Write errors:           0
Non-2xx responses:      100000
Total transferred:      32200000 bytes
HTML transferred:       1600000 bytes
Requests per second:    4125.15 [#/sec] (mean)
Time per request:       145.449 [ms] (mean)
Time per request:       0.242 [ms] (mean, across all concurrent requests)
Transfer rate:          1297.17 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   67 291.0      2    3017
Processing:     5   76  33.3     70     484
Waiting:        4   74  33.1     68     483
Total:          7  143 295.6     75    3195

Percentage of the requests served within a certain time (ms)
  50%     75
  66%     91
  75%    103
  80%    111
  90%    143
  95%   1051
  98%   1103
  99%   1133
 100%   3195 (longest request)

$ docker stats
CONTAINER           CPU %               MEM USAGE / LIMIT     MEM %               NET I/O             BLOCK I/O           PIDS
caddy               287.21%             44.38MiB / 3.686GiB   1.18%               0B / 0B             0B / 0B             12
```

## Alternatives

- https://github.com/gambol99/keycloak-proxy (Golang) 
- https://github.com/bitly/oauth2_proxy (Golang)
- https://github.com/jaimejorge/rig (Golang)
