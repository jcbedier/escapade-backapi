package main

func main() {
	config := NewConfiguration()
	a := App{}
	a.Initialize(config)

	a.Run(":8080")
}

// type Person struct {
// 	ID        bson.ObjectId `bson:"_id,omitempty"`
// 	Name      string
// 	Phone     string
// 	Timestamp time.Time
// }

// var (
// 	IsDrop = false
// )

// const (
// 	hosts      = "127.0.0.1:27017"
// 	database   = "escapade"
// 	username   = "root"
// 	password   = "example"
// 	collection = "people"
// )

// func main() {
// 	info := &mgo.DialInfo{
// 		Addrs:    []string{hosts},
// 		Timeout:  60 * time.Second,
// 		Database: database,
// 		Username: username,
// 		Password: password,
// 	}

// 	session, err := mgo.DialWithInfo(info)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer session.Close()

// 	session.SetMode(mgo.Monotonic, true)

// 	// Drop Database
// 	if IsDrop {
// 		err = session.DB("escapade").DropDatabase()
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	// Collection People
// 	c := session.DB("escapade").C("people")

// 	// Index
// 	index := mgo.Index{
// 		Key:        []string{"name", "phone"},
// 		Unique:     false,
// 		DropDups:   true,
// 		Background: true,
// 		Sparse:     true,
// 	}

// 	err = c.EnsureIndex(index)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Insert Datas
// 	err = c.Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321", Timestamp: time.Now()},
// 		&Person{Name: "Cla", Phone: "+66 33 1234 5678", Timestamp: time.Now()})

// 	if err != nil {
// 		panic(err)
// 	}

// 	// Query One
// 	result := Person{}
// 	err = c.Find(bson.M{"name": "Ale"}).Select(bson.M{"phone": 0}).One(&result)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Phone", result)

// 	// Query All
// 	var results []Person
// 	err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Results All: ", results)

// 	// Update
// 	colQuerier := bson.M{"name": "Ale"}
// 	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
// 	err = c.Update(colQuerier, change)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Query All
// 	err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Results All: ", results)
// }
