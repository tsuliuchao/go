package apis

import (
 "net/http"
 "log"
 "fmt"
 "strconv"
 "gopkg.in/gin-gonic/gin.v1"
 . "github.com/liuchao/project/models"
)

func IndexApi(c *gin.Context) {
 c.String(http.StatusOK, "It works")
}

func GetPersonApi(c *gin.Context){
	 p := Person{}
	 persons, err := p.GetPersons()
	 if err != nil {
	  log.Fatalln(err)
	 }
	 c.JSON(http.StatusOK, gin.H{
	  "data": persons,
	 })
}

func AddPersonApi(c *gin.Context) {
 firstName := c.Request.FormValue("first_name")
 lastName := c.Request.FormValue("last_name")

 p := Person{FirstName: firstName, LastName: lastName}

 ra, err := p.AddPerson()
 if err != nil {
  log.Fatalln(err)
 }
 msg := fmt.Sprintf("insert successful %d", ra)
 c.JSON(http.StatusOK, gin.H{
  "msg": msg,
 })
}

func GetOnePersonApi(c *gin.Context){
   cid := c.Param("id")
   id, _ := strconv.Atoi(cid)
   p := Person{Id:id}
   ra, _ := p.GetOne(id)
   c.JSON(http.StatusOK, gin.H{
     "msg": ra,
   })
}