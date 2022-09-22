package main

import{
	"fmt"
	"log"
	"net/http"
}


func formhandler(w,http.Responsewriter,r *http.Request){
	if err := r.Perseform():err != nil{
		fmt.fprintf(w,"Perseform() err:%v\n",err)
		return

	}
	fmt.fprintf(w,"POST request successfull")
	name=r.formvalue("name")
	address=r.formvalue("address")
	fmt.fprintf(w,"name=%s\n",name)
	fmt.fprintf(w,"address=%s\n",address) 
}



func hellohandler(w http.Responsewriter,r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w,"404 not found",http.StatusNotFound)
		return 

	}
	if r.Method != "GET" {
		http.Error(w,"Method not supported",http.StatusNotFound)
		return
	}
	fmt.fPrintf(w,"hello!")

}


func main(){
	fileServer=http.fileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.Handlefunc("/form",formhandler)
	http.Handlfunc("/hello",hellohandler)

	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.fatal(err)
	}


}