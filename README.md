# Простой веб-сервис для вычисления арифметических выражений

## Описание
Этот проект реализует веб-сервис, который вычисляет арифметические выражения, переданные пользователем через HTTP-запрос.

## Запуск сервиса

1. Установите [Go](https://go.dev/dl/).
2. Склонируйте проект с GitHub:
    ```bash
    git clone https://github.com/Fisterovna/-IT-.git
    ```
3. Перейдите в папку проекта и запустите сервер:
    ```bash
    go run ./cmd/calc_service/...
    ```
4. Сервис будет доступен по адресу: [http://localhost:8080/api/v1/calculate](http://localhost:8080/api/v1/calculate).

---

## Эндпоинты

### POST /api/v1/calculate
**Описание:** принимает JSON с математическим выражением.

**Пример успешного запроса:**
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "2+2*2"}'
Пример успешного ответа:

json
Копировать код
{
    "result": "6.000000"
}
Пример ошибки 422:

bash
Копировать код
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "invalid expression"}'
Ответ:

json
Копировать код
{
    "error": "invalid expression"
}
Пример ошибки 500:

bash
Копировать код
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "10 / $"}'
Ответ:

json
Копировать код
{
    "error": "Internal server error"
}
Тесты
Для проверки функциональности проекта написаны тесты.

1. Тесты для функции вычисления (Calc)
Файл: internal/calculate/calculator_test.go
Тесты покрывают:

Успешные вычисления (например, 2 + 2).
Ошибки, такие как деление на 0 или некорректное выражение.
2. Тесты для HTTP-эндпоинта
Файл: cmd/calc_service/main_test.go
Тесты проверяют:

Корректный запрос и успешный ответ.
Обработку ошибок (422 и 500).
3. Тесты для вспомогательных функций (mathutils)
Файл: pkg/mathutils/mathutils_test.go
Тесты покрывают базовые операции:

Сложение.
Вычитание.
Запуск тестов
Перейдите в корневую директорию проекта.
Выполните команду:
bash
Копировать код
go test ./...
Убедитесь, что все тесты прошли успешно.
