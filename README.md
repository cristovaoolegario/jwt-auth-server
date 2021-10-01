# free-auth-server
An auth server to use in my other projects

![alt diagram](static/jwt.png)
We have 3 endpoints:

- api/v1/register - it's used to register a new user to the database;
- api/v1/login - return a valid jwt token for the user;
- api/v1/keys - return the server public keys that could be used to validate the token.