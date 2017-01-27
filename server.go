package main

import (
	"flag"
	"github.com/julienschmidt/httprouter"
	"os"
	"fmt"
	"net/http"
	"github.com/abrowley/Sense/controllers"
	"github.com/rs/cors"
)

var port  = flag.Int("port",6000,"The port of which to run the server")

func runServer(port int) (err error){
	wd, _ := os.Getwd()
	fmt.Printf("The working directory is %s\n", wd)
	r := httprouter.New()
	//r.ServeFiles("/static/*filepath", http.Dir(wd))
	pc := controllers.NewMongoPostController(controllers.GetSession())
	r.POST("/api/post",pc.CreatePost)
	r.GET("/api/posts", pc.GetPosts)
	r.GET("/api/post/:id", pc.GetPost)
	r.DELETE("/api/post/:id", pc.RemovePost)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET","POST","DELETE","PATCH","PUT","OPTIONS"},
		AllowedHeaders: []string{
			"Origin",
			"Content-Type",
			"X-Auth-Token",
			"Access-Control-Allow-Origin",
			"x-requested-with"},
		AllowCredentials: true,
	})
	handler:= c.Handler(r)
	// Fire up the server
	server_hostname := fmt.Sprintf("localhost:%v",port)
	fmt.Printf("Running on %s \n", server_hostname)
	fmt.Println(http.ListenAndServe(server_hostname, handler))
	return err
}

func main(){
	flag.Parse()
	runServer(*port)
}