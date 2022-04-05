# GraphQL with GORM

<p align="justify">Gorm merupakan ORM yang dikembangkan untuk bahasa GO, seperti halnya SQLAlchemy pada bahasa python. Golang juga mendukung proses auto migrations, ini adalah alat bantu yang cukup keren yang berfungsi sebagai alat bantu untuk mempercepat kerja developer.
<br><br>
(Sumber: https://halovina.com/relasi-table-pada-golang-dengan-gorm)
</p>

### Command To Use

- Buat file direktory workspace

```
$ mkdir graphql-gorm
$ cd graphql-gorm
$ go mod init github.com/[username]/graphql-gorm
```

- Download dependensi library GORM dan mysql driver yang akan digunakan

```
$ go get github.com/jinzhu/gorm
$ go get github.com/go-sql-driver/mysql
```

- Download library GraphQL dari 99designs

```
$ go get github.com/99designs/gqlgen
```

- Terapkan skeleton/struktur folder dari 99designs

```
$ go run github.com/99designs/gqlgen init
```
more info about connecting mongoDB<a href="https://medium.com/programmer-geek/menggunakan-mongodb-pada-golang-ad9eb9ed08c">Here</a><br>
more info about GraphQL <a href="https://github.com/99designs/gqlgen"><b>Here</b></a>
<br>
Thanks, <br>
Alfian
