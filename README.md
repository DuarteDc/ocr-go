# OCR - Full Text PDFs

This a basic project to learn golang, that help to extract data of PDFs and search PDF with his content

## How to start

To start this project install dependencies

```
go mod tidy
```

Run project

```
go run cmd/api/main.go
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`PORT=8082`

`DB_HOST=localhost`

`DB_PORT=5432`

`DB_NAME=document_search`

`DB_USER=postgres`

`DB_PASSWORD=postgres`

## API Reference

#### Upload documents

```http
  POST /documents
```

| Parameter | Type   | Description       |
| :-------- | :----- | :---------------- |
| Document  | `File` | **Required file** |

#### Search Text

```http
  GET /search
```

| Parameter | Type   | Description                               |
| :-------- | :----- | :---------------------------------------- |
| ?q=       | string | **Required**. q of search text inside pdf |

#### Ask about data inside PDFs

```http
  GET /ask
```

| Parameter | Type   | Description                               |
| :-------- | :----- | :---------------------------------------- |
| question  | string | **Required**. q of search text inside pdf |
