# AVITO-TASK
## Описание
Данное приложение позволяет взаимодействовать с уже имеющейся базой данных.
* создавать новых юзеров и добавлять их в список 
* списывать средства со счета
* начислять средства на счет 
* переводить средства со счета на счет
* получать список юзеров
* получать список всех производимых операций
* получать баланс счета конкретного юзера 
(в каждой функции производится проверка наличия и возможность создания id и невозможность создания отрицательного баланса)

Стек: 
* Go 1.16
* Postgresql

## Запуск 
go run main.go checks.go handlers.go

## Создание БД (производилось в psql)

Список юзеров
name | type | value
:----|:----:|-----:|
id | integer| not null|
balance| integer | not null|
status | varchar(20)| not null|




