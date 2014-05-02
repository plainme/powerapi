package main

import ("github.com/go-martini/martini";
		"github.com/martini-contrib/render";
		"github.com/martini-contrib/binding";
		"net/http";
		"labix.org/v2/mgo";
		"labix.org/v2/mgo/bson";
		"fmt";)

type Msg struct {
    Name        string `form:"name"`
    Description string `form:"description"`
}

func main() {
	fmt.Print("Start")
	// init martini-framework
	mObj := martini.Classic()
	/* Add middleware (auth in this case, similiar to adding it
	to the Get method
	.Get( "/greetings/:name", Auth (<--- here), 
			func(params martini.Params) string {
	*/		
	//mObj.Use(Auth)
	mObj.Use(render.Renderer())
	mObj.Use(DB())
	
	
	
	mObj.Get("/test/", 
			func() string {
				return "Test lübbt"
			})
			
	mObj.Get("/infos", 
			func(r render.Render, db *mgo.Database) {
        		//r.HTML(200, "/home/sascha//workspace/firsttest/src/github.com/sheisig0707/templates/test.tpl", GetAll(db))
        		r.HTML(200, "infos", GetAll(db))
   			})
			
	mObj.Get("/htmlmsg", 
			func(r render.Render, db *mgo.Database) {
        		//r.HTML(200, "/home/sascha//workspace/firsttest/src/github.com/sheisig0707/templates/test.tpl", GetAll(db))
        		r.HTML(200, "test", GetAll(db))
   			})
   	
   	mObj.Post("/htmlmsg", binding.Form(Msg{}), 
   		func(msg Msg, r render.Render, db *mgo.Database) {
        db.C("messages").Insert(msg)
        r.HTML(200, "test", GetAll(db))
    })
   	
   	
   	mObj.Get("/api/addmsg/:user/:msg",  
   		func(params martini.Params, r render.Render, db *mgo.Database) {
        fmt.Print("User ist: " + params["user"])
        msg := Msg{Name: params["user"], Description: params["msg"]}
        
        //msg.Name = params["user"]
        //msg.Name = params["msg"]
        db.C("messages").Insert(msg)
        //r.HTML(200, "test", GetAll(db))
        r.HTML(200, "userMessagesAdded", GetByUser(db, params["user"]))
    })
   	
   	mObj.Get("/api/getAllmsg/:format",  
   		func(params martini.Params, r render.Render, db *mgo.Database) {
        if params["format"] == "json" {
        	r.JSON(200, GetAll(db))
        } else {
        	r.HTML(200, "allMessages", GetAll(db))
        }
    })
   	
   	mObj.Get("/api/getUsermsg/:user/:format", 
			func(params martini.Params, r render.Render, db *mgo.Database) {
        		//r.HTML(200, "/home/sascha//workspace/firsttest/src/github.com/sheisig0707/templates/test.tpl", GetAll(db))
        		if params["format"] == "json" {
        			r.JSON(200, GetAll(db))
        		} else {
        			r.HTML(200, "userMessages", GetByUser(db, params["user"]))
        		}
   			})
   			
	mObj.Run()
	
}		

func Auth(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Api-Key") != "testpassword" {
		http.Error(res, "Wadd? Wer bist du denn?", 401)
	}
}


// DB Returns a martini.Handler
func DB() martini.Handler {
    session, err := mgo.Dial("mongodb://localhost")
    if err != nil {
        panic(err)
    }

    return func(c martini.Context) {
        s := session.Clone()
        //c.Map(s.DB("advent"))
        c.Map(s.DB("mydb"))
        defer s.Close()
        c.Next()
    }
}

// GetAll returns all Wishes in the database
func GetAll(db *mgo.Database) []Msg {
    var msgList []Msg
    err := db.C("messages").Find(nil).All(&msgList)
    if err != nil {
    	panic(err)
    }
    return msgList
}

// GetAll returns all Wishes in the database
func GetByUser(db *mgo.Database, user string ) []Msg {
    var msgList []Msg
    err := db.C("messages").Find(bson.M{"name" : user}).All(&msgList)
    if err != nil {
    	panic(err)
    }
    return msgList
}