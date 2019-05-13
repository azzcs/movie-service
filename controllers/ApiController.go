package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"movie-show/commons"
	"movie-show/entity"
	"movie-show/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func ApiTest(c *gin.Context)  {
	c.JSON(http.StatusOK,"成功")
}
func ApiIndexGet(c *gin.Context)  {
	view := c.Request.Header.Get(commons.ViewPrefix)+"index.tmpl"
	tvList := []entity.Movie{}
	filmList := []entity.Movie{}
	varietyList := []entity.Movie{}
	cartoonList := []entity.Movie{}
	now := time.Now()
	fmt.Printf("开始查询时间 ：%s",now)
	models.Db.Table("movie").Where(" nav1 = ? and enable = 1 ","连续剧").Order(" update_time desc ").Limit(9).Find(&tvList)
	models.Db.Table("movie").Where(" nav1 = ? and enable = 1 ","电影片").Order(" update_time desc ").Limit(9).Find(&filmList)
	models.Db.Table("movie").Where(" nav1 = ? and enable = 1 ","综艺片").Order(" update_time desc ").Limit(9).Find(&varietyList)
	models.Db.Table("movie").Where(" nav1 = ? and enable = 1 ","动漫片").Order(" update_time desc ").Limit(9).Find(&cartoonList)
	end := time.Now()
	fmt.Printf("结束查询时间 ：%s",end)
	c.HTML(http.StatusOK,view,gin.H{
		"title":"title",
		"tvList":tvList,
		"filmList":filmList,
		"varietyList":varietyList,
		"cartoonList":cartoonList,
	})
}

func ApiDetailGet(c *gin.Context)  {
	id := c.Param("id")
	movie :=entity.Movie{}
	models.Db.Table("movie").Where(" id = ? and enable = 1 ",strings.Split(id,".")[0]).Order(" update_time desc ").Limit(1).Find(&movie)
	movie.MovieContent = commons.AnalysisMovieContentJson(movie.Content)
	c.JSON(http.StatusOK,movie)
}


func ApiPlay(c *gin.Context)  {
	id := c.Param("id")
	fmt.Printf("播放页id：%s\n",id)
	num, _ := strconv.Atoi(c.Param("num"))
	playType := strings.Split(c.Param("playType"),".")[0]
	view := c.Request.Header.Get(commons.ViewPrefix)+playType+".tmpl"
	movie :=entity.Movie{}
	models.Db.Table("movie").Where(" id = ? and enable = 1 ",strings.Split(id,".")[0]).Order(" update_time desc ").Limit(1).Find(&movie)
	movieContent := commons.AnalysisMovieContentJson(movie.Content)
	var currentMovie string
	if playType == "kuyun"{
		currentMovie = movieContent.Kuyun[num]
	}else {
		currentMovie = movieContent.Ckm3u8[num]
	}
	c.HTML(http.StatusOK,view,gin.H{
		"title":"title",
		"movie":movie,
		"name":strings.Split(currentMovie,"$")[0],
		"url":strings.Split(currentMovie,"$")[1],
		"playType":playType,
	})

}

func ApiSearch(c *gin.Context)  {
	fmt.Printf("进入方法时间：%s",time.Now())
	var search entity.SearchParam
	c.ShouldBindWith(&search,binding.Form)
	if search.PageNum==0{
		search.PageNum = 1
	}
	if search.PageSize==0{
		search.PageSize = 9
	}
	var list []*entity.Movie
	table := models.Db.Table("movie").Where(" enable = 1 ")
	if search.Name != ""{
		table = table.Where(" name like ? ", "%"+search.Name+"%")
	}
	if search.Nav1!=""{
		table = table.Where(" nav1 = ? ",search.Nav1)
	}
	if search.Nav2!=""{
		table = table.Where(" nav2 = ? ",search.Nav2)
	}
	if search.Nav1 != "福利片"{
		table = table.Where(" nav1 <> ? ","福利片")
	}
	count := 0
	fmt.Printf("开始查询前时间：%s",time.Now())
	table.Offset((search.PageNum - 1) * search.PageSize).
		Order(" time desc, update_time desc ").
		//Count(&count).
		Limit(search.PageSize).
		Select("id,name,img,update_status").
		Find(&list)
	fmt.Printf("查询结束后时间：%s",time.Now())
	pages:=0
	if count%search.PageSize==0{
		pages =count/search.PageSize
	}else {
		pages =count/search.PageSize +1
	}
	pageInfo := entity.PageInfo{
		Name:search.Name,
		Nav1:search.Nav1,
		Nav2:search.Nav2,
		PageNum:search.PageNum,
		PageSize:search.PageSize,
		Pages:pages,
		List:list,
		PrePage:search.PageNum-1,
		NextPage:search.PageNum+1,
	}
	fmt.Print(pageInfo)
	c.JSON(http.StatusOK,pageInfo)
}
