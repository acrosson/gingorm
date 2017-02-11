package main
// only need mysql OR sqlite 
// both are included here for reference
import (
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "github.com/alexandercrosson/gingorm/db"
    "github.com/alexandercrosson/gingorm/controllers"
)

func main() {
    // NOTE: See we’re using = to assign the global var
    // instead of := which would assign it only in this function
    //db, err = gorm.Open("sqlite3", "./gorm.db")
    db.Init()

    r := gin.Default()
    r.GET("/people/", controllers.GetPeople)
    r.GET("/people/:id", controllers.GetPerson)
    r.POST("/people", controllers.CreatePerson)
    r.PUT("/people/:id", controllers.UpdatePerson)
    r.DELETE("/people/:id", controllers.DeletePerson)
    r.Run(":8080")

    defer db.CloseDB()
}


