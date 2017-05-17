## 根据swagger spec 生成markdown文档

```sh
go get -v github.com/hudangwei/swag2md
```

```sh
./swag2md -h
NAME:
   swagger2markdown - generate markdown with swagger spec

USAGE:
   swag2md [global options] command [command options] [arguments...]

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
swag2md -f http://xxxxxxxx.com/v2/api-docs
```