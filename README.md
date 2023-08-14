# Rest API
A RESTful API example written in go featuring secure user authentication.
## Endpoints
All endpoints begin with the prefix `/api`. 

### Unauthenticated Endpoints
`GET /register`
Returns an API key which can be used for authenticated endpoints.
##### Response
```json
{
    "apikey": "key=="
}
```
### Authenticated Endpoints
Authentication is expected to be an `API key` passed in by the `Authorization` http header.
#### Get a product
`GET /product/{id}`
Get a product via it's product ID.
##### Response
```json
{
    "ID": 1,
    "Name": "Book"
}
```
#### Create a product
`POST /product`
Create a product with a name and ID.
##### Response
```json
{
    "message": "Product created."
}
```
#### Delete a product
`DELETE /product/{id}`
Deletes a product via its ID.
##### Response
```json
{
    "message": "Product deleted."
}
```