
# Sharing Vision Backend

A backend app for post article. Used Golang (Gin) as framework.


## Tech Stack

**Language**: Golang (Gin)

**ORM**: GORM

**Database**: MySQL (hosted via aiven)

## Endpoint Documentation##
Please refer to the postman collection
https://reinauth-900831.postman.co/workspace/Val's-Workspace~c2ef92cf-e373-4b79-87da-bee1a3cfc962/collection/47723930-5f019c6c-c2ff-4d69-afbb-21c74f6fe8c5?action=share&creator=47723930

## Prerequisites
To run this project locally, you will need:
* [Go](https://go.dev/dl/) installed (v1.18 or higher)
* A MySQL server running locally or a cloud MySQL connection string

## How to Run Locally

**Clone the repository:**
   ```bash
   git clone [YOUR_BACKEND_REPO_URL]
   cd [YOUR_BACKEND_FOLDER_NAME]
   ```
**Install Dependencies:**
    
    go mod tidy

**Configure Env:**

```bash
DB_URL=user:password@tcp(127.0.0.1:3306)/sharing_vision?charset=utf8mb4&parseTime=True&loc=Local
PORT=8080
```


**Build**

```bash
   go build -o main .
```

**Run**

```bash
   ./main
```
