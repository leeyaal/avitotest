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

Список юзеров (все юзеры уникальны. Если юзер имеет статус "new", значит операций по данному счету не производилось и баланс "0" фактически не имеет нулевого значения)
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

![](https://github.com/leeyaal/avitotest/blob/main/images/newuser(unsuccess).jpg)

#### Successful 

![](https://github.com/leeyaal/avitotest/blob/main/images/newuser(success).jpg)

### Income (зачисление средств)

#### Unsuccessful 

![](https://github.com/leeyaal/avitotest/blob/main/images/income(unsuccess).jpg)

#### Successful 

![](https://github.com/leeyaal/avitotest/blob/main/images/income(success).jpg)

### Outcome (списание средств)

#### Unsuccessful 

![](https://github.com/leeyaal/avitotest/blob/main/images/outcome(unsuccess).jpg)

#### Successful 

![](https://github.com/leeyaal/avitotest/blob/main/images/outcome(success).jpg)

### Transit (перевод средств от юзера к юзеру)

#### Unsuccessful 

![](https://github.com/leeyaal/avitotest/blob/main/images/transfer(unsuccess).jpg)

#### Successful 

![](https://github.com/leeyaal/avitotest/blob/main/images/transfer(success).jpg)

### Check (получение списка всех юзеров)

#### Successful 

![](https://github.com/leeyaal/avitotest/blob/main/images/check(success).jpg)

### Infobalance (получение баланса конкретного юзера)

#### Successful 

![](https://github.com/leeyaal/avitotest/blob/main/images/infobalance.jpg)
![](https://github.com/leeyaal/avitotest/blob/main/images/infobalance(success).jpg)

### History (получение списка истории операций)

#### Successful 

![](https://github.com/leeyaal/avitotest/blob/main/images/history.jpg)
![](https://github.com/leeyaal/avitotest/blob/main/images/history(success).jpg)

