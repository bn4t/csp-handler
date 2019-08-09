# CSP-Handler
*A simple web application to send csp reports to a email address*


### Important 
**CSP-Handler needs to be behind a reverse proxy which forwards either the `X-Forwarded-For` or `X-Real-IP` header, else ratelimiting won't work.**