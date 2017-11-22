# Cloudgo-io

Note:

- [negroni-gzip](https://github.com/phyber/negroni-gzip)'s code analysis is in `gzip-code-analysis.md`

## Tests

- Static File Server

Use `curl -v` and see the response simply contain some `<a>` tags, which are File Server's showing file list.

```stdio
✗ curl -v http://localhost:9090/static
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 9090 (#0)
> GET /static HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: text/html; charset=utf-8
< Last-Modified: Wed, 22 Nov 2017 05:08:40 GMT
< Date: Wed, 22 Nov 2017 05:13:22 GMT
< Content-Length: 127
<
<pre>
<a href="css/">css/</a>
<a href="favicon.ico">favicon.ico</a>
<a href="images/">images/</a>
<a href="js/">js/</a>
</pre>
* Connection #0 to host localhost left intat
```

- Simple js Request

I add an `asset/js/apitest.js` to call `service/apiTest.go` for a random integer between [0, 2000)

```out
✗ curl -v http://localhost:9090/
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 9090 (#0)
> GET / HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: text/html; charset=UTF-8
< Date: Wed, 22 Nov 2017 05:35:09 GMT
< Content-Length: 890
<
<html>

<head>
    <link rel="stylesheet" href="static/css/main.css" />
    <script src="http://code.jquery.com/jquery-latest.js"></script>
    <script src="static/js/apitest.js"></script>
</head>

<body>
    <img src="static/images/cng.png" height="48" width="48" /> Sample Go Web Application!!
    <div>
        <p class="greeting-id">The ID is 8675309</p>
        <p class="greeting-content">The content is Hello from Go! of homeHandler</p>
        <p class="rand">Get a random number by js: </p>
    </div>
    <div>
        <form action="/checkform" method="post">
            <p>用户名: <input type="text" name="username"></p>
            <p>密码: <input type="password" name="password"></p>
            <p><input type="hidden" name="token" value="5ebd2b356155437a4833044860b40d2d"></p>
            <input type="submit" value="登录">
        </form>
    </div>
</body>

* Connection #0 to host localhost left intact
</html>%
```

- Receive a form and return a table

`checkform.go` is the handler for a form submit. Only supporting `POST` and if you visit it by `GET`, it will return a `307 Temporary Redirect` and redirects to `/unknown`, which returns a `501 Not Implemented`.

```out
✗ curl -v -d "username=aaa&password=bbb&token=curlpost" http://localhost:9090/checkform
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 9090 (#0)
> POST /checkform HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.54.0
> Accept: */*
> Content-Length: 40
> Content-Type: application/x-www-form-urlencoded
>
* upload completely sent off: 40 out of 40 bytes
< HTTP/1.1 200 OK
< Date: Wed, 22 Nov 2017 05:39:19 GMT
< Content-Length: 350
< Content-Type: text/html; charset=utf-8
<
<html>

<head>
    <link rel="stylesheet" href="static/css/main.css" />
</head>

<body>
    <table>
        <tr>
            <td>Username</td>
            <td>Password</td>
            <td>Token</td>
        </tr>
        <tr>
            <td>aaa</td>
            <td>bbb</td>
            <td>curlpost</td>
        </tr>
    </table>
</body>

* Connection #0 to host localhost left intact
</html>%
```

- /unknown returns status code `5xx`

In `unknown.go` I created a handler func `inDevelopment`, which simply returns a code `501 Not Implemented` and a string telling it's in development.

```out
✗ curl -v http://localhost:9090/unknown
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 9090 (#0)
> GET /unknown HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 501 Not Implemented
< Date: Wed, 22 Nov 2017 05:42:44 GMT
< Content-Length: 30
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
Sorry, this is in development.%
```