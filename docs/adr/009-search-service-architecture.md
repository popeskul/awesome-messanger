# ADR: Search Service Architecture

## Context
The Search Service is responsible for searching users by nickname and other criteria.

## Decision
We will design the Search Service using a microservice architecture, with PostgreSQL as the primary database for indexing and searching user data.

## Status
Accepted

## Consequences
- **Positive**:
    - Clear separation of concerns.
    - Scalable and maintainable service.
    - Strong consistency for search index data.
- **Negative**:
    - Requires careful management of search index updates and concurrency.
    - Potentially complex interactions with other services.

## Alternatives Considered
- **Monolithic Architecture**: Easier to implement initially but harder to scale and maintain.
- **Elasticsearch**: Better for full-text search but adds operational complexity.

## Implementation
We will implement the Search Service with RESTful APIs, using PostgreSQL for data storage and Kafka for event publishing.
