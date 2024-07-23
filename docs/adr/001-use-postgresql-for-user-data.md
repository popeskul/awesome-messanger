# ADR: Use PostgreSQL for User Data Storage

## Context
We need a reliable and scalable database solution for storing user data, including authentication information, profiles, and friend relationships. The solution should support ACID transactions and provide strong consistency guarantees.

## Decision
We will use PostgreSQL as the primary database for user data storage.

## Status
Accepted

## Consequences
- **Positive**:
    - Strong ACID compliance ensures data consistency.
    - Wide adoption and strong community support.
    - Rich feature set, including JSON support for semi-structured data.
- **Negative**:
    - Requires careful tuning for performance optimization at scale.
    - Vertical scaling limitations compared to some NoSQL databases.

## Alternatives Considered
- **MySQL**: Similar to PostgreSQL but has fewer advanced features and less flexible licensing.
- **MongoDB**: Offers better horizontal scaling but lacks ACID compliance for all operations.
- **Cassandra**: Highly scalable but does not support ACID transactions.

## Implementation
We will set up a PostgreSQL cluster using AWS RDS, configure it for high availability, and ensure regular backups are in place.
