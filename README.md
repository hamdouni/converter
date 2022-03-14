# Generate Go struct from mysql table schema

## Usage
    ```sh
    go run cmd/cli.go -file model.go -dsn "us:pw@tcp(localhost:3306)/db" -table user
    ```

## Parameters
```sh
-dsn            string database dsn configuration
-enableJsonTag  bool whether to add json tag, default false
-file           string save path
-packageName    string package name
-prefix         string table prefix
-realNameMethod string The table name corresponding to the structure
-table          string table to migrate
-tagKey         string tag key, default orm
-dateToTime     bool whether to convert sql date to Time, default true
```
