# Algebra Tutor

This is an app that uses ML to assist with a user learning Algebra.

## To run

`docker compose up --build`

## To tear down

`docker compose down -v`

# Modules

## ml-service

Handles embeddings and inference-related functionality.

- Generates embeddings
- Runs model inference
- Exposes ML capabilities to the main application

## db-migration

Responsible for database schema management and seed data.

- Creates and updates database schema
- Runs migrations in order
- Seeds initial/reference data

## web-app

The main application containing both backend and frontend rendering logic.

This is a modular monolith using:

- Clean Architecture principles
- Go templates for server-side rendering
- HTMX for partial interactivity

The app is broken into the following components:

### domain

Defines the core business model of the system.

- Entities (e.g. Course, Unit, Question, InteractiveElement)
- Core interfaces
- Business rules (no framework or transport logic)

### service

Contains business logic and application workflows.

- Orchestrates domain behavior
- Applies rules and validations
- Coordinates repositories

### repo

Handles data persistence.

- Database queries and commands
- Implements repository interfaces defined in domain

### controller

Handles HTTP layer concerns.

Web controllers (HTML + HTMX):
- Serves full pages using Go templates
- Returns HTMX partial updates

API controllers (optional):
- JSON endpoints for external use

### view

Defines presentation models used for rendering UI.

- ViewModels for pages and components
- Shapes data for templates
- HTMX response models

### middleware

Handles cross-cutting concerns.

- Authentication / authorization
- Logging
- Request context injection
- Rate limiting (if needed)

### web
Frontend presentation layer (server-rendered UI + assets)

templates:
- Full page templates (quiz, course, unit)
- Layout templates

components:
- question components
- graph components
- interactive elements
- HTMX partials

static:
- JavaScript (HTMX, graph rendering, interactivity)
- CSS
- images and fonts

## Request Flow
Browser -> controller/web -> service -> domain -> repo -> view (ViewModel) -> web/templates + components -> HTML / HTMX response -> Browser renders UI

## Key Ideas

- Backend renders HTML (not JSON-first)
- HTMX handles partial updates
- JS is only for advanced UI (graphs, interactivity)
- Domain layer stays completely framework-free