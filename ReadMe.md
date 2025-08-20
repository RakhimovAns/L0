# Order Service
_____
Демонстрационный сервис с Kafka, PostgreSQL, кешем.


## Задание

Необходимо разработать демонстрационный сервис с простейшим интерфейсом, отображающий данные о заказе.
____
## Пример видео
Пример можно посмотреть по ссылке ==> https://disk.yandex.ru/i/0O3ulKq7Jds3Cg

____
## Архитектура
Go — бизнес-логика и API.

Kafka — брокер сообщений.

PostgreSQL — основная база данных.

Redis — кеширование данных.

DI (Dependency Injection) — используется для управления зависимостями между компонентами (репозитории, логгер, кэш и др.).

____

## Запуск и проверка
1. **Скопировать Репозиторий**
```bash
  git clone https://github.com/RakhimovAns/L0.git
```
2. **Запуск проекта:**
```bash
  docker-compose up -d
```
3. **Добавление в кафку**

Чтоб не писать через терминал, можно перейти по пути ./cmd/producer/first_producer/first_main.go и запустить его
либо воспользоваться командой 
```bash
  go run ./cmd/producer/first_producer/first_main.go
 ```
4. **Доступ к АПИ**

АПИ будет доступен по ссылке ==> http://localhost:9000

_____
## Использованные технологии
- **Go**

- **Kafka**

- **Redis**

- **PostgreSQL**

- **Docker**
_____ 
## Использованные библиотеки
- Для подключение к бд и для транзакции использовал свою либу https://github.com/RakhimovAns/txmananger 
- Для логов свою либу https://github.com/RakhimovAns/logger
- Для легкого управления проектами https://github.com/RakhimovAns/wrapper
- Web фреймворк https://github.com/gofiber/fiber