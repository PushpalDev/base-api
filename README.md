# base-API  [![Build Status](https://travis-ci.com/pushpaldev/base-api.svg?token=AbEANjysKDJ24sgJwcmH&branch=master)](https://travis-ci.com/pushpaldev/base-api)

This is the main repo of the base API. This API is a framework to build an API for a project. It contains the basics like users management, stripe implementation, authentication with JWT and security features like rate limits and access tokens management. 

## How to install

Set up and run MongoDB on your computer.
Set up and run Redis on your computer.

Create 2 environment files, .env and .env.prod - Use .env.sample skeleton to fill them and make sure to configure them correctly. As you can see default port for MongoDB is 27017, default port for Redis is 6379 and the application listens on port 4000. Feel free to update the ports with your own configuration. 

Set up Golang on your computer. Clone the project in your gopath and run the following commands at the project's root level:

openssl genrsa -out base.rsa 1024

openssl rsa -in base.rsa -pubout > base.rsa.pub

You can try to run the integration tests to make sure that everything is fine (Redis must be running since there is an integration test for rate limits): BASEAPI_ENV=testing go test ./...

If something fails look at the error and make sure that you didn't miss any configuration. Also, feel free to send an email at support@pushpal.io



