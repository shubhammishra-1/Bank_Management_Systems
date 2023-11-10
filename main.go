package main

import (

    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/shubhammishra-1/bank-system/Routers"

)


func home(w http.ResponseWriter, r *http.Request){


json.NewEncoder(w).Encode("Welcome to Bank Management Systems API")


}






func StartServer(PORT string){

    
    router := mux.NewRouter()

    router.HandleFunc("/home",home)

    //Account router

    router.HandleFunc("/accounts/create",route.Create_Acc).Methods("POST")
    router.HandleFunc("/accounts/id/{id}",route.Get_Acc_Details).Methods("GET")
    router.HandleFunc("/accounts/all",route.Get_All_Acc).Methods("GET")
    router.HandleFunc("/accounts/update/name",route.Update_Name).Methods("PUT")
    router.HandleFunc("/accounts/update/balance",route.Update_Balance).Methods("PUT")
    router.HandleFunc("/accounts/delete/{id}",route.Delete_Acc ).Methods("DELETE")



    //Entiries router

    router.HandleFunc("/entiries/create",route.Create_Ent).Methods("POST")
    router.HandleFunc("/entiries/id/{id}",route.Get_Ent_Details).Methods("GET")
    router.HandleFunc("/entiries/all",route.Get_All_Ent).Methods("GET")
 


    //Transfers router

    router.HandleFunc("/transfers/create",route.Create_Trans).Methods("POST")
    router.HandleFunc("/transfers/id/{id}",route.Get_Trans_Details).Methods("GET")
    router.HandleFunc("/transfers/all",route.Get_All_Trans).Methods("GET")




    log.Println("Server is running on PORT",PORT)
    log.Fatal(http.ListenAndServe(PORT,router))



}



func main(){
    
    var PORT=":9000"


    StartServer(PORT)


}