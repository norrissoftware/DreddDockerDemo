# DreddDockerDemo
A Quick Demo to get Dredd API testing in Docker.

This demo will build a docker container for [Dredd](https://dredd.readthedocs.io/en/latest/)
API testing with a simple [Go](https://golang.org/)
web application and utilizing [Node.js](https://dredd.readthedocs.io/en/latest/hooks-nodejs/) hooks for Dredd.

To run test
```
make dreddtest
```

Requires
- Docker

Travis CI [![Build Status](https://travis-ci.org/jasonrichardsmith/DreddDockerDemo.svg?branch=master)](https://travis-ci.org/jasonrichardsmith/DreddDockerDemo)

Circle CI [![CircleCI](https://circleci.com/gh/jasonrichardsmith/DreddDockerDemo.svg?style=svg)](https://circleci.com/gh/jasonrichardsmith/DreddDockerDemo)

Codeship [ ![Codeship Status for jasonrichardsmith/DreddDockerDemo](https://app.codeship.com/projects/acd30640-67f0-0134-445a-2e0f9d8470fa/status?branch=master)](https://app.codeship.com/projects/176326)
