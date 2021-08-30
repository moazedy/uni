package couchbaseQueries

const (
	WriteNewCuorseQuery = `
  INSERT INTO cuorse (KEY, VALUE) VALUES ($1, $2)
  `

	ReadCourseQuery = `
  SELECT cuorse.* FROM cuorse WHERE id= $1 
  `

	GetCuorseListQuery = `
  SELECT cuorse.* FROM cuorse LIMIT $1
  `
)
