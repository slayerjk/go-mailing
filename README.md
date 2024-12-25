# go-mailing
Go - va mailing helper functions for standart "net/smtp" package

Email using builtin "net/smtp" package based on appname, message, date and smtp data that could be some json file.

<h3>Mailing json file example to use with package</h3>

```
{
    "host": "my-smtp@example.com",
    "port": 25,
    "from_addr": "my-app@example.com",
    "to_addr_errors": [
        "admin-1@example.com",
        "admin-2@example.com"
    ],
    "to_addr_reports": [
        "user-1@example.com",
        "user-2@example.com"
    ]
}
```
