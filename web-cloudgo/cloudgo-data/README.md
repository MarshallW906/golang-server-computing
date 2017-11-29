# cloudgo-data

Go "with" xorm & mysql

<!-- TOC -->

- [cloudgo-data](#cloudgo-data)
    - [Comparisons between xorm & database/sql](#comparisons-between-xorm--databasesql)
    - [TEST Using xorm](#test-using-xorm)
        - [test POST : insert 1000 entries](#test-post--insert-1000-entries)
        - [test GET FindById](#test-get-findbyid)
        - [test GET FindAllUsers](#test-get-findallusers)
    - [TEST Using database/sql](#test-using-databasesql)
        - [test POST : insert 1000 entries](#test-post--insert-1000-entries-1)
        - [test GET FindById](#test-get-findbyid-1)
        - [test GET FindAllUsers](#test-get-findallusers-1)

<!-- /TOC -->

## Comparisons between xorm & database/sql

1. The coding progress with **xorm** is much faster than **database/sql**;

2. When using xorm, we don't need to create a DAO Service, we just need to create a struct and call `Insert()`, `Get()` etc, and it's quite convenient;

3. But using xorm means sacrifice some efficiency for the development speed. In the following test, I found that it takes much more time to `FindAllUsers` using `xorm` than `database/sql`(2115 vs 1119).

## TEST Using xorm

### test POST : insert 1000 entries

```bash
✗  ab -n 1000 -c 100 -p post.txt  -T  'application/x-www-form-urlencoded'  "http://localhost:9090/service/userinfo"
This is ApacheBench, Version 2.3 <$Revision: 1796539 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            9090

Document Path:          /service/userinfo
Document Length:        111 bytes

Concurrency Level:      100
Time taken for tests:   0.830 seconds
Complete requests:      1000
Failed requests:        86
   (Connect: 0, Receive: 0, Length: 86, Exceptions: 0)
Non-2xx responses:      2
Total transferred:      243095 bytes
Total body sent:        194000
HTML transferred:       119111 bytes
Requests per second:    1204.62 [#/sec] (mean)
Time per request:       83.014 [ms] (mean)
Time per request:       0.830 [ms] (mean, across all concurrent requests)
Transfer rate:          285.97 [Kbytes/sec] received
                        228.22 kb/s sent
                        514.19 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    5   8.9      2      50
Processing:     2   74  54.9     69     238
Waiting:        2   70  52.2     68     233
Total:          2   79  54.9     78     243

Percentage of the requests served within a certain time (ms)
  50%     78
  66%     97
  75%    112
  80%    125
  90%    163
  95%    177
  98%    204
  99%    213
 100%    243 (longest request)
```

### test GET FindById

```bash
✗ ab -n 1000 -c 100 "http://localhost:9090/service/userinfo?userid=1929"
This is ApacheBench, Version 2.3 <$Revision: 1796539 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            9090

Document Path:          /service/userinfo?userid=1929
Document Length:        4244 bytes

Concurrency Level:      100
Time taken for tests:   0.781 seconds
Complete requests:      1000
Failed requests:        999
   (Connect: 0, Receive: 0, Length: 999, Exceptions: 0)
Non-2xx responses:      1000
Total transferred:      241123 bytes
HTML transferred:       108140 bytes
Requests per second:    1280.06 [#/sec] (mean)
Time per request:       78.121 [ms] (mean)
Time per request:       0.781 [ms] (mean, across all concurrent requests)
Transfer rate:          301.42 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   11  17.5      2      62
Processing:     1   64  68.6     35     343
Waiting:        1   58  63.2     31     323
Total:          1   75  68.4     72     347

Percentage of the requests served within a certain time (ms)
  50%     72
  66%     85
  75%    130
  80%    148
  90%    173
  95%    188
  98%    220
  99%    265
 100%    347 (longest request)
```

### test GET FindAllUsers

```bash
✗ ab -n 1000 -c 100 "http://localhost:9090/service/userinfo?userid="
This is ApacheBench, Version 2.3 <$Revision: 1796539 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            9090

Document Path:          /service/userinfo?userid=
Document Length:        4216 bytes

Concurrency Level:      100
Time taken for tests:   10.479 seconds
Complete requests:      1000
Failed requests:        998
   (Connect: 0, Receive: 0, Length: 998, Exceptions: 0)
Non-2xx responses:      2
Total transferred:      116413388 bytes
HTML transferred:       116310362 bytes
Requests per second:    95.43 [#/sec] (mean)
Time per request:       1047.903 [ms] (mean)
Time per request:       10.479 [ms] (mean, across all concurrent requests)
Transfer rate:          10848.80 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.0      0       5
Processing:     8 1013 381.9   1008    2115
Waiting:        7 1011 382.0   1005    1985
Total:         10 1014 381.8   1009    2115

Percentage of the requests served within a certain time (ms)
  50%   1009
  66%   1168
  75%   1273
  80%   1343
  90%   1523
  95%   1661
  98%   1788
  99%   1820
 100%   2115 (longest request)
```

## TEST Using database/sql

Code is in github.com/pmlpml/golang-learning/tree/master/web/cloudgo-data

### test POST : insert 1000 entries

```bash
✗ ab -n 1000 -c 100 -p post.txt  -T  'application/x-www-form-urlencoded'  "http://localhost:8080/service/userinfo"
This is ApacheBench, Version 2.3 <$Revision: 1796539 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo
Document Length:        4124 bytes

Concurrency Level:      100
Time taken for tests:   0.926 seconds
Complete requests:      1000
Failed requests:        997
   (Connect: 0, Receive: 0, Length: 997, Exceptions: 0)
Non-2xx responses:      4
Total transferred:      251914 bytes
Total body sent:        194000
HTML transferred:       127946 bytes
Requests per second:    1080.46 [#/sec] (mean)
Time per request:       92.553 [ms] (mean)
Time per request:       0.926 [ms] (mean, across all concurrent requests)
Transfer rate:          265.80 [Kbytes/sec] received
                        204.70 kb/s sent
                        470.50 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   3.1      1      41
Processing:     2   89  68.8     74     443
Waiting:        2   88  68.2     73     443
Total:          3   91  68.9     76     444

Percentage of the requests served within a certain time (ms)
  50%     76
  66%    118
  75%    141
  80%    153
  90%    177
  95%    215
  98%    247
  99%    296
 100%    444 (longest request)
```

### test GET FindById

```bash
✗ ab -n 1000 -c 100 "http://localhost:8080/service/userinfo?userid=1929"
This is ApacheBench, Version 2.3 <$Revision: 1796539 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=1929
Document Length:        4288 bytes

Concurrency Level:      100
Time taken for tests:   0.849 seconds
Complete requests:      1000
Failed requests:        989
   (Connect: 0, Receive: 0, Length: 989, Exceptions: 0)
Non-2xx responses:      1000
Total transferred:      283095 bytes
HTML transferred:       150299 bytes
Requests per second:    1178.16 [#/sec] (mean)
Time per request:       84.878 [ms] (mean)
Time per request:       0.849 [ms] (mean, across all concurrent requests)
Transfer rate:          325.71 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    8  13.7      1      48
Processing:     1   75  70.5     46     335
Waiting:        1   71  69.0     39     335
Total:          1   83  69.6     69     335

Percentage of the requests served within a certain time (ms)
  50%     69
  66%    118
  75%    150
  80%    160
  90%    172
  95%    198
  98%    212
  99%    237
 100%    335 (longest request)
```

### test GET FindAllUsers

```bash
✗ ab -n 1000 -c 100 "http://localhost:8080/service/userinfo?userid="
This is ApacheBench, Version 2.3 <$Revision: 1796539 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=
Document Length:        4307 bytes

Concurrency Level:      100
Time taken for tests:   4.084 seconds
Complete requests:      1000
Failed requests:        996
   (Connect: 0, Receive: 0, Length: 996, Exceptions: 0)
Non-2xx responses:      11
Total transferred:      111463760 bytes
HTML transferred:       111360617 bytes
Requests per second:    244.87 [#/sec] (mean)
Time per request:       408.372 [ms] (mean)
Time per request:       4.084 [ms] (mean, across all concurrent requests)
Transfer rate:          26654.95 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.0      0       5
Processing:     6  395 165.0    384    1119
Waiting:        6  394 164.9    383    1119
Total:          9  396 164.7    385    1119
WARNING: The median and mean for the initial connection time are not within a normal deviation
        These results are probably not that reliable.

Percentage of the requests served within a certain time (ms)
  50%    385
  66%    468
  75%    532
  80%    555
  90%    606
  95%    644
  98%    737
  99%    805
 100%   1119 (longest request)
```