package controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/julienschmidt/httprouter"
// 	"gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"
// )

// type UserController struct {
// 	session *mgo.Session
// }

// func NewUserController(s *mgo.Session) *UserController {
// 	return &UserController{s}
// }
// func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	id := p.ByName("id")
// 	if !bson.IsObectIdHex(id) {
// 		w.WriteHeader(http.StatusNotFound)
// 	}
// 	oid := bson.ObjectHex(id)
// 	u := models.User{}
// 	if err := uc.session.DB("monogo-golang").C("users").FindId(oid).One(&u); err != nil {
// 		w.WriteHeader(404)
// 		return
// 	}
// 	u, err := json.Marshal(u)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
// func main() {

// }
//
