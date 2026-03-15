# Algebra Tutor

This is an app that uses AI to assist with a user learning Algebra.

### To run

`docker compose up --build`

### To tear down

`docker compose down -v`

## Modules

`ai-service`

This is the embedding & inference service.

`db-migration`

This module handles DB migrations and seeding the DB with data.

`web-app`

This is where the front-end and back-end code lives. The "main" application.

## Todos

[ ] Complete the data model for the main application
