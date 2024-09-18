# GoLangVue3Application

Following [this](https://www.udemy.com/course/working-with-vue-3-and-go/) course.

- [DemoApp](./DemoApp/) : This is a small books list app, to get familiar with Vue, created some components, fetched some data from an API.


## [VueGo Application](./VueGo/) : An app including DB, API, Vue Frontend.

#### Setup:
- inside vue-api : `go mod tidy`
- inside vue-application : `npm install`
- inside vue-api (for postgres database): `docker-compose up -d`

#### Postgres
- [Beekeeper Client](https://www.beekeeperstudio.io/download/?ext=exe&arch=&type=installer&edition=ultimate)
- Tables:
    - users:  
        `Email also must be unique`  
        ![alt text](/VueGo/vue-api/users_table.png)
    
    - tokens:  
        ![alt text](/VueGo/vue-api/tokens_table.png)
    
    - PK:
        ![alt text](/VueGo/vue-api/users_tokens_pk.png)
#### Runnig:
- API
    - start - `make start`
    - stop - `make stop`

- Vue App
    - start - `npm run serve`
