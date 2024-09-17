# GoLangVue3Application

Following [this](https://www.udemy.com/course/working-with-vue-3-and-go/) course.

- [DemoApp](./DemoApp/) : This is a small books list app, to get familiar with Vue, created some components, fetched some data from an API.


## [VueGo Application](./VueGo/) : An app including DB, API, Vue Frontend.

#### Setup:
- inside vue-api : `go mod tidy`
- inside vue-application : `npm install`
- inside vue-api (for database): `docker-compose up -d`

#### Runnig:
- API
    - start - `make start`
    - stop - `make stop`