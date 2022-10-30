
# Online election system

An online simulation of real world elections.




## API Reference

#### Add new user

```http
  POST localhost:8080/api/add-user
```

```curl
curl --location --request POST 'localhost:8080/api/add-user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "password": "Lorem",
    "phone_number": "Lorem",
    "personal_info": {
        "name": "Lorem",
        "father_name": "Lorem",
        "dob": "2016-04-08",
        "age": "Lorem",
        "document_type": "Lorem",
        "address": {
            "street": "Lorem",
            "city": "Lorem",
            "state": "Lorem",
            "zip_code": "Lorem",
            "country": "Lorem"
        }
    },
    "uploaded_docs": {
        "document_type": "Lorem",
        "document_identification_no": "Lorem",
        "document_path": "Lorem"
    }
}'
```

#### Verify user

```http
  POST localhost:8080/api/verify-user
```
```curl
curl --location --request POST 'localhost:8080/api/verify-user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "_id": "0e1efc01f9961f8af5cee4f6",
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "is_verified": true
}'
```

#### Update user

```http
  POST localhost:8080/api/update-user
```
```curl
curl --location --request POST 'localhost:8080/api/update-user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "_id": "0e1efc01f9961f8af5cee4f6",
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "password": "Lorem",
    "phone_number": "Lorem",
    "personal_info": {
        "name": "Lorem",
        "father_name": "Lorem",
        "dob": "2016-04-08",
        "age": "Lorem",
        "voter_id": "Lorem",
        "document_type": "Lorem",
        "address": {
            "street": "Lorem",
            "city": "Lorem",
            "state": "Lorem",
            "zip_code": "Lorem",
            "country": "Lorem"
        }
    },
    "is_verified": true,
    "verified_by": {
        "_id": "6c6c2d84de0f09d5007bdef8",
        "name": "Lorem"
    },
    "uploaded_docs": {
        "document_type": "Lorem",
        "document_identification_no": "Lorem",
        "document_path": "Lorem"
    },
    "voted": [
        "ccebdc4611f903debe1b15ac"
    ]
}'
```

#### Search one user

```http
  POST localhost:8080/api/search-one-user
```
```curl
  curl --location --request POST 'localhost:8080/api/search-one-user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "_id": "0e1efc01f9961f8af5cee4f6",
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem"
}'
```
#### Search multiple users

```http
  POST localhost:8080/api/search-multiple-user
```
```curl
curl --location --request POST 'localhost:8080/api/search-multiple-user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "role": "Lorem",
    "city": "Lorem",
    "state": "Lorem",
    "zip_code": "Lorem",
    "country": "Lorem",
    "is_verified": true,
    "voted": [
        "ccebdc4611f903debe1b15ac"
    ]
}'
```

#### Delete user

```http
  POST localhost:8080/api/delete-user
```
```curl
curl --location --request POST 'localhost:8080/api/delete-user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "_id": "0e1efc01f9961f8af5cee4f6",
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem"
}'
```

#### Deactivate user

```http
  POST localhost:8080/api/deactivate-user
```
```curl
curl --location --request POST 'localhost:8080/api/deactivate-user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "_id": "0e1efc01f9961f8af5cee4f6",
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "is_verified": false
}'
```

### Add new election
```http
POST localhost:8080/api/add_election
```

```curl
curl --location --request POST 'localhost:8080/api/add_election' \
--header 'Content-Type: application/json' \
--data-raw '{
    "election_date": " ",
    "result_date": " ",
    "election_staus": " ",
    "result": " ",
    "location": " "
}'
```

### Add new candidate
```http
POST localhost:8080/api/add_candidate
```

```curl
curl --location --request POST 'localhost:8080/api/add_candidate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "candidate": [
        {
            "election_id": "",
            "candidate_name": " ",
            "vote_count": " ",
            "vote_sign": " ",
            "nomination_status": " ",
            "nomination_verified_by": " "
        }
    ]
}'
```

### Verify candidate
```http
POST localhost:8080/api/verify-candidate
```

```curl
curl --location --request POST 'localhost:8080/api/verify-candidate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": " ",
    "nomination_status": " "
}'
```

### Update election
```http
POST localhost:8080/api/update_election/{id}
```

```curl
curl --location -g --request POST 'localhost:8080/api/update_election/{id}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "election_date": " ",
    "result_date": " ",
    "election_staus": " ",
    "location": " "
}'
```

### Search election
```http
POST localhost:8080/api/election_search
```

```curl
curl --location --request POST 'localhost:8080/api/election_search' \
--header 'Content-Type: application/json' \
--data-raw '{
    "_id": "0e1efc01f9961f8af5cee4f6",
    "election_date": " ",
    "result_date": " ",
    "election_staus": " ",
    "candidate_name": " ",
    "location": " "
}'
```

### Deactivate election
```http
DELETE localhost:8080/api/deactivate_election/{id}
```

```curl
curl --location -g --request DELETE 'localhost:8080/api/deactivate_election/{id}'
```

### Cast a vote
```http 
POST localhost:8080/cast-vote
```

```curl
curl --location --request POST 'localhost:8080/cast-vote' \
--header 'Content-Type: application/json' \
--data-raw '{
    "election_id": "",
    "candidate_id": ""
}'
```

### Search election result
```http
POST localhost:8080/election-result
```

```curl
curl --location --request POST 'localhost:8080/election-result' \
--header 'Content-Type: application/json' \
--data-raw '{
    "election_id": "",
    "location": "",
    "election_date": "",
    "result_date": "",
    "election_status": ""
}'
```

### Search candidate profile

```http
POST localhost:8080/candidates-profile
```

```curl
curl --location --request POST 'localhost:8080/candidates-profile' \
--header 'Content-Type: application/json' \
--data-raw '{
    "election_id": "",
    "location": "",
    "election_date": "",
    "election_status": ""
}'
```
## Authors

- [Mukesh Kumar Jangir](https://www.github.com/gic-mukesh)

