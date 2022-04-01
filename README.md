# GraphQL with GORM

<p align="justify">Gorm merupakan ORM yang dikembangkan untuk bahasa GO, seperti halnya SQLAlchemy pada bahasa python. Golang juga mendukung proses auto migrations, ini adalah alat bantu yang cukup keren yang berfungsi sebagai alat bantu untuk mempercepat kerja developer.
<br>
(Sumber: https://halovina.com/relasi-table-pada-golang-dengan-gorm)
</p>

### Command To Use

- Buat file direktory workspace

```
$ mkdir graphql-gorm
$ cd graphql-gorm
$ go mod init github.com/[username]/graphql-gorm
```

- download dependensi library GORM dan mysql driver yang akan digunakan

```
$ go get github.com/jinzhu/gorm
$ go get github.com/go-sql-driver/mysql
```
