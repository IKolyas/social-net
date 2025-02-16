# social-net

## Настройка конфигурации Master -> Slave
https://github.com/OtusTeam/highload/blob/master/lessons/02/05/live/guide.md

## Развёртывание

```bash
make all
```

### Публичные маршруты (не требуют авторизации)

```
POST /login - авторизация пользователя
POST /user/register - регистрация нового пользователя
```

### Защищенные маршруты (требуют JWT токен)

#### Пользователи
```
GET /user/get/:id - получение информации о пользователе
GET /user/search - поиск пользователей
```

#### Друзья
```
PUT /friend/set/:user_id - добавить друга
PUT /friend/delete/:user_id - удалить из друзей
```

#### Посты
```
POST /post/create - создание поста
PUT /post/update - обновление поста
PUT /post/delete/:id - удаление поста
GET /post/get/:id - получение поста по ID
GET /post/feed - получение ленты постов (поддерживает параметры offset и limit)
```

Все защищенные маршруты требуют заголовок `Authorization: Bearer <token>`, который проверяется через `middleware.AuthMiddleware`.

Маршрутизация реализована с использованием фреймворка Gin и настроена в файле `server.go`. Каждый маршрут связан с соответствующим обработчиком из пакета `handlers`.