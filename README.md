# GraphQL Course Management API

A GraphQL server built with Go that manages courses and categories. This project uses [gqlgen](https://github.com/99designs/gqlgen) for GraphQL implementation and SQLite as the database.

## Features

- **Categories Management**: Create and list categories
- **Courses Management**: Create and list courses with category relationships
- **GraphQL Playground**: Interactive interface for testing queries and mutations
- **SQLite Database**: Lightweight database for data persistence
- **Relational Data**: Full support for course-category relationships

## Tech Stack

- **Go 1.24.4**
- **GraphQL** with [gqlgen](https://github.com/99designs/gqlgen)
- **SQLite** database
- **UUID** for unique identifiers

## Project Structure

```
server-graphql/
├── graph/
│ ├── model/ # Generated GraphQL models
│ ├── schema.graphqls # GraphQL schema definition
│ ├── schema.resolvers.go # GraphQL resolvers
│ ├── resolver.go # Main resolver struct
│ └── generated.go # Generated GraphQL code
├── internal/
│ └── database/ # Database layer
│ ├── category.go # Category repository
│ └── course.go # Course repository
├── server/
│ └── server.go # HTTP server setup
├── data.db # SQLite database file
├── gqlgen.yml # gqlgen configuration
├── tools.go # Build tools
├── go.mod # Go modules
└── go.sum # Go modules checksums
```
## Installation & Setup

1. **Clone the repository**

   ```bash
   git clone https://github.com/sergioc0sta/server-graphql
   cd server-graphql

   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Generate GraphQL code**

   ```bash
   go run github.com/99designs/gqlgen generate
   ```

4. **Run the server**

   ```bash
   go run server/server.go
   ```

5. **Access GraphQL Playground**
   Open your browser and navigate to: `http://localhost:8080/


GraphQL Schema

### Types

```GraphQL
type Category {
  id: ID!
  name: String!
  description: String
  courses: [Courses!]!
}

type Courses {
  id: ID!
  name: String!
  description: String
  category: Category!
}

input NewCategory {
  name: String!
  description: String
}

input NewCourse {
  name: String!
  description: String
  categoryId: ID!
}

type Query {
  categories: [Category!]!
  courses: [Courses!]!
}

type Mutation {
  createCategory(input: NewCategory!): Category!
  createCourse(input: NewCourse!): Courses!
}
```
API Examples

### Mutations

#### Create a New Category
```graphql
mutation NewCategory {
  createCategory(input: {name: "mistas", description: "pao ou rosca"}) {
    id
    name
    description
  }
}
```

**Expected Response:**
```json
{
  "data": {
    "createCategory": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "name": "mistas",
      "description": "pao ou rosca"
    }
  }
}
```

#### Create a New Course

```graphql
mutation NewCourse {
  createCourse(input: {
    name: "novo",
    description: "top",
    categoryId: "918d7b83-ca5c-4916-8988-63e3b44b974c"
  }) {
    id
    name
    description
    category {
      id
      name
      description
    }
  }
}
```

**Expected Response:**
```json
{
  "data": {
    "createCourse": {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "name": "novo",
      "description": "top",
      "category": {
        "id": "918d7b83-ca5c-4916-8988-63e3b44b974c",
        "name": "mistas",
        "description": "pao ou rosca"
      }
    }
  }
}
```

### Queries

#### Get All Categories

```graphql
query FindCategories {
  categories {
    id
    name
    description
  }
}
```

#### Get Categories with Their Courses

```graphql
query FindCategoriesAndCourses {
  categories {
    id
    name
    description
    courses {
      id
      name
      description
    }
  }
}
```

#### Get All Courses

```graphql
query FindCourses {
  courses {
    id
    description
    name
  }
}
```

#### Get Courses with Their Categories

```graphql
query FindCoursesAndCategory {
  courses {
    name
    description
    id
    category {
      name
      description
      id
    }
  }
}
```

## Database Schema

The application uses SQLite with the following tables:

### Categories Table
```sql
CREATE TABLE categories (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT
);
```

### Courses Table
```sql
CREATE TABLE courses (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  category_id TEXT,
  FOREIGN KEY (category_id) REFERENCES categories(id)
);
```

Development

### Generate GraphQL Code

After modifying the schema file (`graph/schema.graphqls`), regenerate the code:

```bash
go run github.com/99designs/gqlgen generate
```

### Adding New Resolvers

1. Update the GraphQL schema in `graph/schema.graphqls`
2. Run code generation
3. Implement the resolver methods in `graph/schema.resolvers.go`

### Database Operations

The database layer is located in `internal/database/`:
- `category.go` - Handles category CRUD operations
- `course.go` - Handles course CRUD operations

## Configuration

### Environment Variables

- `PORT` - Server port (default: 8080)

### gqlgen Configuration

The gqlgen configuration is in `gqlgen.yml`. Key settings:

- **Schema location**: `graph/*.graphqls`
- **Generated models**: `graph/model/models_gen.go`
- **Resolvers**: `graph/schema.resolvers.go`

## Testing with GraphQL Playground

1. Start the server: `go run server/server.go`
2. Open `http://localhost:8080/` in your browser
3. Use the examples above to test the API
4. Explore the schema using the built-in documentation
