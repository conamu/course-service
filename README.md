## Local Environment
to test stuff locally:
```
make run
```

## Authentication
#### Request Headers:
API Authentication Header:
```
X-KBU-Auth: abcdefghijklmnopqrstuvwxyz
```
Login Token:
```
X-KBU-Login: TOKEN
```
## Endpoint Documentation

#### Create Course:
Requires `admin` role
```
/create | Responds with 201 Created if successfull

Payload:
{
    "title": "Testing",
    "subtitle": "Subtest",
    "description": "lorem ipsum dolor sit amet addendum",
    "instructor": "Hans Günther",
    "difficulty": 5,
    "fee": "230.50€",
    "certpath": "path/to/cert.pdf",
    "enlisted": "constantin,maria,karl",
    "likes": 43
}
```

#### Update Course:
Requires `admin` role
```
/update?courseId=1 | Responds with 200 OK if modified sucessfully

Payload:
{
    "title": "Testing",
    "subtitle": "Subtest",
    "description": "lorem ipsum dolor sit amet addendum",
    "instructor": "Hans Günther",
    "difficulty": 5,
    "fee": "230.50€",
    "certpath": "path/to/cert.pdf",
    "enlisted": "constantin,maria,karl",
    "likes": 43
}
```

#### Delete Course:
Requires `admin` role
```
/delete?courseId=1 | Responds with 200 OK 
```

#### Get Course List:
Requires to be logged in as user
without pageLength attribute it defaults to 5
```
/courses?pageLength=5
```
returns an array of Minimal Course Entry

#### Get Course:
Requires to be logged in as user
```
/course?courseId=1 | Responds 200 OK and course
                   | Responds 404 Not Found if course not found
```

#### Ping:
Pings, requires no login
```
/ping | responds 200 OK and "pong"
```