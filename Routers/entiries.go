package route

import (
    "fmt"
    "net/http"
    "strconv"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/shubhammishra-1/bank-system/CRUD"
    "github.com/shubhammishra-1/bank-system/Models"
    "github.com/shubhammishra-1/bank-system/Database"


)



//::CREATE Entireis

func Create_Ent(w http.ResponseWriter, r *http.Request){

//decoding

decoder := json.NewDecoder(r.Body)

var ent model.Entiries

//note only that fields of Payload will be filled whose values are passed into body

err := decoder.Decode(&ent)

if err != nil {
    panic(err)
}


//opening the database

db , _:=setup.Open_DB()

//creating  entiers  means either money added or removed so change that account also 

crud.Create_Ent(db,ent)

//closing the database

setup.Close_DB(db)



json.NewEncoder(w).Encode("Entiries has been created successfully")



}


//Get Details using Entiries ID


func Get_Ent_Details(w http.ResponseWriter, r *http.Request){

id_str:=mux.Vars(r)["id"]

id,_ := strconv.Atoi(id_str)  // converting id string into integer

//opening the database

db , _:=setup.Open_DB()

//Getting the details of account

var details model.Entiries

details=crud.Get_Ent_Details(db,id)

//closing the database

setup.Close_DB(db)


if details.Id!=0 {

str:=fmt.Sprintf("Here is Details of Entiries ID no. %d",id)


json.NewEncoder(w).Encode(str)

json.NewEncoder(w).Encode(details)

//fmt.Fprint(w,details)


}else{

   str:=fmt.Sprintf("Sorry there is no User exits with ID no. %d",id)
   
   json.NewEncoder(w).Encode(str)


}

}


//Get All Entiries Details


func Get_All_Ent(w http.ResponseWriter, r *http.Request){


//opening the database

db , _:=setup.Open_DB()

//Getting the details of account

var AllEntiries []model.Entiries

AllEntiries=crud.Get_All_Ent(db)

//closing the database

setup.Close_DB(db)


if len(AllEntiries)!=0 {

str:=fmt.Sprintf("Here is Details of All Entiries")


json.NewEncoder(w).Encode(str)

json.NewEncoder(w).Encode(AllEntiries)

//fmt.Fprint(w,details)


}else{

   str:=fmt.Sprintf("Sorry there is no Entiries exits")
   
   json.NewEncoder(w).Encode(str)


}

}