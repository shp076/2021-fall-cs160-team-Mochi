# 2021-fall-cs160-team-Mochi

2021 Fall CS160 Team Mochi

##### Team Members

1. April Chao
2. Feng Zhang
3. Fudong Huang
4. Xiaoshu Xiao
5. Shuang Pan

### Run the program

##### 1. Start db server

Download PostgreSQL

```
https://www.postgresql.org/download/
```

Start the `psql` server and listen to port `5432`

##### 2. Start backend Server

Download Go env

```
https://golang.org/dl/
```

Create a `go/src` dic from your home directory, and clone the repo to the `src` repo.

```
git clone https://github.com/shuangpan5217/2021-fall-cs160-team-Mochi.git
```

Open `/2021-fall-cs160-team-Mochi/backend/source/generated/restapi/configure_coreapi.go`
For the following line,

```
db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=shuangpan user=shuangpan sslmode=disable")
```

change `dbname` and `user` as required, add `password=` if needed

```
cd 2021-fall-cs160-team-Mochi/backend/source/generated/cmd/coreapi-server
```

Run

```
1. go build
2. ./coreapi-server
```

The backend will listen to the `localhost:3000`

##### 3. Start frontend server

Have `node`, `npm` installed

<pre>
1. npm install
2. npm start
</pre>

Note: If there are any issues with any node modules, try running `npm cache clean --force` and deleting the `node_modules` folder before retrying the steps outlined above.

Listen to a different port if needed when run `npm start`.
Open your browser and input `localhost:front_end_port_number`

Now, you are ready to use the website.

### Testing
##### Front-End test automation, `Cypress`

1. Start db server
2. Clear any data in the db server (run `TRUNCATE TABLE users CASCADE;`)
3. Start backend Server
4. Start frontend server

<pre>
1. npm init -y
2. npm install cypress
3. npx cypress open
</pre>

Now, you can click the test file in `Cypress` to test.

##### Back-end test automation, go `testing` and `net/http/httptest` packages
<pre>
1. Start db server
2. From <b>2021-fall-cs160-team-Mochi/backend</b>, run <b>go test ./source/apis/notes/...</b>
3. From <b>2021-fall-cs160-team-Mochi/backend</b>, run <b>go test ./source/apis/usermgmt/...</b>
</pre>

Successful test results.
```
1. ok  	2021-fall-cs160-team-Mochi/backend/source/apis/notes	1.221s
2. ok  	2021-fall-cs160-team-Mochi/backend/source/apis/usermgmt	1.108s
```
