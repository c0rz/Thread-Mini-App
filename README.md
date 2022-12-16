
# Thread Mini Apps with Go

A simple social media application according to the category you want to submit, which has 2 ADMIN and MEMBER levels.
## Installation

Tutorial Install Thread Mini Apps with Golang, don't forget to setting .env (change name file .env_example)

```bash
  git clone https://github.com/c0rz/Thread-Mini-App.git
  cd Thread-Mini-App
  go get
  go run main.go
```
    
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file, example in file .env_example

## API Reference

#### Auth Login

```http
  POST /api/auth/login
```

| Request | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required, Format Email**. Email Address |
| `password` | `string` | **Required**. Password |

#### Auth Register

```http
  POST /api/auth/register
```

| Request | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name` | `string` | **Required**. Full Name |
| `email` | `string` | **Required, Format Email**. Email Address |
| `password` | `string` | **Required**. Password |


#### Send Thread

```http
  POST /api/thread/send
```

| Request | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `title` | `string` | **Required**. Title Thread |
| `text` | `string` | **Required**. Body Thread |
| `category_id` | `string` | **Required**. Category ID |


#### Get All Thread

```http
  GET /api/thread
```

#### Update Thread

```http
  PUT /api/thread/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id` | `integer` | **Required**. ID Thread |


#### Delete Thread

```http
  DELETE /api/thread/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id` | `integer` | **Required**. ID Thread |


#### View Thread

```http
  GET /api/thread/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id` | `integer` | **Required**. ID Thread |

#### Insert/Add Comment in Thread

```http
  POST /api/thread/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id_post` | `integer` | **Required**. ID Thread |
| `comment` | `text` | **Required**. Comment |


#### Admin - Add Category

```http
  POST /admin/category/add
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name` | `string` | **Required**. Name Category |

#### Admin - DELETE Category

```http
  DELETE /admin/category/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id` | `integer` | **Required**. ID Category |

## Tech Stack

**Database:** PostgreSQL

**Server:** Golang, Gin, sql-migrate

