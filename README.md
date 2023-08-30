# Avito Backend Internship Task 2023
Сервис динамического сегментирования пользователей
<details>
  <summary>Содержание</summary>
  <ol>
    <li><a href="#установка-и-запуск">Установка и запуск</a></li>
    <li><a href="#реализовано">Реализовано</a></li>
    <li><a href="#возникшие-вопросы">Возникшие вопросы</a></li>
    <li><a href="#основные-сведения">Основные сведения</a></li>
  </ol>
</details>

## Установка и запуск

Клонировать проект.
Далее через `docker compose`:
- `docker-compose up --build testingavito`
- `docker-compose up testingavito`

Также подключить базу данных хранящуюся по адресу ./schema
Пример через миграции:  `migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up`
Данные для подключения могут быть другие. Поменять DB_PASS - в файле .env

## Реализовано

1. Метод создания пользователя.
2. Метод создания сегмента. Также при создании возможно указание процента пользователей автоматически попадающих в сегмент.
3. Метод удаления сегмента.
4. Метод добавления пользователя в несколько сегментов.
5. Метод удаления указанных сегментов у пользователя.
6. Метод получения списока сегментов в котором состоит пользователь.
7. Метод получения ссылки на csv файл с отчетом об удалении/добавлении сегментов пользователю.
8. Метод получения данные по таблицам базы данных.

## Возникшие вопросы

1. Нужно ли добавлять метод создания пользователей? - Да, так как без этого метода логики работы с апи не будет.
2. Нужно ли добавлять метод получения данных из таблиц БД через запрос? - Да, для удобства проверки результатов запросов.

## Основные сведения

Данная программа содержит реализацию сервиса динамического сегментирования пользователей. 

Оригинальное задание -> https://github.com/avito-tech/backend-trainee-assignment-2023 

Структура проекта выглядит следующим образом:
```markdown
├── build
|   └── config.yml                      
├── cmd
|   └── main.go                         
├── configs
|   └── config.yml                        
├── pkg
|   ├── handler
|   |   ├── command.go              
|   |   ├── handler.go                 
|   |   └── response.go            
|   |
|   ├── repository
|   |   ├── commanddb.go
|   |   ├── db.go   
|   |   └── repository.go             
|   |      
|   |                  
|   |                    
|   |                   
|   └── service   
|       ├── enty                
|       |      ├── history.go    
|       |      ├── segment.go
|       |      ├── server.go
|       |      └── user.go
|       |  
|       ├── mocks 
|       |       └── mock.go
|       |
|       ├── commservice.go
|       └── service.go
|
├── schema
|   ├── 000001_init.down.sql
|   └── 000001_init.up.sql
|
├── static
|   └── сюда загружается csv файл
|        
├── .env          
├── .gitignore
├── docker-compose.yml
├── Dockerfile                  
├── go.mod
├── Makefile
└── README.md
```

Для взаимодействия с сервером есть 8 способов:
1. (POST) /main/adduser -- создать пользователя (обязательно). Принимает JSON в качестве тела запроса;
2. (POST) /main/addsegment -- добавить сегмент. Принимает JSON в качестве тела запроса;
3. (DELETE) /main/delsegment -- удалить сегмент. Принимает JSON в качестве тела запроса;
4. (POST) /main/insertsegmentuser -- добавить пользователя в несколько сегментов. Принимает JSON в качестве тела запроса;
5. (DELETE) /main/delsegmentuser -- удалить указанные сегменты у пользователя. Принимает JSON в качестве тела запроса;
6. (POST) /main/getusersegment -- получить список сегментов в котором состоит пользователь. Принимает JSON в качестве тела запроса;
7. (GET) /main/history -- получить ссылку на csv файл с отчетом об удалении/добавлении сегментов пользователю. Принимает JSON в качестве тела запроса;
8. (GET) /main/showdatadb -- получить данные по таблицам базы данных. Принимает JSON в качестве тела запроса.

В качестве ответа возвращается JSON с данными.

Сами структуры запросов и ответов в этих файлах:

Примеры запросов и ответов:
1. Создать пользователя (корректный запрос) http://localhost:8080/main/adduser POST:
  - INPUT
  ```json
  {
    "user_number":1
  }
  ```
  OUTPUT
    ```json
  {
      "Message": "1" //возвращает user_number созданного пользователя
  }
    ```

2. Создать пользователя (некорректный запрос):
 INPUT:
    ```json
  {
      "user_number":-4
  }
    ```
  OUTPUT:
    ```json
  {
      "Message": "User number must > 0"
  }
    ```

3. Создать сегмент (коректный запрос) http://localhost:8080/main/addsegment POST:
 INPUT:
    ```json
  {
      "name":"AVITO_VOICE_MESSAGES"
  }
    ```
  OUTPUT:
    ```json
  {
      "Message": "AVITO_VOICE_MESSAGES" //возвращает slug_name созданного сегмента
  }
    ```

4. Создать сегмент с использованием процентов - количество пользователей автоматически попадающих в сегмент (корректный запрос) .
  http://localhost:8080/main/addsegment POST
- когда указываем процент пользователей подразумевается что в таблице Users существует хотя бы один пользователь, если это не так - получите ошибку
 - INPUT:
    ```json
  {
      "name":"AVITO_PERFORMANCE_VAS",
      "percent":50
  }
    ```
  - OUTPUT
    ```json
  {
      "Message": "AVITO_PERFORMANCE_VAS" //возвращает slug_name созданного сегмента
  }
    ```

5. Создать сегмент (некорректный запрос) http://localhost:8080/main/addsegment POST:
 INPUT:
    ```json
  {
      "name":"AVITO_PERFORMANCE_VAS"
  }
    ```
  OUTPUT:
    ```json
  {
      "message": "pq: duplicate key value violates unique constraint \"slug_name_unique\"" //нельзя дублировать значения
  }
    ```

6. Создать сегмент с использованием процентов - количество пользователей автоматически попадающих в сегмент (некорректный запрос):
  http://localhost:8080/main/addsegment POST
 INPUT:
    ```json
  {
     "name":"AVITO_DISCOUNT_30",
     "percent": -30 
  }
    ```
  OUTPUT:
    ```json
  {
      "Message": "Fix Percent ,pls" // проценты не могут быть меньше 0, 
  }
    ```

7. Удалить сегмент (корректный запрос) http://localhost:8080/main/delsegment DELETE:
- при удалении сегмента - сегменты также удаляются у тех пользователей, которые были в него включены, следовательно обновляется таблица History, в которой добаляется запись о том когда был удален тот или иной пользователь.
 INPUT:
    ```json
  {
      "name":"AVITO_PERFORMANCE_VAS"
  }
    ```
  OUTPUT:
    ```json
  {
      "Status": "OK"
  }
    ```

8. Удалить сегмент (некорректный запрос) http://localhost:8080/main/delsegment DELETE:
 INPUT:
    ```json
  {
      "name":"AVITO_PERFORMANCE_VAS1"
  }
    ```
  OUTPUT:
    ```json
  {
      "Status": "No such data"
  }
    ```

9. Добавить пользователя в список сегментов (корректный запрос) http://localhost:8080/main/insertsegmentuser POST:
- при добавлении пользователя в сегмент, данные также обновляется таблица History, в которой добаляется запись о том когда был добавлен тот или иной пользователь.
 INPUT:
    ```json
  {
      "name":["AVITO_VOICE_MESSAGES","AVITO_DISCOUNT_50","AVITO_blockchain"],
      "UserID":1
  }
    ```
  OUTPUT:
    ```json
  {
      "id": [
        3,
        4,
        5
    ] //id созданных записей в таблице Usersslug
  }
    ```

10. Добавить пользователя в список сегментов (некорректный запрос) http://localhost:8080/main/insertsegmentuser POST:
 INPUT:
    ```json
  {
      "name":["AVITO_VOICE_MESSAGES","AVITO_DISCOUNT_51","AVITO_DISCOUNT_50"], //2 элемент списка не существует в таблице slugs
      "UserID":6
  }
    ```
  OUTPUT:
    ```json
  {
      "message": "pq: insert or update on table \"usersslug\" violates foreign key constraint \"usersslug_name_slug_fkey\""
      // выполнит запрос до этого элемента и вернет ошибку, т.е 1 элемент списка добавиться, а второй уже нет и вернется ошибка
  }
    ```

11. Удалить сегменты у пользователя (корректный запрос) http://localhost:8080/main/delsegmentuser DELETE:
- при удалении пользователя из сегмента, данные также обновляется таблица History, в которой добаляется запись о том когда был удален тот или иной пользователь.
 INPUT:
    ```json
  {
      "name":["AVITO_VOICE_MESSAGES","AVITO_DISCOUNT_50"],
      "UserID":5
  }
    ```
  OUTPUT:
    ```json
  {
      "Status": "OK"
  }
    ```

12. Удалить сегменты у пользователя (некорректный запрос) http://localhost:8080/main/delsegmentuser DELETE:
 INPUT:
    ```json
  {
      "name":["AVITO_VOICE_MESSAGES","AVITO_DISCOUNT_50"],
      "UserID":25 
  }
    ```
  OUTPUT:
    ```json
  {
      "Status": "no such data" // пользователя с таким UserID - нет
      // в этом примере в отличии от предыдущего если указан элемент списка который не существует в Usersslug - просто пропускается
  }
    ```

13. Получить список сегментов в которых состоит пользователь (корректный запрос) http://localhost:8080/main/getusersegment POST:
 INPUT:
    ```json
  {
      "user_number":1
  }
    ```
  OUTPUT:
    ```json
  {
      "Names": [
        {
            "name": "AVITO_VOICE_MESSAGES"
        },
        {
            "name": "AVITO_DISCOUNT_50"
        },
        {
            "name": "AVITO_blockchain"
        }
    ]
  }
    ```

14. Получить список сегментов в которых состоит пользователь (некорректный запрос) http://localhost:8080/main/getusersegment POST:
 INPUT:
    ```json
  {
      "user_number":25
  }
    ```
  OUTPUT:
    ```json
  {
      "Message": "has no segments or non-existent user entered"
  }
    ```

15. Получить ссылку на csv файл с отчетом об удалении/добавлении сегментов пользователю (корректный запрос):
  http://localhost:8080/main/history GET
 INPUT:
    ```json
  {
      "id":5,
      "start":"2022/06",
      "end":"2023/09"
  }
    ```
  OUTPUT:
    ```json
  {
      http://localhost:8080/static/otchet.csv
  }
    ```
    Содержимое файла:
      NAME | MODE | DATE
      AVITO_VOICE_MESSAGES,ADD,2023-08-30
      AVITO_DISCOUNT_50,ADD,2023-08-30
      AVITO_VOICE_MESSAGES,DELETE,2023-08-30
      AVITO_DISCOUNT_50,DELETE,2023-08-30

16. Получить ссылку на csv файл с отчетом об удалении/добавлении сегментов пользователю (некорректный запрос):
  http://localhost:8080/main/history GET
 INPUT:
    ```json
  {
      "id":25,
      "start":"2022/06",
      "end":"2023/09"
  }
    ```
  OUTPUT:
    ```json
  {
      http://localhost:8080/static/otchet.csv // при некоректном запросе будет таке выдаваться сслыка, но файл будет пустым
      // вводим ссылку в браузере - загружается файл
  }
    ```

17. Получить данные по таблицам базы данных (корректный запрос) http://localhost:8080/main/showdatadb GET:
- сделано для удобства проверки работы сервиса с базой данных, возможные данные для запроса: Users - таблица пользователей, Slugs - таблица сегментов, UsersSlug - таблица сегментов в которых состоит пользователь, History - таблица добавления/удаления сегментов у пользователя.
 INPUT:
    ```json
  {
      "name_table":"usersslug"
  }
    ```
  OUTPUT:
    ```json
  {
      {"Status":[{"ID":3,"User_id":1,"Name":"AVITO_VOICE_MESSAGES"},{"ID":4,"User_id":1,"Name":"AVITO_DISCOUNT_50"},{"ID":5,"User_id":1,"Name":"AVITO_blockchain"},{"ID":6,"User_id":2,"Name":"AVITO_VOICE_MESSAGES"},{"ID":7,"User_id":2,"Name":"AVITO_DISCOUNT_50"},{"ID":14,"User_id":3,"Name":"AVITO_VOICE_MESSAGES"},{"ID":15,"User_id":3,"Name":"AVITO_DISCOUNT_50"},{"ID":22,"User_id":6,"Name":"AVITO_VOICE_MESSAGES"}]}
  }
    ```

18. Получить данные по таблицам базы данных (некорректный запрос) http://localhost:8080/main/showdatadb GET:
 INPUT:
    ```json
  {
      "name_table":"error"
  }
    ```
  OUTPUT:
    ```json
  {
      {"Status":"Nil result"} // так как таблицы error не существует
  }
    ```
