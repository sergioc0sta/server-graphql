package graph

import "github.com/sergioc0sta/server-graphql/internal/database"

type Resolver struct{
	CategoryDB *database.Category 
	CourseDB *database.Course
}
