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

https://www.postgresql.org/

### Расширения для указанной в **.env** БД:

```
CREATE EXTENSION "uuid-ossp";
CREATE EXTENSION "plpgsql";
```