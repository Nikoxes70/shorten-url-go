# Shorten URL Server

## Quickstart

### Create short url
#### URI

```text
http://localhost:8080/shorturl?type={id_type}
```

#### HTTP Method

POST

#### Required parametrs

| Field     | Type       | Description          |
| --------- | ---------- | -------------------- |
| int-int   | string     | The key represents the time interval to redirect the entire url to, which is the value |

#### Optional Parameters

| Field    | Type           | Description                              |
| -------- | -------------- | ---------------------------------------- |
| id type  | [alphanumeric/uuid] | type of id, default is alphanumeric |

#### Sample Request

```text
http://localhost:8080/shorturl?type=alphanumeric
```

```json
{
  "1-13": "https://ksp.co.il",
  "13-15":"https://walla.co.il",
  "15-0":"https://ynet.co.il"
}
```

#### Sample Response

```json
{
  "id":"99a628"
}
```


### Create short url
#### URI

```text
http://localhost:8080/{id}
```

#### HTTP Method

Get

#### Required parametrs

| Field    | Type           | Description                              |
| -------- | -------------- | ---------------------------------------- |
| id  | string | short url id |

#### Sample Request

```text
http://localhost:8080/99a628
```

#### Sample Response

```text
  Redirecting to the relevant page -> <!DOCTYPE html>
```
