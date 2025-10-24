# Task Manager API
A RESTful API for managing tasks built with Go (no frameworks) and PostgreSQL.

## Learning Objective
Understanding Go's HTTP server, database connections, and error handling without relying on frameworks.

## 🛠️ Tech Stack
- **Language:** Go 1.22+
- **Database:** PostgreSQL
- **Driver:** `github.com/lib/pq`
- **Testing:** Postman

## 📋 Features
- ✔️ Create new tasks
- ✔️ Get all tasks with pagination
- ✔️ Get single task by ID
- ✔️ Update task status
- ✔️ Delete tasks

## Step 1:
1. Build and start the server.
2. Connect to the database.
3. Complete the `createTask` handler for the `POST` method.

### Lessons learned:
1. Don't make the start more complex by immediately restructuring folders.
2. First, work in a single file.
3. Complete one handler first (`createTask`); then you can start structuring.

## Step 2:
1. Structuring the `CreateTask` first.
2. Yes, it's not easy. See what has been done.
3. For `GetAllTasks` we need two things:
    1. fetching data from the database 
    2. need an HTTP route that returns JSON

## Step 3:
1. Complete `GetTask` by ID (Wildcard route).
2. Similar to `GetAllTasks` but need to get the ID first.

## Step 4:
1. Complete `DeleteTask`.
2. Now I feel it's easy.

## Step 5:
1. Complete `UpdateTask`.
2. After updating, the order changed. Add `ORDER BY ... ASC` in `SELECT` statement of `GetAllTasks`.

## Step 6:
1. Intended to implement pagination, but ended up structuring HTTP requests instead.
2. 🥵

## Step 7:
1. Implement pagination.
2. 🎯


