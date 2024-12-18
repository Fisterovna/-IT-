calc_service
Описание
Простой веб-сервис для вычисления арифметических выражений.

Запуск сервиса
Установите Go: https://go.dev/dl/
Склонируйте проект из репозитория GitHub.
Перейдите в папку проекта и запустите сервер:
bash
Копировать код
go run ./cmd/calc_service/...
Сервер будет доступен по адресу: http://localhost:8080/api/v1/calculate.

Пример запроса с использованием PowerShell
Для тестирования работы вашего сервиса с помощью PowerShell используйте команду Invoke-RestMethod:

powershell
Копировать код
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/calculate" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"expression": "2+2*2"}'
Пояснение:

-Uri "http://localhost:8080/api/v1/calculate" — URL, к которому отправляется запрос.
-Method POST — метод запроса.
-Headers @{"Content-Type"="application/json"} — заголовок, указывающий, что тело запроса будет в формате JSON.
-Body '{"expression": "2+2*2"}' — тело запроса с арифметическим выражением для вычисления.
