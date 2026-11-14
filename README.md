# Pastebin — Text Sharing Backend

**A personal backend development project exploring REST APIs, authentication, and persistent storage.**

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![REST API](https://img.shields.io/badge/REST_API-005571?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Work_in_Progress-orange?style=for-the-badge)

---

## Overview

Pastebin is my first personal backend pet project, inspired by text-sharing services such as Pastebin.

The goal is to build a service that allows users to submit text, store it persistently, and retrieve it using a unique identifier.

Beyond the application itself, this project serves as a practical introduction to backend architecture, HTTP APIs, authentication, databases, and object storage.

## Project Goals

- Design a RESTful API for managing text snippets.
- Support paste creation and retrieval without requiring an account.
- Explore authentication and user management.
- Separate application responsibilities into logical components.
- Integrate database storage for paste metadata.
- Explore S3-compatible storage for paste contents.

## Architecture

The repository is organized into two main directories:

```text
Pastebin/
├── auth/          # Authentication-related component
├── paste/         # Paste-related component
├── .gitignore
└── README.md
```

The intended storage model separates the actual text content from its associated metadata.

```text
Client
  |
  v
HTTP API
  |
  +-- Authentication component
  |
  +-- Paste component
        |
        +-- Metadata database
        |
        +-- S3-compatible object storage
```

*The diagram represents the intended system architecture, not a verified description of the current implementation.*

## Planned API

| Method | Endpoint | Purpose |
|---|---|---|
| `POST` | `/pastes` | Create a new paste |
| `GET` | `/pastes/{id}` | Retrieve an existing paste |

### Paste creation workflow

1. The client sends text to the backend.
2. The service processes the submitted content.
3. Paste content is stored in object storage.
4. Associated metadata is recorded in a database.
5. A unique paste identifier is returned to the client.

### Paste retrieval workflow

1. The client requests a paste by its identifier.
2. The backend looks up the corresponding paste metadata.
3. The associated content is retrieved.
4. The result is returned to the client.

These workflows describe the project's intended functionality.

## Technical Focus

This project is an opportunity to gain hands-on experience with:

- **Backend development:** Request handling and service organization.
- **REST API design:** Resource-oriented HTTP endpoints.
- **Authentication:** Separating identity-related functionality from application features.
- **Data persistence:** Working with structured metadata and text content.
- **Object storage:** Exploring S3-compatible storage solutions.
- **Software architecture:** Designing components with clear responsibilities.

## Development Status

This repository is an ongoing learning project.

The initial focus is building the core paste creation and retrieval workflow. Additional functionality and deployment documentation will be included as development progresses.

## Author

**Alan Tokin**

Computer Science Student

[GitHub](https://github.com/minidiva)

---

*My first personal backend project — built to learn by designing and implementing a real application.*
