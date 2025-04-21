# Go + Templ Fullstack Application Boilerplate

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue.svg)](https://golang.org/dl/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) <!-- Choose your license -->

A project boilerplate designed to kickstart the development of fullstack web applications using Go for the backend and [Templ](https://templ.guide/) for type-safe HTML templating.

## Overview

This boilerplate provides a structured starting point for building modern web applications with Go. It integrates Templ for generating server-side rendered HTML, potentially combined with tools like HTMX and Tailwind CSS (or your preferred CSS framework) for a dynamic user experience without writing complex JavaScript.

## Features

-   **Go Backend:** Solid foundation using idiomatic Go.
-   **Templ:** Type-safe HTML templating directly in Go.
-   **Server-Side Rendering (SSR):** Fast initial page loads and SEO benefits.
-   **Project Structure:** Organized layout for scalability and maintainability.
-   **Live Reload:** Integrated development workflow using `air` and `make`.
-   **Static File Serving:** Configured for serving CSS, JS, and images.
-   **Basic Routing:** (Specify your router, e.g., `net/http`, `chi`, `gorilla/mux`) Example routes included.
-   **Database Migrations:** Simple migration management via `make` commands.
-   **Frontend Build:** Integrated Tailwind CSS and Bun build processes.
-   **Environment Configuration:** (Optional, e.g., using `.env` files) Easy configuration management.
-   **Makefile:** Streamlined common development tasks.
-   **[Add other features specific to your boilerplate]**

## Tech Stack

-   **Language:** Go (version 1.21 or higher recommended)
-   **Templating:** Templ
-   **Routing:** [Specify Router, e.g., `net/http` stdlib, `go-chi/chi`, `gorilla/mux`]
-   **CSS:** Tailwind CSS
-   **JavaScript Bundling:** Bun
-   **Live Reload:** Air, Templ watch mode
-   **Task Runner:** Make
-   **Database:** PostgreSQL (as indicated by `psql` and migration commands)
-   **Migrations:** Custom Go script (`./cmd/migrate/main.go`)
-   **[Add/Adjust based on your actual stack]**

## Prerequisites

Before you begin, ensure you have the following installed:

-   Go (version 1.21+)
-   Templ CLI (`go install github.com/a-h/templ/cmd/templ@latest`)
-   Air (`go install github.com/cosmtrek/air@latest`)
-   Node.js & npm (for `npx tailwindcss`)
-   Bun (`curl -fsSL https://bun.sh/install | bash` or other methods)
-   Make
-   PostgreSQL Client (`psql`) - for `make db_login`
-   (Optional) PostgreSQL Server - if running the database locally.

## Getting Started

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/your-repo-name.git
    cd your-repo-name
    ```

2.  **Install Go dependencies:**

    ```bash
    go mod tidy
    ```

3.  **Install Node.js dependencies (for Tailwind):**

    ```bash
    npm install
    ```

    _(This assumes you have a `package.json` defining `@tailwindcss/cli`)_

4.  **(Optional) Set up environment variables:**

    -   Copy the example environment file (if you have one):
        ```bash
        cp .env.example .env
        ```
    -   Update the `.env` file with your specific configuration (database credentials, ports, etc.).

5.  **Generate initial Templ components:**

    ```bash
    make templ
    ```

6.  **(Optional) Set up the database:**
    -   Ensure your PostgreSQL server is running and accessible.
    -   Run the database migrations:
        ```bash
        make db_up
        ```
    -   _(You might need to create the database (`dev_db`) and user (`admin`) manually first if they don't exist)_

## Usage

### Development

To run the application in development mode with live reloading for Go code, Templ templates, CSS, and JS:

1.  **Start the development server and watchers:**

    ```bash
    make dev
    ```

    -   This command runs `air` (which rebuilds/restarts your Go app on changes) and `templ generate -watch` in the background.

2.  **Start the CSS watcher (in a separate terminal):**

    ```bash
    make css-watch
    ```

3.  **Start the JS watcher (in a separate terminal):**
    ```bash
    make js-watch
    ```

The application should now be running (check `air` output or your `.air.toml` for the address, typically `http://localhost:3000` or similar). Changes to Go files, `.templ` files, input CSS, and input JS should trigger automatic rebuilds and browser refreshes (if `air` is configured for it).

### Production Build

_(Your current Makefile doesn't have explicit production build targets, but here's how you would typically do it, potentially adding these targets to your Makefile)_

1.  **Clean previous builds (optional):**
    ```bash
    make clean
    ```
2.  **Generate Templ components:**
    ```bash
    make templ
    ```
3.  **Build frontend assets:**

    ```bash
    # Build CSS (adjust output path if different from dev)
    npx @tailwindcss/cli -i ./frontend/css/input.css -o ./public/css/main.css --minify

    # Build JS (adjust output path if different from dev)
    bun build --outdir ./public/js --target node ./frontend/js/*.js --minify
    ```

    -   _(Suggestion: Add `make build-css` and `make build-js` targets to your Makefile for these)_

4.  **Build the Go application:**

    ```bash
    go build -o ./bin/server ./cmd/server/main.go # Adjust paths as needed
    ```

    -   _(Suggestion: Add a `make build` target for this)_
    -   _(Consider adding `-ldflags="-s -w"` for smaller production binaries)_

5.  **Run the compiled binary:**
    ```bash
    ./bin/server
    ```
    -   Ensure any required environment variables are set in the production environment.

## Makefile Targets

Here's a summary of the available `make` commands:

-   `make clean`: Removes generated files (`tmp/`, CSS, JS).
-   `make db_login`: Attempts to log into the development database using `psql`.
-   `make db_up`: Runs database migrations upwards.
-   `make db_down`: Runs database migrations downwards.
-   `make db_reset`: Runs `db_down` then `db_up`.
-   `make templ`: Generates Go code from Templ files once.
-   `make templ-watch`: Watches Templ files and generates Go code on changes.
-   `make css-watch`: Watches input CSS files and rebuilds using Tailwind CSS on changes.
-   `make js-watch`: Watches input JS files and rebuilds using Bun on changes.
-   `make dev`: Starts the `air` live reloader and `templ-watch` concurrently. _Note: You still need to run `make css-watch` and `make js-watch` in separate terminals for full live-reloading._

_(Consider adding `build`, `build-css`, `build-js` targets for production builds)_

## Project Structure
