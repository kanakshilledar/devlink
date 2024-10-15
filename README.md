[![Go](https://github.com/kanakshilledar/devlink/actions/workflows/go.yml/badge.svg)](https://github.com/kanakshilledar/devlink/actions/workflows/go.yml)

# Devlink

A simple webapp to find you upcoming conferences or meetups.

## Tech Stack

Built using Go as backend and Next.js as frontend.

## Building

To build this project make sure you have `Go` and `Next.js` installed.

### Running the backend

* Get your MongoDB connection URL and place it in the **.env** file.
* Add you JWT token secret KEY in the **.env** file.
* Run the project using

```shell
$ go run main.go
```

This will open a webserver on port `:8080`.

### Running the frontend

* Navigate to the `client/` directory and run

```shell
$ npm install
$ npm run dev
```

This will open the frontend on port `:3000`.

## Documentation

The documentation can be found in the `docs/` directory.