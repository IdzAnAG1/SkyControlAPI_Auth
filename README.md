# SkyControlAPI_Auth

[![Go Version](https://img.shields.io/badge/Go-1.25.4-blue.svg)](https://golang.org/)
[![gRPC](https://img.shields.io/badge/gRPC-1.79.1-green.svg)](https://grpc.io/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

> Микросервис аутентификации для экосистемы SkyControl. Предоставляет функционал регистрации и входа пользователей через gRPC API.

## Обзор

SkyControlAPI_Auth — высокопроизводительный сервис аутентификации, построенный на Go и gRPC. Обрабатывает регистрацию и аутентификацию пользователей для платформы SkyControl.

## Возможности

- **Регистрация пользователей** — Создание новых учетных записей с email, паролем, никнеймом и именем пользователя
- **Вход пользователей** — Аутентификация существующих пользователей и получение токенов доступа
- **gRPC API** — Высокопроизводительное взаимодействие на основе protocol buffers
- **Генерация кода Protobuf** — Автоматическая генерация кода с помощью Buf

## Структура проекта

```
.
├── cmd/main/                    # Точка входа приложения
│   └── main.go
├── generated/skycontrol/        # Сгенерированный код protobuf
│   └── proto/auth/v1/
│       ├── auth.pb.go
│       └── auth_grpc.pb.go
├── internal/
│   ├── entry/app/               # Инициализация приложения
│   │   └── app.go
│   └── logger/                  # Утилиты логирования
│       └── logger.go
├── buf.gen.yaml                 # Конфигурация генерации кода Buf
├── Dockerfile                   # Контейнерный образ
└── go.mod                       # Зависимости Go-модуля
```

## Методы API

| Метод | Запрос | Ответ | Описание |
|-------|--------|-------|----------|
| `Register` | `RegisterRequest` | `RegisterResponse` | Создание новой учетной записи |
| `Login` | `LoginRequest` | `LoginResponse` | Аутентификация пользователя и получение токена |

### Типы сообщений

**RegisterRequest**
- `email` (string) — Email-адрес пользователя
- `password` (string) — Пароль пользователя
- `nickname` (string) — Отображаемое имя
- `username` (string) — Уникальное имя пользователя

**RegisterResponse**
- `user_id` (int64) — ID созданного пользователя
- `token` (string) — Токен аутентификации
- `err_message` (string) — Сообщение об ошибке (если есть)

**LoginRequest**
- `email` (string) — Email-адрес пользователя
- `password` (string) — Пароль пользователя

**LoginResponse**
- `user_id` (int64) — ID аутентифицированного пользователя
- `token` (string) — Токен аутентификации
- `err_message` (string) — Сообщение об ошибке (если есть)

## Начало работы

### Требования

- Go 1.25.4 или новее
- Buf CLI (для генерации кода protobuf)
- Docker (опционально, для контейнеризации)

### Установка

```bash
# Клонирование репозитория
git clone https://github.com/IdzAnAG1/SkyControlAPI_Auth.git
cd SkyControlAPI_Auth

# Загрузка зависимостей
go mod download

# Генерация кода protobuf (при необходимости)
buf generate
```

### Локальный запуск

```bash
go run cmd/main/main.go
```

### Запуск через Docker

```bash
# Сборка образа
docker build -t skycontrol-auth .

# Запуск контейнера
docker run -p 50051:50051 skycontrol-auth
```

## Генерация кода

Этот проект использует [Buf](https://buf.build/) для генерации кода protobuf.

```bash
# Генерация Go-кода из protobuf-определений
buf generate
```

Конфигурация определена в [`buf.gen.yaml`](buf.gen.yaml).

## Зависимости

| Пакет | Версия | Назначение |
|-------|--------|------------|
| google.golang.org/grpc | v1.79.1 | gRPC фреймворк |
| google.golang.org/protobuf | v1.36.11 | Protocol buffers |

## Лицензия

Этот проект распространяется под лицензией MIT.

---