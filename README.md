# Code Snippet Application

This is a simple web application for managing code snippets. It is built using Go and provides basic functionality to create, view, and list code snippets.

## Features

- Create new code snippets
- View existing code snippets
- List all code snippets

## Requirements

- Go 1.16 or higher

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/skyespirates/codsnips.git
    cd codsnips
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Run the application:
    ```sh
    go run cmd/web/main.go
    ```

2. Open your web browser and navigate to `http://localhost:8080`.

## Project Structure

- `cmd/web/` - Main application entry point
- `internal/models/` - Data models and database interactions
- `ui/html/` - HTML templates
- `ui/static/` - Static files (CSS, JavaScript, images)

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.