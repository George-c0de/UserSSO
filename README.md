# Сервис авторизации

## Основные требования:

* Авторизация (Auth)
* Работа с пермишеннами (Permissions)
* Предоставление информации о пользователе (User Info)

### Генерация Proto files

```bash
    protoc -I proto proto/sso/sso.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
```

#### Для более удобного запуска с помощью утилиты Task

```bash
  brew install go-task
  task generate
```

### Команда для запуска приложения go run cmd/sso/main.go --config=./config/local.yaml

```bash
go run cmd/sso/main.go --config=./config/local.yaml
```

## Зависимости

### Request

* Go

### Optional

* brew for install other packages
* task for Taskfile.yml

