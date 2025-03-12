# Book Search Engine

This project is a complete book search system that crawls the web for book content, indexes it, and provides a user-friendly search interface.

## System Components

- **Web Crawler**: Scrapes book data from the web
- **Index Engine**: Processes and indexes the book data
- **Search Engine**: API service that handles search queries
- **Search UI**: User interface for searching books
- **Supporting Services**: Redis, MinIO, PostgreSQL

## Prerequisites

- Docker and Docker Compose
- Git

## Getting Started

### 1. Clone the repository

```bash
git clone <repository-url>
cd webir-book-search
```

### 2. Run the Web Crawler to collect data

```bash
docker compose up web-crawler
```

Wait until enough books have been crawled. You can monitor the progress through:
- Console output (stdout)
- Connecting to the PostgreSQL database using a database GUI client

Once sufficient data is collected, stop the crawler with `Ctrl+C`.

### 3. Build the search index

```bash
docker compose up index-engine
```

Wait for the indexing process to complete. The console will indicate when indexing is finished.

### 4. Start the Search Engine API

```bash
docker compose up search-engine -d
```

This will start the search API service in the background.

### 5. Launch the Search UI

```bash
docker compose up search-ui -d
```

This will start the user interface in the background.

### 6. Access the Book Search Engine

Open your browser and navigate to:

```
http://localhost:8080
```

## Additional Information

- The search index data is persisted in a Docker volume (`index_data`)
- Book data is stored in PostgreSQL
- Raw HTML content is stored in MinIO
- Redis is used for queue management during crawling

## Configuration

You can modify the configuration parameters in the `docker-compose.yaml` file to adjust:
- Number of crawler threads
- Database credentials
- Storage locations
- Service ports

## Stopping the Services

To stop all services:

```bash
docker compose down
```

To stop a specific service:

```bash
docker compose stop <service-name>
```