package model

type Course struct {
	Id        int
	Title     string
	CreatedAt string
}

type Unit struct {
	Id         int
	CourseId   int
	Title      string
	OrderIndex int
	CreatedAt  string
}

type UnitComponent struct {
	Id         int
	UnitId     int
	Type       string
	Title      string
	OrderIndex int
	CreatedAt  string
}

type Explanation struct {
	Id              int
	UnitComponentId int
	Content         string
	OrderIndex      int
}
