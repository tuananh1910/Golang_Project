package api

import (
	"awesomeProject/src/api/router"
	"awesomeProject/src/config"
	"fmt"
	"log"
	"net/http"
)

func Run(){
	config.Load()
	fmt.Printf("\n\tListening [::]:%d\n", config.PORT)
	listen(config.PORT)

}

func listen(port int)  {
	r := router.NEW()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))

}

