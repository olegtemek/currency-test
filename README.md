# Для того чтобы развернуть приложение

1. Нужен docker (compose)

2.

```bash
   cd <path_to_app> && make start
```

3. Миграции будут автоматически запущены

4. Сервис будет работать на порту :8000 и так же к нему добавил pgadmin для того, чтобы смотреть что в бд на порту :8001

5. Swagger документация по пути /swagger

6. Чтобы остановить сервис

```bash
make down
```

## Доп использованные библиотеки

1. pgx || pgx/pool
2. github.com/swaggo/swag|| github.com/swaggo/http-swagger
3. github.com/ilyakaznacheev/cleanenv
4. github.com/golang-migrate/migrate/v4

## Вопросы после ТЗ:

1. Почему ms sql server?
   > хотя на проекте postgres, mongo
2. Зачем хранить config в json?
   > почему не в .env?
