# 01-Task-Management-API
Introducing an api for task managing in Golang using Go Fiber as framework.

## quick start
if you already installed Makefile, try this
```
make tidy
make run
```

for those who didn't
```
go mod tidy
go run ./cmd/server/
```

## requirement
- this repository required Go version atleast 1.18 >> [Go](https://go.dev/doc/install)

# api walkthrough
1. **Task Model:**
    - I've define a struct like this:
        - **`ID`**: Unique identifier for the task (auto-incremented by the database).
        - **`Title`**: Title of the task.
        - **`Description`**: Description of the task.
        - **`Status`**: Status of the task (To Do, In Progress, Done).
        - **`CreatedAt`**: Time when this task was created.
        - **`UpdatedAt`**: Time when this task was updated.
2. **Status Enumeration** 
    - I've define an enumeration for **`status`** like this:
        - **`0`**: To Do.
        - **`1`**: In Progress.
        - **`2`**: Done.
    - You'll need to use this on update status operation
    - If you're entered other value than the following, I'll returned an error of invalid status. So, make sure you entering it correctly.
3. **Routing** 
    - All api's routing operation:
        - Method: **`GET`**  **`{{baseURL}}/api/v1/tasks/:task_id?`**: leave task_id empty if ypu wanna list all tasks, specify task_id to get task by ID.
        - Method: **`POST`**  **`{{baseURL}}/api/v1/tasks/`**: create a task.
        - Method: **`PATCH`**  **`{{baseURL}}/api/v1/tasks/:task_id?`**: update task title and description.
        - Method: **`PATCH`**  **`{{baseURL}}/api/v1/tasks/:task_id?`**: update task status.
        - Method: **`DELETE`**  **`{{baseURL}}/api/v1/tasks/:task_id?`**: delete a task.