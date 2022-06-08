# Tech Stack
- Go v1.17.7
- Gin HTTP Framework
- GORM for SQL ORM
- PostgreSQL as DBMS
- Testify for testing tools

# Project Structure
![alt text](https://miro.medium.com/max/1400/1*phecRia6It8AnwlFjhjx2w.jpeg)
In this project, im using Clean Architecture principles with Domain Design Driver (DDD) that consist 4 layer:
1. domain/entity/model
2. repository
3. usecase/service
4. handler/controller/delivery


# Installation
```
1. Clone this repository: git clone https://github.com/yerobalg/delos-farm.git
2. Enter to directory   : cd delos-farm-backend
3. Copy or rename the .env.example file, enter your credential
3. Run postgres db      : docker compose up -d
4. Run test             : go test -v ./...
5. Run go app           : go run ./app/http/
6. If you wanna shut down db: docker compose down 
```

**Note:**: to avoid error, do not run the postgres DB in port 5432 

# API Documentation
📃: https://documenter.getpostman.com/view/14494329/Uz5JHFBj
