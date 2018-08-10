[LogDisplayPlatform](https://logdisplayplatform.com) [![Circle CI](https://circleci.com/gh/logdisplayplatform/logdisplayplatform.svg?style=svg)](https://circleci.com/gh/logdisplayplatform/logdisplayplatform) [![Go Report Card](https://goreportcard.com/badge/github.com/logdisplayplatform/logdisplayplatform)](https://goreportcard.com/report/github.com/logdisplayplatform/logdisplayplatform) [![codecov](https://codecov.io/gh/logdisplayplatform/logdisplayplatform/branch/master/graph/badge.svg)](https://codecov.io/gh/logdisplayplatform/logdisplayplatform)
================
[Website](https://logdisplayplatform.com) |
[Twitter](https://twitter.com/logdisplayplatform) |
[Community & Forum](https://community.logdisplayplatform.com)

LogDisplayPlatform is an open source, feature rich metrics dashboard and graph editor for
Graphite, Elasticsearch, OpenTSDB, Prometheus and InfluxDB.

![](http://docs.logdisplayplatform.org/assets/img/features/dashboard_ex1.png)

## Installation
Head to [docs.logdisplayplatform.org](http://docs.logdisplayplatform.org/installation/) and [download](https://logdisplayplatform.com/get)
the latest release.

If you have any problems please read the [troubleshooting guide](http://docs.logdisplayplatform.org/installation/troubleshooting/).

## Documentation & Support
Be sure to read the [getting started guide](http://docs.logdisplayplatform.org/guides/gettingstarted/) and the other feature guides.

## Run from master
If you want to build a package yourself, or contribute - Here is a guide for how to do that. You can always find
the latest master builds [here](https://logdisplayplatform.com/logdisplayplatform/download)

### Dependencies

- Go 1.10
- NodeJS LTS

### Building the backend
```bash
go get github.com/logdisplayplatform/logdisplayplatform
cd $GOPATH/src/github.com/logdisplayplatform/logdisplayplatform
go run build.go setup
go run build.go build
```

### Building frontend assets

For this you need nodejs (v.6+).

To build the assets, rebuild on file change, and serve them by LogDisplayPlatform's webserver (http://localhost:3000):
```bash
npm install -g yarn
yarn install --pure-lockfile
npm run watch
```

Build the assets, rebuild on file change with Hot Module Replacement (HMR), and serve them by webpack-dev-server (http://localhost:3333):
```bash
yarn start
# OR set a theme
env GRAFANA_THEME=light yarn start
```
Note: HMR for Angular is not supported. If you edit files in the Angular part of the app, the whole page will reload.

Run tests 
```bash
npm run jest
```

Run karma tests
```bash
npm run karma
```

### Recompile backend on source change

To rebuild on source change.
```bash
go get github.com/Unknwon/bra
bra run
```

Open logdisplayplatform in your browser (default: `http://localhost:3000`) and login with admin user (default: `user/pass = admin/admin`).

### Dev config

Create a custom.ini in the conf directory to override default configuration options.
You only need to add the options you want to override. Config files are applied in the order of:

1. logdisplayplatform.ini
1. custom.ini

In your custom.ini uncomment (remove the leading `;`) sign. And set `app_mode = development`.

### Running tests

#### Frontend
Execute all frontend tests
```bash
npm run test
```

Writing & watching frontend tests (we have two test runners)

- jest for all new tests that do not require browser context (React+more)
   - Start watcher: `npm run jest`
   - Jest will run all test files that end with the name ".jest.ts"
- karma + mocha is used for testing angularjs components. We do want to migrate these test to jest over time (if possible).
  - Start watcher: `npm run karma`
  - Karma+Mocha runs all files that end with the name "_specs.ts".

#### Backend
```bash
# Run Golang tests using sqlite3 as database (default)
go test ./pkg/... 

# Run Golang tests using mysql as database - convenient to use /docker/blocks/mysql_tests
GRAFANA_TEST_DB=mysql go test ./pkg/... 

# Run Golang tests using postgres as database - convenient to use /docker/blocks/postgres_tests
GRAFANA_TEST_DB=postgres go test ./pkg/... 
```

## Contribute

If you have any idea for an improvement or found a bug, do not hesitate to open an issue.
And if you have time clone this repo and submit a pull request and help me make LogDisplayPlatform
the kickass metrics & devops dashboard we all dream about!

## Plugin development

Checkout the [Plugin Development Guide](http://docs.logdisplayplatform.org/plugins/developing/development/) and checkout the [PLUGIN_DEV.md](https://github.com/logdisplayplatform/logdisplayplatform/blob/master/PLUGIN_DEV.md) file for changes in LogDisplayPlatform that relate to
plugin development.

## License

LogDisplayPlatform is distributed under Apache 2.0 License.

