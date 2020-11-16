package main
        import "github.com/gin-gonic/gin"
        import "fmt"
        func main() {
        	r := gin.Default()
        	r.GET("/ping", func(c *gin.Context) {
        		c.JSON(200, gin.H{
        			"message": "pong",
        		})
        	})
        	fmt.Printf("0.0.0.0:3000")
        	// Listen and Server in 0.0.0.0:3000
        	r.Run(":3000")
        }