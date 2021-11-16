# Веб приложение для управления ТС 
## Как установить?
клонируем репозиторий `git clone https://github.com/OnlyM1ss/Transport-Manager.git`
Есть 2 части 
- Server (go сервер `backend`)
- Web (angular вет `frontend`)
### Web
Запускаем visual studio code, открываем терминал в папке Web и вводим команду `npm install` 
после скачивания пакетов запускаем `ng serve`
Наш веб запущен по адресу: `http://localhost:4200/`
### Server
Запускаем GoLand (или другую IDE или текстовый редактор)
Открываем терминал и скачиваем все необходимые пакеты `go get github.com/OnlyM1ss/transport-manager/v2`, если не получается, то скачиваем необходимые пакеты по отдельности) 

запускаем проект 

Теперь веб должен общаться с сервером
