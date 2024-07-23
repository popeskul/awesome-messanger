# ADR: Use Cassandra for Messages

## Context
We need a database solution that can handle a high volume of messages with low latency and high availability. The solution should support horizontal scaling and provide eventual consistency.

## Decision
We will use Cassandra as the primary database for storing messages.

## Status
Accepted

## Consequences
- **Positive**:
    - High write and read throughput.
    - Supports horizontal scaling.
    - High availability and fault tolerance.
- **Negative**:
    - Eventual consistency model.
    - Complex maintenance and operations.

## Alternatives Considered
- **PostgreSQL**: ACID compliance but less suitable for high write throughput.
- **MongoDB**: Better horizontal scaling but lacks the same level of fault tolerance as Cassandra.

## Implementation
We will set up a Cassandra cluster, configure it for high availability, and integrate it with the Message Service.
