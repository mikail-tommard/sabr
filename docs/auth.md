# Auth Service

## Назначение

`auth` отвечает только за identity и аутентификацию.

Что находится в зоне ответственности `auth`:

- регистрация identity
- хранение `email`
- хранение `password_hash`
- хранение `role`
- выпуск access token
- выпуск refresh token
- обновление access token через refresh token
- публикация события о регистрации пользователя в `users`

Что `auth` не хранит:

- `name`
- `username`
- профиль пользователя
- кампус
- avatar
- bio

Эти данные должны жить в `users`.

## Архитектура

Структура `auth` построена по слоям:

- `domain`
  - сущности и domain errors
- `usecase`
  - бизнес-сценарии
  - порты `IdentityRepository`, `RefreshTokenRepository`, `EventPublisher`
- `adapters/postgres`
  - работа с PostgreSQL
- `adapters/http`
  - HTTP handlers и router
- `adapters/events/http`
  - HTTP publisher для отправки события в `users`

Точка сборки модуля:

- [backend/services/auth/module.go](/Users/Shared/sabr/backend/services/auth/module.go:1)

Точка входа приложения:

- [backend/cmd/api/api.go](/Users/Shared/sabr/backend/cmd/api/api.go:1)

## Domain

Основная сущность:

- [backend/services/auth/internal/domain/identity.go](/Users/Shared/sabr/backend/services/auth/internal/domain/identity.go:1)

`Identity` содержит:

- `ID`
- `Email`
- `PasswordHash`
- `Role`
- `CreatedAt`
- `UpdatedAt`

Также в domain есть:

- `RefreshToken`
- `UserRegistered`

`UserRegistered` это событие, которое публикуется после успешной регистрации:

- `UserID`
- `Email`
- `Name`
- `Username`
- `OccurredAt`

## База данных

Сейчас `auth` хранит identities в таблице:

- `auth_identities`

Миграция:

- [backend/migrations/20260510220001_create_users.sql](/Users/Shared/sabr/backend/migrations/20260510220001_create_users.sql:1)

Поля `auth_identities`:

- `id`
- `email`
- `password_hash`
- `role`
- `created_at`
- `updated_at`

Refresh токены хранятся отдельно:

- [backend/migrations/20260510220002_create_refresh_tokens.sql](/Users/Shared/sabr/backend/migrations/20260510220002_create_refresh_tokens.sql:1)

Поля `refresh_tokens`:

- `id`
- `user_id`
- `token_hash`
- `expires_at`
- `created_at`
- `revoked_at`

Важно:

- refresh token хранится не в открытом виде
- в БД хранится только его SHA-256 hash

## Usecase

Основная бизнес-логика:

- [backend/services/auth/internal/usecase/service.go](/Users/Shared/sabr/backend/services/auth/internal/usecase/service.go:1)

### Register

`Register(input RegisterInput)` делает следующее:

1. нормализует `email`
2. нормализует `username`
3. проверяет, что `email` не занят
4. хеширует пароль через `bcrypt`
5. создает `Identity`
6. публикует событие `UserRegistered`
7. выдает access token и refresh token

В `auth` при регистрации сохраняются только identity-данные.

`name` и `username` не пишутся в таблицу `auth_identities`, а используются только для отправки события в `users`.

### Login

`Login(input LoginInput)`:

1. находит identity по `email`
2. сравнивает пароль с `password_hash`
3. выдает новую пару токенов

### Refresh

`Refresh(input RefreshInput)`:

1. берет `refreshToken`
2. вычисляет его hash
3. ищет запись в `refresh_tokens`
4. проверяет, что токен не отозван и не истек
5. отзывает старый refresh token
6. создает новую пару access/refresh token

То есть refresh token ротируется.

### Me

`Me(userID)`:

1. берет `userID` из middleware
2. ищет identity по id
3. возвращает текущую identity

## HTTP API

HTTP слой:

- [backend/services/auth/internal/adapters/http/handler.go](/Users/Shared/sabr/backend/services/auth/internal/adapters/http/handler.go:1)
- [backend/services/auth/internal/adapters/http/router.go](/Users/Shared/sabr/backend/services/auth/internal/adapters/http/router.go:1)

Маршруты монтируются в `api.go` под префиксом `/auth` через `http.StripPrefix("/auth", ...)`.

Итоговые endpoint'ы:

- `POST /auth/register`
- `POST /auth/login`
- `POST /auth/refresh`
- `GET /auth/me`

### POST /auth/register

Request:

```json
{
  "name": "Ivan",
  "username": "ivan_dev",
  "email": "ivan@example.com",
  "password": "secret123"
}
```

Response:

```json
{
  "user": {
    "id": "uuid",
    "email": "ivan@example.com",
    "role": "Student",
    "createdAt": "2026-05-11T10:00:00Z"
  },
  "tokens": {
    "accessToken": "...",
    "accessTokenExpiresAt": "2026-05-11T10:15:00Z",
    "refreshToken": "...",
    "refreshTokenExpiresAt": "2026-06-10T10:00:00Z"
  }
}
```

### POST /auth/login

Request:

```json
{
  "email": "ivan@example.com",
  "password": "secret123"
}
```

Response совпадает по структуре с `register`.

### POST /auth/refresh

Request:

```json
{
  "refreshToken": "..."
}
```

Response совпадает по структуре с `register`.

### GET /auth/me

Требует заголовок:

```text
Authorization: Bearer <accessToken>
```

Response:

```json
{
  "user": {
    "id": "uuid",
    "email": "ivan@example.com",
    "role": "Student",
    "createdAt": "2026-05-11T10:00:00Z"
  }
}
```

## Middleware

JWT middleware:

- [backend/pkg/middleware/auth.go](/Users/Shared/sabr/backend/pkg/middleware/auth.go:1)

Что делает `Auth(...)`:

1. читает `Authorization`
2. достает Bearer token
3. валидирует JWT
4. кладет в context:
   - `UserID`
   - `Role`

Также есть `RequireRoles(...string)` для проверки ролей.

Это позволит потом защищать маршруты по ролям, например `Admin`.

## JWT

JWT manager:

- [backend/pkg/jwt/manager.go](/Users/Shared/sabr/backend/pkg/jwt/manager.go:1)

В access token кладутся claims:

- `uid`
- `role`
- стандартные registered claims

Сейчас используется:

- `HS256`

## Интеграция с users

После успешной регистрации `auth` публикует событие в `users` через HTTP adapter:

- [backend/services/auth/internal/adapters/events/http/publisher.go](/Users/Shared/sabr/backend/services/auth/internal/adapters/events/http/publisher.go:1)

Вызов:

- `POST {USERS_SERVICE_URL}/internal/users/events/user-registered`

Payload:

```json
{
  "userId": "uuid",
  "email": "ivan@example.com",
  "name": "Ivan",
  "username": "ivan_dev",
  "occurredAt": "2026-05-11T10:00:00Z"
}
```

Важно:

- `users` не импортирует domain `auth`
- `users` должен принимать свой собственный request DTO
- `auth` отправляет integration payload, а не domain object наружу

Сейчас схема синхронная:

1. `auth` создает identity
2. `auth` вызывает `users`
3. если `users` вернул не-2xx или не ответил, `register` завершится ошибкой
4. если `users` ответил успешно, `auth` выдает токены

## Конфигурация

Конфиг:

- [backend/pkg/config/config.go](/Users/Shared/sabr/backend/pkg/config/config.go:1)

Обязательные env:

- `APP_ADDR`
- `DB_DSN`
- `JWT_SECRET`
- `USERS_SERVICE_URL`

Необязательные env:

- `ACCESS_TOKEN_TTL`
- `REFRESH_TOKEN_TTL`

Значения по умолчанию:

- `APP_ADDR=":8080"`
- `ACCESS_TOKEN_TTL=15m`
- `REFRESH_TOKEN_TTL=30d`

## Текущие ограничения

- `auth` не делает rollback identity, если `users` не принял событие после создания записи
- `register` сейчас синхронно зависит от `users`
- `me` возвращает только identity-данные, без профильной информации
- `auth` пока не реализует `logout`

## Что важно помнить

Главная идея текущей реализации:

- `auth` владеет identity
- `users` владеет profile
- связка идет через общий `userId`
- `auth` после регистрации сообщает `users`, что нужно создать профиль
