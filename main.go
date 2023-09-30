package main

import(
	//"fmt"
	"log"
)

func main(){
	store, err := NewPostgresStore()
	if err!=nil {
		log.Fatal(err)
	}
	
	//fmt.Printf("%+v\n",store)
	
	if err := store.Init(); err!=nil{
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", store)
	server.Run()

	//fmt.Println("Yeah Buddy it works!")
	// http.HandleFunc("/login",Login)
	// http.HandleFunc("/home",Home)
	// http.HandleFunc("/refresh",Refresh)

	// log.Fatal(http.ListenAndServe(":8080",nil))
}



//---------------- API ------------------------------------------


//---------------------------------- Types ---------------------------------------------

