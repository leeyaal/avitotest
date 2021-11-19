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
(в каждой функции производится проверка наличия и возможность создания id и невозможность создания отрицательного баланса и существующего id)

Стек: 
* Go 1.16
* Postgresql

## Запуск 
go run main.go checks.go handlers.go

## Создание БД (производилось в psql)

Список юзеров (все юзеры уникальны. Если юзер имеет статус "new", значит операций по данному счету не производилось и баланс "0" фактически не имеет нулвого значения)
name | type | value
:----|:----:|-----:|
id | integer| not null|
balance| integer | not null|
status | varchar(20)| not null|


Список операций ("idcreator" - создатель операции, не указывается если производится списание или начисление, "idreciever" - юзер, который получает средства или у кого они списываются)
name | type | value
:----|:----:|-----:|
idcreator | integer| not null|
idreciever| integer | not null|
operation | varchar(20)| not null|
date | varchar(20)| not null|
sum | integer| not null|

## Примеры запросов 

### NewUser (создание юзера)

#### Unsuccessful 

![](newuser(unsuccess).png)

#### Successful 

![](newuser(success).png)

### Income (зачисление средств)

#### Unsuccessful 

![](income(unsuccess).png)

#### Successful 

![](income(success).png)

### Outcome (списание средств)

#### Unsuccessful 

![](outcome(unsuccess).png)

#### Successful 

![](outcome(success).png)

### Transit (перевод средств от юзера к юзеру)

#### Unsuccessful 

![](transfer(unsuccess).png)

#### Successful 

![](transfer(success).png)

### Check (получение списка всех юзеров)

#### Successful 

![](check(success).png)

### Infobalance (получение баланса конкретного юзера)

#### Successful 

![](infobalance.png)
![](infobalance(success).png)

### History (получение списка истории операций)

#### Successful 

![](history.png)
![](history(success).png)

