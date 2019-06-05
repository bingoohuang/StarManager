Star Manager
============

Organize, tag, and search your GitHub stars.


```bash
➜  StarManager git:(main) ✗ http ":8080/stars"
HTTP/1.1 200 OK
Content-Length: 93
Content-Type: text/plain; charset=utf-8
Date: Wed, 05 Jun 2019 11:27:01 GMT

[
    {
        "day": "2019-06-05T19:23:37+08:00",
        "description": "xxxx",
        "name": "bingoohuang",
        "url": "hahah"
    }
]


➜  StarManager git:(main) ✗ http --form POST :8080/stars name=dingoohuang description=xxxx url=hahah day="2019-06-01 19:23:37"
HTTP/1.1 201 Created
Content-Length: 0
Date: Wed, 05 Jun 2019 11:27:36 GMT
Location: /stars/dingoohuang


➜  StarManager git:(main) ✗ http ":8080/stars"
HTTP/1.1 200 OK
Content-Length: 185
Content-Type: text/plain; charset=utf-8
Date: Wed, 05 Jun 2019 11:27:57 GMT

[
    {
        "day": "2019-06-05T19:23:37+08:00",
        "description": "xxxx",
        "name": "bingoohuang",
        "url": "hahah"
    },
    {
        "day": "2019-06-01T19:23:37+08:00",
        "description": "xxxx",
        "name": "dingoohuang",
        "url": "hahah"
    }
]

➜  StarManager git:(main) ✗ http --form PUT ":8080/stars/dingoohuang" description="blabla..."
HTTP/1.1 204 No Content
Date: Wed, 05 Jun 2019 11:30:19 GMT


➜  StarManager git:(main) ✗ http ":8080/stars/dingoohuang"
HTTP/1.1 200 OK
Content-Length: 96
Content-Type: text/plain; charset=utf-8
Date: Wed, 05 Jun 2019 11:30:36 GMT

{
    "day": "2019-06-01T19:23:37+08:00",
    "description": "blabla...",
    "name": "dingoohuang",
    "url": "hahah"
}

➜  StarManager git:(main) ✗ http ":8080/db/stats"
HTTP/1.1 200 OK
Content-Length: 136
Content-Type: text/plain; charset=utf-8
Date: Wed, 05 Jun 2019 11:31:04 GMT

{
    "Idle": 0,
    "InUse": 0,
    "MaxIdleClosed": 0,
    "MaxLifetimeClosed": 2,
    "MaxOpenConnections": 10,
    "OpenConnections": 0,
    "WaitCount": 0,
    "WaitDuration": 0
}

```