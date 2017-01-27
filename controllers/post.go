package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"fmt"
	"encoding/json"
	"github.com/abrowley/Sense/models"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func GetSession() *mgo.Session {
	s,err := mgo.Dial("mongodb://localhost")
	if(err != nil){
		panic(err)
	}else{
		fmt.Println("Connected to mongodb")
	}
	return s
}

type(
	PostController interface {
		CreatePost(w http.ResponseWriter, r *http.Request, p httprouter.Params)
		GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params)
		GetPosts(w http.ResponseWriter, r *http.Request, p httprouter.Params)
		DeletePost(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	}
)

type(
	MongoPostController struct {
		session *mgo.Session
	}
)

func NewMongoPostController(s *mgo.Session) *MongoPostController{
	return &MongoPostController{s}
}

// Method to list all posts
func (uc MongoPostController) GetPosts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("GetPosts called")
	var posts []models.Post
	// Fetch all users
	if err := uc.session.DB("sense").C("posts").Find(nil).Limit(100).All(&posts); err != nil {
		w.WriteHeader(404)
		return
	}
	postsJson, _ := json.Marshal(posts);
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", postsJson)
}

// Method to retrieve a post by id
func (uc MongoPostController)  GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Printf("GetPost %s called \n", id)

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	pst := models.Post{}

	// Fetch user
	if err := uc.session.DB("sense").C("posts").FindId(oid).One(&pst); err != nil {
		w.WriteHeader(404)
		return
	}
	pj, _ := json.Marshal(pst);
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", pj)
}

// Method to create a post
func (uc MongoPostController) CreatePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("CreatePost called")
	pst := models.Post{}
	json.NewDecoder(r.Body).Decode(&pst)
	pst.Id = bson.NewObjectId()
	pst.TimeReceived = time.Now()
	uc.session.DB("sense").C("posts").Insert(pst)
	pj, _ := json.Marshal(pst)
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", pj)
}

// Method to delete a post
func (uc MongoPostController) RemovePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")
	fmt.Printf("RemovePost %s called \n", id)

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	// Grab id
	oid := bson.ObjectIdHex(id)
	// Remove user
	if err := uc.session.DB("sense").C("posts").RemoveId(oid); err != nil {
		fmt.Printf("Could not remove post %v the error was %v\n",id,err)
		w.WriteHeader(404)
		return
	}
	fmt.Printf("Removed post %v\n",id)
	// Write status
	w.WriteHeader(200)
}