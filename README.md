# Простой веб-сервис для вычисления арифметических выражений

## Описание
Этот проект реализует веб-сервис, который вычисляет арифметические выражения, переданные пользователем через HTTP-запрос. Сервис поддерживает базовые арифметические операции, такие как сложение, вычитание, умножение и деление.

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
**Описание:** принимает JSON с математическим выражением, вычисляет его и возвращает результат.

**Пример успешного запроса:**
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "2+2*2"}'
Пример успешного ответа:
{
    "result": "6.000000"
}
Пример ошибки 422 (Некорректное выражение):
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "invalid expression"}'
Ответ:

json
Копировать код
{
    "error": "invalid expression"
}
Пример ошибки 500 (Внутренняя ошибка сервера):
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "10 / $"}'
Ответ:
{
    "error": "Internal server error"
}
Тесты
Для проверки функциональности проекта написаны тесты. Они покрывают:

Функцию вычисления (Calc):

Успешные вычисления (например, 2 + 2).
Ошибки, такие как деление на 0 или некорректное выражение.
HTTP-эндпоинт:

Корректный запрос и успешный ответ.
Обработку ошибок (422 и 500).
Вспомогательные функции (mathutils):

Сложение.
Вычитание.
Запуск тестов
Перейдите в корневую директорию проекта.
Выполните команду:
go test ./...
Убедитесь, что все тесты прошли успешно.

Структура проекта
Проект организован по следующим папкам:

cmd/ — точка входа в приложение.
internal/ — бизнес-логика приложения.
pkg/ — утилиты и вспомогательные пакеты.
scripts/ — автоматизационные скрипты (например, сборка, деплой).
