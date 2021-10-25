![The stonehenge](https://static.turbosquid.com/Preview/2019/02/13__15_50_21/Lowpoly_Stonehenge_05.jpg720E3DA3-2A5D-4E28-A953-C096C27DB0D1Large.jpg) 

# Stonehenge
Stonehenge is a go app for you to send and receive money to and from your friends. It is easy to use and has a public and simple API that you can run directly on your computer through Docker

# Dependencies
This application uses the following packages:
* [UUID](github.com/google/uuid) - to create unique and safe entity identifiers
* [CHI](github.com/go-chi/chi/v5) - to create separate and concise routes for you
* [Firebase Firestore](https://pkg.go.dev/firebase.google.com/go) - to store your stuff in a safe place

# Usage
For you to run the API on your computer, you've got to have Docker installed. On a terminal, use the command below
```bash
docker pull eakeur/stonehenge
```
After the image is completely pulled to your computer, create a container on it with this command (you can change the container name or the output port):
```bash
docker run -d --name eakeur-stonehenge -p 3000:8080 eakeur/stonehenge:latest
```
That's all! You can access localhost:[THE_PORT_YOU_CHOSE] and have access to the API

# Endpoints
This section has useful and relevant information on how to consume the API. Please read everything below, so that you don`t make any mistake when making transfers










## Authentication
```
PATH:   /login
METHOD:   POST
BODY: 
```
Request body:
```json
{ "cpf": "string", "secret": "string"}
```
For you to authenticate to the API, you have to first access this endpoint with a json body containing your cpf and your password, like the above example. If you do not have an account yet, please refer to [this endpoint](#post-accounts). The response is empty, but has a Set-Cookie and an Authorization header member with the JWT token of your session









## Accounts

### Entity definition
This section of the app has several endpoints. Many of them use the same response structure, which is defined below and is named Account
```json
{ "cpf": "string", "secret": "string", "name": "string", "balance": "int", "created_at": "string", "id": "string"}
```
**The balance is always informed in cents**


### POST Accounts
```
PATH:   /accounts
METHOD:   POST
```
Request body:
```json
{ "cpf": "string", "secret": "string", "name": "string"}
```
This endpoint creates an account with the data in the request body and authenticates it. The expected body scheme is mentioned above in the BODY property.
The cpf must be a string with 11 numbers. After created, the user receives R$ 500.00 as a starter budget. The response is empty, but has a Set-Cookie and an Authorization header member with the JWT token of your session

### GET Accounts
```
PATH:   /accounts
METHOD:   GET
```
This endpoint returns a list of all accounts registered, with the account entity type. For safety of our users, the balance, the secret and the cpf properties are 0, "" and "", respectively.

### GET Account by ID
```
PATH:   /accounts/{accountId}
METHOD:   GET
```
This endpoint returns the account entity that corresponds to the accountId passed as parameter. **You can only access the account that is logged in at the moment**

### GET Balance by ID
```
PATH:   /accounts/{accountId}/balance 
METHOD:   GET 
```
This endpoint returns the account's balance that corresponds to the accountId passed as parameter. **You can only access the balance of the account that is logged in at the moment. The balance is sent in cents**










## Transfers
### POST transfers
```
PATH:   /transfers
METHOD:   POST
```
Request body:
```json
{ "destination_account_id": "string", "amount": "int" }
```
This endpoint tells the server that a transfer must be made. The request body json must be like the one specified in the body property above
**The amount of money must always be sent in cents**, otherwise the request will fail or consider the wrong amount of money. Please be sure that the destination id is correct, so that you don't make the transfer to the wrong account, as it can not be undone

### GET transfers
```
PATH:   /transfers
METHOD:   GET
QUERY: toMe (true or false)
```
This endpoint returns a list of all transfers registered, with the transfers entity type. The toMe parameter is a modifier that indicates whether you want to fetch the transfers made to you (true), or by you (false). The default is false. The response body should be an array of this model (remember that the amount is always in cents): 
```json
    { "id": "string", "origin_account_id": "string", "destination_account_id": "string", "amount": "int", "created_at": "string"}
```
