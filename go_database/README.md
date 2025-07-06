## How to run
### Setup docker for PostgreSQL
1. docker pull postgres
1. docker container run -v \$PWD/sql:/sql --name go-learn -e POSTGRES_PASSWORD=<mysecretpassword> -d postgres
1. docker container exec -it go-learn bash -c 'su postgres -c psql'
1. Create role
    ```psql
    postgres=# CREATE ROLE root LOGIN CREATEDB CREATEROLE PASSWORD '<mysecretpassword>';
    postgres=# \q
    ```
1. docker container exec -it go-learn bash -c 'createuser -d -P golearn'
1. Enter password and confirm
1. docker container exec -it go-learn bash -c 'createdb golearn' 
1. docker container exec -it go-learn bash -c 'psql -U golearn -f /sql/create.sql -d golearn'
### Intall lib/pg
1. go get github.com/lib/pq
