# Soon to be translated
# Шаламкова Алиса БПИ-218
Задание сделано по этому ТЗ

https://docs.google.com/document/d/1M_AQEhjyq7GtiQJ4TWt3bQsYcq1x1pIvxVepMY8S0Ps/edit?usp=sharing

### Запуск программы:
```
docker-compose up --build
```

При логине пользователю в куки сохраняется jwt токен, 
при последующих запросах токен вычитывается из куки, а из токена расшифровываются id, username, role

Запросы в сервис order (кроме получения меню) может отправлять только авторизованный пользователь.
Все конечные точки и примеры запросов экспортированы в коллекцию Postman,
коллекцию можно импортировать с помощью файла kpo.postman_collection.json
