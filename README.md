## 根据swagger spec 生成markdown文档

```sh
go get -v github.com/hudangwei/apidoctool
```

```sh
./apidoctool -h
NAME:
   apidoctool - generate markdown docs with swagger spec

USAGE:
   apidoctool [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --loglevel value, -l value  set log level (default: "debug") [$LOG_LEVEL]
   --FilePath value, -f value  swagger spec file or swagger spec url (default: "swagger.json") [$SWAGGER_SPEC_FILEPATH]
   --help, -h                  show help
   --version, -v               print the version
```

### 示例
```sh
apidoctool -f http://xxxxxxxx.com/v2/api-docs
或
apidoctool -f swagger.json
```