# CSP-Handler
*A simple web application to send csp reports to a email address*


### Important 
**CSP-Handler needs to be behind a reverse proxy which forwards either the `X-Forwarded-For` or `X-Real-IP` header, else ratelimiting won't work.**



## Setup


1. Clone the repository and enter the directory: `git clone https://git.bn4t.me/bn4t/csp-handler.git
 && cd csp-handler`
2. Edit the environment variables in `docker-compose.yml`
3. Build the image and start the container: `docker-compose up --build -d`

## Usage
Include the `report-uri` directive in your content security policy:

`report-uri https://csp-report.example.com/report-uri/mydomain.com`

Replace `mydomain.com` with the domain on which this content security policy is deployed.

## License

GPLv3
