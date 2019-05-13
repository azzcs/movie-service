package entity


import (
	"movie-show/commons"
	"time"
)

type Movie struct {
	Id string `gorm:"primary_key"`
	Name, Alias, UpdateStatus, Nav1, Nav2, Img, Director, Actor, Type, Area, Language, Time, UpdateTime, Introduce, Content, Enable, CreateTime string
	MovieContent commons.MovieContent
}
type MovieNew struct {
	Id string `gorm:"primary_key"`
	Name, Alias, UpdateStatus, Nav1, Nav2, Img, Director, Actor, Type, Area, Language, Time, UpdateTime, Introduce, Content string
	Enable bool
	CreateTime time.Time
	MovieContent commons.MovieContent
}

type PageInfo struct {
	Name     string
	Nav1     string
	Nav2     string
	Pages    int
	PageNum  int
	PageSize int
	PrePage  int
	NextPage int
	List     []*Movie
}

type SearchParam struct {
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
	Name     string `form:"name"`
	Nav1     string `form:"nav1"`
	Nav2     string `form:"nav2"`
}