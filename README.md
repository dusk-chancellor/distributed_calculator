# Распределенный калькулятор

Этот проект представляет собой выполнение финальной задачи в
Яндекс.Лицее ("распределенная система для выполнения вычислений").
Система состоит из 2х основных компонентов: *оркестратор* - "менеджер", 
управляющий и координирующий вычислениями; *агент* - "вычислитель", отвечающий 
за обработку вычислений.

## Содержание
### - [Установка](#установка)
### - [Мок-тест](#мок-тест)
### - [Структура проекта](#структура-проекта)
### - [Технологии](#технологии)
### - [Контакты](#контакты)

## Установка
### Старый добрый метод:
1. Откройте нужную директорию в терминале и клонируйте туда репозиторий:

`git clone https://github.com/dusk-chancellor/distributed_calculator.git`

Вы также можете скачать zip файл проекта и извлечь его в нужную директорию.

2. Перейдите в директорию проекта:

`cd distributed_calculator`

3. Установите обязательные зависимости:

`go mod tidy`

4. Запустите оркестратор с рабочей директории:

`go run ./cmd/orchestrator/main.go`

В терминале должно отобразиться 'running Orchestrator server at
localhost:8080' и 'running Orchestrator manager'

5. Запустите агента с рабочей директории:

`go run ./cmd/agent/main.go`

В терминале вы должны увидеть 'tcp listener started at localhost:5000'

6. Перейти на [localhost:8080](http://localhost:8080/) и начать проверять !

### Запуск докером:

1. Запустите [Docker Desktop](https://www.docker.com/products/docker-desktop/)
или же запустите docker daemon

2. Поднимите контейнеры:

`docker-compose up`

3. Ждите, много.

4. Отправляйтесь на [localhost:8080](http://localhost:8080/)
и начинайте проверять !

P.S ~~Автор не рекоммендует данный метод, т.к в его случае при запуске
контейнеров вылетели все фиксики из компьютера~~ 😃

## Мок-тест

> При проверке проекта с браузера рекомендуется постоянно отслеживать консоль
и состояние сети нажатием на F12. Так, Вам будет удобнее проверять что и как происходит.

- **После этапа запуска всех компонентов и перехода по [localhost:8080](http://localhost:8080/)
вас должно было перекинуть в [localhost:8080/auth](http://localhost:8080/auth)
(если так не произошло, то пожалуйста обновите страничку)**

![Auth Page](/presentation_imgs/auth_page.png)

- **Создайте для начала аккаунт: подберите запоминающееся имя и пароль, а
затем нажмите на кнопку "Sign up"**

![Sign Up](/presentation_imgs/sign_up.png)

- **Теперь можете заходить в свой аккаунт: введите те же данные и нажмите
на кнопку "Login". После успешного входа Вас должно перебросить на главную
страницу проекта**

![Login](/presentation_imgs/login.png)

> Созданный при входе токен протухнет через 3 минуты и вам надо будет снова
перезаходить в свой аккаунт. Это было сделано для того, чтобы проверить его работоспособность.

- **Введем пример '2+2*2', отправим его на вычисление кнопкой "Send" и обновим страничку**

![Status-stored](/presentation_imgs/status_stored.png)

- **Обновляем страницу пока пример не решится**

![Status-done](/presentation_imgs/status_done.png)

> Кнопка удалить работает. Зайдя в другой аккаунт, вы не обнаружите выражений других пользователей.

## Структура проекта

```
├── cmd - запуск компонентов
│ ├── agent - "вычислитель"
│ │ └── main.go
│ └── orchestrator - "менеджер"
│ │ └── main.go
├── database - база данных
│ └── storage.db    #sqlite3
├── frontend - внешняя оболочка
│ ├── auth - страница аутентификации
│ │ ├── index.html
│ │ ├── script.js
│ │ └── style.css
│ ├── main - главная страница
│ │ ├── index.html
│ │ ├── script.js
│ │ └── style.css
├── internal - внутренние пакеты
│ ├── grpc - коннекторы грпц связи
│ │ ├── agent - для агента (запуск tcp слушателя)
│ │ │ └── agent.go
│ │ └── orchestrator - для оркестратора (методы отправки запросов)
│ │ │ └── orchestrator.go
│ ├── http - для http-запросов
│ │ ├── handlers - REST API хендлеры
│ │ │ ├── auth - хендлеры авторизации
│ │ │ │ └── auth.go
│ │ │ ├── expression - хендлеры работы с выражениями
│ │ │ │ └── expression.go
│ ├── storage - методы для взаимодействия с базой данных
│ │ ├── expression_storage.go - методы для expressions table
│ │ ├── storage.go - инициализация и создание таблиц
│ │ └── user_storage.go - методы для users table
│ ├── utils - доп.инструменты для полноценной работы компонентов
│ │ ├── agent - инструменты агента
│ │ │ ├── calculation - произведение вычисления
│ │ │ │ ├── calculation.go
│ │ │ │ └── stack.go
│ │ │ ├── infix_to_postfix - превращение выражения в постфиксную запись
│ │ │ │ ├── infix_to_postfix.go
│ │ │ │ └── stack.go
│ │ │ ├── validator - валидация допустимости выражения
│ │ │ │ └── validator.go
│ │ ├── orchestrator - инструменты оркестратора
│ │ │ ├── jwts - генерация и валидация jwt-токенов
│ │ │ │ └── jwts.go
│ │ │ ├── manager - служба постоянной связи с базой данных и отправки выражения агенту
│ │ │ │ └── manager.go
├── proto - прото и сгенерированные файлы для grpc общения
│ ├── agent_grpc.pb.go
│ ├── agent.pb.go
│ └── agent.proto
├── docker-compose.yaml
├── Dockerfile.Agent
├── Dockerfile.Orchestrator
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

### REST API запросы

| POST /auth/signup/       | -> | RegisterUserHandler     | -> | RegisterUser          |

| POST /auth/login/        | -> | LoginUserHandler        | -> | LoginUser             |

| POST /expression/        | -> | CreateExpressionHandler | -> | InsertExpression      |

| GET /expression/         | -> | GetExpressionsHandler   | -> | SelectExpressionsByID |

| DELETE /expression/{id}/ | -> | DeleteExpressionHandler | -> | DeleteExpression      |

## Технологии

В этом проекте используются следующие технологии и инструменты:

<div style="display: flex; justify-content: space-around; flex-wrap: wrap;">

<a href="https://golang.org/" target="_blank">
 <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png" alt="Go" width="65">
</a>

<a href="https://grpc.io/" target="_blank">
 <img src="https://grpc.io/img/grpc_square_reverse_4x.png" alt="gRPC" width="65">
</a>

<a href="https://jwt.io/" target="_blank">
 <img src="https://jwt.io/img/logo.svg" alt="JWT" width="65">
</a>

<a href="https://developers.google.com/protocol-buffers" target="_blank">
 <img src="https://www.codespot.org/assets/cover/protocol-buffers.png" alt="Protocol Buffers" width="65">
</a>

<a href="https://www.sqlite.org/index.html" target="_blank">
 <img src="https://www.sqlite.org/images/sqlite370_banner.gif" alt="SQLite" width="65">
</a>

<a href="https://developer.mozilla.org/en-US/docs/Web/HTML" target="_blank">
 <img src="https://cdn.dribbble.com/users/66221/screenshots/1655593/media/63d9b0acd7e81cde54f291bdcf8a24df.png?resize=400x300&vertical=center" alt="HTML" width="65">
</a>

<a href="https://developer.mozilla.org/en-US/docs/Web/CSS" target="_blank">
 <img src="https://i.pinimg.com/736x/a9/dc/c7/a9dcc740cad3149598307b5de8bc10c3.jpg" alt="CSS" width="65">
</a>

<a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript" target="_blank">
 <img src="https://upload.wikimedia.org/wikipedia/commons/6/6a/JavaScript-logo.png" alt="JavaScript" width="65">
</a>

<a href="https://www.docker.com/" target="_blank">
 <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQTQtnXQkjxe8xcXLmq8OvpW5ugiLGN_Y8Jepra6Wt1AA&s" alt="Docker" width="65">
</a>

<a href="https://docs.docker.com/compose/" target="_blank">
 <img src="https://i0.wp.com/codeblog.dotsandbrackets.com/wp-content/uploads/2016/10/compose-logo.jpg?ssl=1" alt="Docker Compose" width="65">
</a>

</div>


#### Основной язык программирования (Backend)

- **Golang**: Основной язык программирования, используемый для разработки всех компонентов системы

#### Фреймворки и библиотеки

- **gRPC**: Используется для создания высокопроизводительного, открытого
и универсального RPC-фреймворка

- **JWT**: Используется для аутентификации и авторизации пользователей

- **Protocol Buffers**: Используется для определения структуры данных и сервисов,
обеспечивая эффективное и быстрое сериализование данных

#### База данных

- **SQLite3**: Используется в качестве легковесной базы данных для хранения информации
о пользователях и выражениях. <u>Запланировано</u>: миграция базы данных на PostgreSQL

#### Интерфейс пользователя (Frontend)

- **HTML, CSS, JavaScript**: Используются для создания веб-интерфейса, обеспечивая
интерактивность и удобство использования для пользователей. Веб-интерфейс
для этого проекта был максимально упрощен

#### Инструменты разработки (опционально)

- **Docker**: Используется для контейнеризации компонентов системы, обеспечивая
удобство развертывания и масштабирования

- **Docker Compose**: Используется для определения и запуска многоконтейнерных приложений
Docker, упрощая процесс развертывания и управления контейнерами.

Стоит учесть, что все выше перечисленные технологии не были применены в их полную силу

## Контакты

Если у вас вдруг что-то не работает, есть вопросы или предложения, то пожалуйста обратитесь ко [мне](https://t.me/duskchancellor)
