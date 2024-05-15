# Grpc-gateway simple example

Простой пример "Приветствие". Описание простого proto файла, генерация кода с помощью компилятора buf. И простой код для запуска grpc сервера и клиента с проксирование http запросов.

Из особенностей. Все сторонние зависимости перенесены в репозиторий поскольку существует ограничения на запуск `buf mod update` команда не выполниться, из за ограничения доступа.

Так же разнесена генерация клиента, сервера и api swagger на три различных конфигурации. И создание двух команд генерации под клиент и сервер. Как продолжение темы клиент может быть как отдельный модуль go со своей версией что позволяет сторонним сервисам легко переиспользовать клиента, не опасаясь за целостность и актуальность.

*ps*
Много было сложностей с работой buf и настройкой модулей buf, в особенности с ограничением доступа к документации и некоторым командам комп. buf.

# Перегенерация grpc

*сервер и описание swagger*
```bash
make buf-gen
```

*client*
```bash
make buf-gen-client
```

Более подробно о работе команд в файл `Makefile`

# Как запустить


```bash
make run
```

Пробный запрос после запуска.

```bash
curl -X GET  http://localhost:8090/v1/sayhello?name=vasa
```

```bash
{"message":"vasa world"}
```

# References

- [Introduction to the gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/#introduction-to-the-grpc-gateway)
- [For more examples (it's docs gRPC-Gateway)](https://github.com/johanbrandhorst/grpc-gateway-boilerplate)
- [Source code of example that taken for a basis](https://github.com/iamrajiv/helloworld-grpc-gateway)