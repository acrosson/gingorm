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

    v1 := r.Group("/v1")
    {
        v1.GET("/people/", controllers.GetPeople)
        v1.GET("/people/:id", controllers.GetPerson)
        v1.POST("/people", controllers.CreatePerson)
        v1.PUT("/people/:id", controllers.UpdatePerson)
        v1.DELETE("/people/:id", controllers.DeletePerson)
    }

    r.Run(":8080")

    defer db.CloseDB()
}


