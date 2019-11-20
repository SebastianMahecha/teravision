# TeraVision Test

## Steps to install the project

1. Install go and configure GOPATH and GOROOT environment variables \
    `https://golang.org/doc/install` \
2. Create folder \
    `{{project_name}}/src/github.com` \
3. Clone repository \
    If you have ssh key in your PC `git@github.com:SebastianMahecha/teravision.git` \
    Else `https://github.com/SebastianMahecha/teravision.git` \
4. If you do not have a database in the cloud for testing, then create a postgres database in your local. Skip the first two commands if you already have postgres installed. \
    `sudo apt-get update` \
    `sudo apt-get install postgresql postgresql-contrib` v
    `sudo -u postgres psql` \
    `CREATE USER teravision_user WITH PASSWORD 'fDG6k6nAb#YG';` \
    `CREATE DATABASE teravision_db;` \
    `GRANT ALL PRIVILEGES ON DATABASE "teravision_db" to teravision_user;` \
    `\q` \
5. Configure your local or remote database access in file: \
    `./config/config.env` \

### Normal way to run the project

1. Install dependencies \
    `go get -u -v github.com/gin-gonic/gin` \
    `go get -u -v github.com/jinzhu/gorm` \
    `go get -u -v github.com/jinzhu/gorm/dialects/postgres` \
    `go get -u -v github.com/lib/pq` \
2. Run project \
    `go run application.go` \

### Run the project with Docker

1. Update docker-compose \
    `docker-compose up` \
2. Run project \
    `docker-compose run teravision` \

### If the service is already running, then try the api.

1. Import the following collection into PostMan: \
    `teravision.postman_collection.json` \
2. Parameterize the {{domain}} variable of your collection. Example: \
    `http://127.0.0.1:3000/api/v1/user` \
    `{
        "name": "Juan Sebastian Mahecha Macias",
        "fiscal_number": "10706541587",
        "birthdate": "1996-12-29"
    }` \

3. Modify the json of the body and send several requests \

### Send me your feedback