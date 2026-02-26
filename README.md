# SelectelTest Log Linter

Линтер для проверки лог-сообщений в Go приложениях. Совместим с `log/slog` и `go.uber.org/zap`.

## Правила проверки

Линтер проверяет 4 правила для лог-сообщений:

### 1. Сообщения должны начинаться со строчной буквы

❌ **Неправильно:**
```go
log.Info("Starting server on port 8080")
slog.Error("Failed to connect to database")
```

✅ **Правильно:**
```go
log.Info("starting server on port 8080")
slog.Error("failed to connect to database")
```

### 2. Сообщения должны быть только на английском языке

❌ **Неправильно:**
```go
slog.Info("запуск сервера")
logger.Error("fehler aufgetreten")
```

✅ **Правильно:**
```go
slog.Info("starting server")
logger.Error("error occurred")
```

### 3. Сообщения не должны содержать спецсимволы или эмодзи

❌ **Неправильно:**
```go
log.Info("server started!🚀")
log.Error("connection failed!!!")
log.Warn("warning: something went wrong")
```

✅ **Правильно:**
```go
log.Info("server started")
log.Error("connection failed")
log.Warn("something went wrong")
```

### 4. Сообщения не должны содержать чувствительные данные

❌ **Неправильно:**
```go
log.Info("user password: " + password)
log.Debug("api_key=" + apiKey)
log.Info("token: " + token)
```

✅ **Правильно:**
```go
log.Info("user authenticated successfully")
log.Debug("api request completed")
log.Info("token validated")
```

Список чувствительных ключевых слов:
- `password`, `passwd`, `pwd`
- `token`, `jwt`, `bearer`
- `api_key`, `apikey`, `api-key`
- `secret`, `private_key`, `private-key`
- `credit_card`, `card_number`, `cvv`
- `ssn`, `social_security`
- `authorization`

## Установка

### Способ 1: Standalone использование

```bash
go install github.com/VladimirGladky/SelectelTest/cmd/selectellinter@latest
selectellinter ./...
```

### Способ 2: Интеграция с golangci-lint (Module Plugin)

1. Создайте файл `.custom-gcl.yml` в корне проекта:

```yaml
version: v1.62.2
plugins:
  - module: 'github.com/VladimirGladky/SelectelTest'
    import: 'github.com/VladimirGladky/SelectelTest'
    version: v1.0.0
```

2. Создайте или обновите `.golangci.yml`:

```yaml
linters-settings:
  custom:
    loglinter:
      type: "module"
      description: "Checks log messages for established rules"

linters:
  enable:
    - loglinter
```

3. Соберите кастомный golangci-lint:

```bash
golangci-lint custom
```

4. Используйте собранный бинарник:

```bash
./custom-gcl run ./...
```

### Способ 3: Локальное использование без публикации

```bash
# Клонируйте репозиторий
git clone https://github.com/yourusername/SelectelTest.git
cd SelectelTest

# Соберите бинарник
go build -o selectellinter ./cmd/selectellinter

# Запустите на вашем проекте
./selectellinter /path/to/your/project/...
```

## Использование

### Прямой запуск

```bash
selectellinter ./...
selectellinter ./pkg/...
selectellinter ./cmd/server
```

### С golangci-lint

```bash
golangci-lint run
```

### В CI/CD

```yaml
# GitHub Actions
- name: Run golangci-lint with custom linter
  run: |
    golangci-lint custom
    ./custom-gcl run ./...
```

## Поддерживаемые логеры

- `log/slog` (стандартная библиотека Go)
- `go.uber.org/zap`
- Любые кастомные обертки с методами: `Info`, `Error`, `Debug`, `Warn`, `Fatal`, `Panic`

## Разработка

### Запуск тестов

```bash
go test -v
```

### Структура проекта

```
SelectelTest/
├── analyzer.go           # Основной анализатор
├── analyzer_test.go      # Тесты анализатора
├── utils/
│   └── helpers.go        # Вспомогательные функции
├── cmd/
│   └── selectellinter/
│       └── main.go       # CLI интерфейс
├── plugin/
│   └── plugin.go         # Плагин для golangci-lint
├── testdata/             # Тестовые данные
│   └── src/
│       ├── a/            # Тесты для slog
│       ├── c/            # Тесты для zap
│       ├── slog_test/    # Unit тесты slog
│       └── zap_test/     # Unit тесты zap
├── .golangci.yml         # Конфигурация golangci-lint
├── .custom-gcl.yml       # Конфигурация custom plugin
└── README.md
```
