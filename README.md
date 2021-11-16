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

### Data base
В качестве бд используем `mongodb`, Для того, что бы все работало добавим в базу данных `admin` 3 коллекции
- Transport
- TransportGroup
- User
логика добавления группы нереализована (сделан только поиск по группе), поэтому нам нужно добавить в TransportGroup следующую запись:

`{"_id":{"$oid":"618f87a1df2aa74d1f2c50fa"},"name":"Hi","unitsids":[{"$oid":"619120ae64c51eab1ca3dbe2"}]}`

Далее для теста добавим Транспорт (таблица Transport):

`{"_id":{"$oid":"619120ae64c51eab1ca3dbe2"},"creatorname":"Никита","name":"Тачела11","description":"qwfqwf","createdat":"12/31/2021 23:59:59","type":"Хитчбек"}`

Добавим пользователя в группу (User)

`{"_id":{"$oid":"618ca3a53b1d0279b9415176"},"Email":"admin","Password":"admin","Permission":"admin","UserName":"Никита"}`
