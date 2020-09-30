# todo-Api

Api for CRUD operations for a todo application.

#Setup

##Import Database

``pg_dump -U postgres todo_app < pgDB.pgsql

Password:ishank111``

##Run API

``go build

go run todoApp``

##Endpoints

``GET := /todos  Fetches the list of todos from the API

POST := /add  Adds a new todo with initial status todo. Put data in body

PUT := /update/id=1 Updates the status of the todo with the given id

DELETE := /delete/id=1 Deletes the todo of the given task``
