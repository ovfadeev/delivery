# Настройки сервера

### Настройки подключения к серверу устанавливаются в файле:

```
config/.env
```

Пример заполнения в файле:

```
config/.env.example
```

### БД:

- https://www.postgresql.org/

### Необходимо для указанной в .env БД установить расширения:

```
CREATE EXTENSION "uuid-ossp";
CREATE EXTENSION "plpgsql";
```