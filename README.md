

## Сетевой сервис сбора данных о состоянии систем связи сотового провайдера.

>Сервис проводит сбор, предварительнцю обработку, и выдачу в графическом виде информацию о состоянии систем:

- SMS
- MMS
- Voice Cal
- Email
- Billing
---
- Так-же, собирает и обрабатаывает информацию о степени загруженности службы поддержки,
  количестве текущих заявок, и времнеи ожидания решения проблем.

---

## Запуск:

- Сервис стартует по адресу 127.0.0.1:8888/api
- Для запуска используйте `go run cmd/main.go`
---
- Симулятор провайдера 127.0.0.1:8383
- Для запуска `go run simulator/skillbox-diploma/main.go`
