# POC Caddy OAuth

## Components

- backend
  it's only example test component, backend http service, which will print hostname and request headers

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

