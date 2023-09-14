# bankservice


[![Bank Service CI/CD Workflow Status](https://github.com/shivshankarm/bankservice/actions/workflows/test.yml/badge.svg)](https://github.com/shivshankarm/bankservice/actions/workflows/test.yml) 

This project aims to build a complete backend service for an application called BankService. The service performs multiple transactions from different users with different accounts. 

1. Design DB schema for PostgreSQL and general SQL code using dbdiagram.io. Use Docker to run this entire application. Use TablePlus to connect to PostgreSQL Server to look at the raw data.
2. Database migration using golang - golang-migrate. Generate CRUD (Create, Read, Update and Delete) Golang code from SQL code using Sqlc. Unit tests for the Go code for ~95% code coverage.
3. Perform transactions that are ACID compliant using Golang. 
