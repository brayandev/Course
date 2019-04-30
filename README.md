# Course

Service to manipulate course.

---

## Pre-requisites

`docker version 17+` See how to download and install in [Docker site.](https://docs.docker.com/install/linux/docker-ce/ubuntu/)

`golang version 1.11+`  See how to download and install in [Golang site.](https://golang.org/doc/install)

## API

* [Error handling](#error-handling)
* [Resources](#resources)
  * [Create course](#post-course)

## Error handling

if something went wrong on request, the application should return http code

### Error codes

| description | http status code |
| :-----------| :----------------|
| unknown error | `500`

## Body

## Resources

### `POST` /course

Create one course

#### `POST` parameters description

| parameter | type | description |
| :---------| :----| :-----------|
| courseID | string | A course ID |

#### `POST` body response

```json
{
    "courseId": "ff2736ed-8275-433c-8fb1-7d90935ea014"
}