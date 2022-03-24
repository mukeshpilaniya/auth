[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=mukeshpilaniya_auth&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=mukeshpilaniya_auth)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=mukeshpilaniya_auth&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=mukeshpilaniya_auth)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=mukeshpilaniya_auth&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=mukeshpilaniya_auth)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=mukeshpilaniya_auth&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=mukeshpilaniya_auth)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=mukeshpilaniya_auth&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=mukeshpilaniya_auth)

[MakeFile](https://raw.githubusercontent.com/mukeshpilaniya/auth/main/Makefile) 
- Start the application
  - make up
- Stop the application
  - make stop
- Build docker images form source code
  - make docker-build
  
### RestAPI call endpoints
```
1. Create User 
  Endpoint url:- http://localhost:8081/api/v1/create_user
  Request Type-POST method
  Example:
  Request:- http://localhost:8081/api/v1/create_user
  {
      "first_name": "mukesh",
      "last_name": "pilaniya",
      "email": "test@gmail.com",
      "password": "pilaniya"
  }
  Response:-
  {
      "id": "bf325346-ab3f-11ec-8284-86023e8131c0",
      "first_name": "mukesh",
      "last_name": "pilaniya",
      "email": "test@gmail.com",
      "password": "pilaniya",
      "is_verified": false,
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z"
  }
```
``` 
2. Generate Token
   Endpoint url:- http://localhost:8081/api/v1/generate_token
   Request Type-POST method
   Example:
   Request:- http://localhost:8081/api/v1/generate_token
   {   
      "email": "test@gmail.com",
      "password": "pilaniya"
   }
   Response:-
   {
      "message": "token generated",
      "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImI5NDBiMjhjLTczYTgtNDJjNS05ZjczLWYxODg4YmJlNjlkYSIsInVzZXJfaWQiOiJiZjMyNTM0Ni1hYjNmLTExZWMtODI4NC04NjAyM2U4MTMxYzAiLCJpc3MiOiJwaWxhbml5YS5hdXRoLnNlcnZpY2UiLCJpc3N1ZWRfYXQiOjE2NDgxMDU0NzgsImV4cGlyZWRfYXQiOjE2NDgxMDU1OTgsImNsYWltcyI6e319.4k1PNH3Mxthw5NEF4RDepMLlYMnkOpYfOIHw49b9ocE",
      "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjA2Nzk5MzUwLTAzZTQtNDhhOS05NmNlLTJiOWYyYTU4NzViOCIsInVzZXJfaWQiOiJiZjMyNTM0Ni1hYjNmLTExZWMtODI4NC04NjAyM2U4MTMxYzAiLCJpc3MiOiJwaWxhbml5YS5hdXRoLnNlcnZpY2UiLCJpc3N1ZWRfYXQiOjE2NDgxMDU0NzgsImV4cGlyZWRfYXQiOjE2NDgxMDYwNzgsImNsYWltcyI6e319.1LowDhPuKm46nQXPZ7Wp2RihOfiFEUE9edr43vm5ogU"
   }
```
``` 
3. Get User by Id
   Endpoint url:- http://localhost:8081/api/v1/get_user
   Request Type-POST method
   Example:
   Request:- http://localhost:8081/api/v1/get_user
   Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImI5NDBiMjhjLTczYTgtNDJjNS05ZjczLWYxODg4YmJlNjlkYSIsInVzZXJfaWQiOiJiZjMyNTM0Ni1hYjNmLTExZWMtODI4NC04NjAyM2U4MTMxYzAiLCJpc3MiOiJwaWxhbml5YS5hdXRoLnNlcnZpY2UiLCJpc3N1ZWRfYXQiOjE2NDgxMDU0NzgsImV4cGlyZWRfYXQiOjE2NDgxMDU1OTgsImNsYWltcyI6e319.4k1PNH3Mxthw5NEF4RDepMLlYMnkOpYfOIHw49b9ocE
   {   
      "id": "bf325346-ab3f-11ec-8284-86023e8131c0"
   }
   Response:-
   {
      "id": "bf325346-ab3f-11ec-8284-86023e8131c0",
      "first_name": "mukesh",
      "last_name": "pilaniya",
      "email": "test@gmail.com",
      "password": "pilaniya",
      "is_verified": false,
      "created_at": "2022-03-24T12:28:02.068075Z",
      "updated_at": "2022-03-24T12:28:02.068075Z"
   }
```
```
4. Generate Access Token Form Refresh Token
   Endpoint url:- http://localhost:8081/api/v1/generate_access_token
   Request Type-POST method
   Example:
   Request:- http://localhost:8081/api/v1/generate_access_token
   {
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjA2Nzk5MzUwLTAzZTQtNDhhOS05NmNlLTJiOWYyYTU4NzViOCIsInVzZXJfaWQiOiJiZjMyNTM0Ni1hYjNmLTExZWMtODI4NC04NjAyM2U4MTMxYzAiLCJpc3MiOiJwaWxhbml5YS5hdXRoLnNlcnZpY2UiLCJpc3N1ZWRfYXQiOjE2NDgxMDU0NzgsImV4cGlyZWRfYXQiOjE2NDgxMDYwNzgsImNsYWltcyI6e319.1LowDhPuKm46nQXPZ7Wp2RihOfiFEUE9edr43vm5ogU"
   }
   Response:-
   {
	"message": "Access token generated",
	"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImY2MTYwZmI2LTY3ODgtNDI5Zi04ZjEwLWE4NTMyOTIxYjA4YiIsInVzZXJfaWQiOiJiZjMyNTM0Ni1hYjNmLTExZWMtODI4NC04NjAyM2U4MTMxYzAiLCJpc3MiOiJwaWxhbml5YS5hdXRoLnNlcnZpY2UiLCJpc3N1ZWRfYXQiOjE2NDgxMDU3OTEsImV4cGlyZWRfYXQiOjE2NDgxMDU5MTEsImNsYWltcyI6e319.E9GotSTOh6Gqn1Oupv_X4_a9IoTtF5pvyAjXKkC3_Ok"
    }
```