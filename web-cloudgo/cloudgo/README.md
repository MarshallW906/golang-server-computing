# Simple HTTP Server

<!-- TOC -->

- [Simple HTTP Server](#simple-http-server)
    - [Create a simple Http Server with net/http](#create-a-simple-http-server-with-nethttp)
        - [test with curl](#test-with-curl)
        - [test with ab](#test-with-ab)
    - [A prodoct-level Http Server](#a-prodoct-level-http-server)
        - [test with curl](#test-with-curl-1)
        - [test with ab](#test-with-ab-1)

<!-- /TOC -->

## Create a simple Http Server with net/http

You just need to call `http.HandleFunc()` and `http.ListenAndServe()`, then implement a Handle Function.

```go
func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()                   // Parse (will not automatic parse unless you call it manually)
    fmt.Println("r.Form:", r.Form)  // Print out the Form to Server Side
    fmt.Println("path", r.URL.Path) // Print out the Path of the Request
    // Print out the form submitted via GET
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ", "))
    }
    // Write Response back to Client (Will be displayed on Browser)
    fmt.Fprintf(w, "Hello %+v!\n", r.Form.Get("name"))
}

func main() {
    http.HandleFunc("/", sayhelloName)       // Set Router
    err := http.ListenAndServe(":9090", nil) // Set Listening Port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
```

### test with curl

```
✗ curl -v "http://localhost:9090/?url_long=111&url_long=222&name=aab"

*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 9090 (#0)
> GET /?url_long=111&url_long=222&name=aab HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Tue, 14 Nov 2017 12:10:45 GMT
< Content-Length: 11
< Content-Type: text/plain; charset=utf-8
<
Hello aab!
* Connection #0 to host localhost left intact
```

### test with ab

```
✗ ab -n 1000 -c 100 "http://localhost:9090/?url_long=111&url_long=222&name=aab"
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

Document Path:          /?url_long=111&url_long=222&name=aab
Document Length:        11 bytes

Concurrency Level:      100
Time taken for tests:   0.541 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      128000 bytes
HTML transferred:       11000 bytes
Requests per second:    1847.26 [#/sec] (mean)
Time per request:       54.134 [ms] (mean)
Time per request:       0.541 [ms] (mean, across all concurrent requests)
Transfer rate:          230.91 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        1   22  15.5     18      77
Processing:     2   31  18.0     29      82
Waiting:        1   14  13.0     10      75
Total:          9   53  21.2     53     104

Percentage of the requests served within a certain time (ms)
  50%     53
  66%     60
  75%     65
  80%     73
  90%     84
  95%     86
  98%     95
  99%     97
 100%    104 (longest request)
```

## A prodoct-level Http Server
Code files are stored in this folder.

`main.go` simply sets port & applies a server to the port, calling service/server to create a server

`service/server.go` Create a Server and add routes to it, then return to the caller.

### test with curl
```
✗ curl -v http://localhost:4396/hello/clearlove
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 4396 (#0)
> GET /hello/clearlove HTTP/1.1
> Host: localhost:4396
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Tue, 14 Nov 2017 13:15:55 GMT
< Content-Length: 32
<
{
  "Test": "Hello clearlove"
}
* Connection #0 to host localhost left intact
```

### test with ab

```
✗ ab -n 1000 -c 100 http://localhost:4396/hello/clearlove
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
Server Port:            4396

Document Path:          /hello/clearlove
Document Length:        32 bytes

Concurrency Level:      100
Time taken for tests:   0.433 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      155000 bytes
HTML transferred:       32000 bytes
Requests per second:    2307.57 [#/sec] (mean)
Time per request:       43.336 [ms] (mean)
Time per request:       0.433 [ms] (mean, across all concurrent requests)
Transfer rate:          349.29 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        2   19   7.3     18      37
Processing:     3   23   7.6     23      46
Waiting:        2   15   8.2     13      31
Total:         14   42   6.0     43      57

Percentage of the requests served within a certain time (ms)
  50%     43
  66%     45
  75%     46
  80%     47
  90%     49
  95%     50
  98%     52
  99%     53
 100%     57 (longest request)
```