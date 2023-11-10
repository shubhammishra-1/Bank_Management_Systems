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



//::CREATE a transfer  money "from" will (-) & will added to "To" account

//for reusebility create two "Entiries" one for "from_acc" as "-" and another "to_acc" as (+)

func Create_Trans(w http.ResponseWriter, r *http.Request){

//decoding

decoder := json.NewDecoder(r.Body)

var trans model.Transfers

//note only that fields of Payload will be filled whose values are passed into body

err := decoder.Decode(&trans)

if err != nil {
    panic(err)
}


//opening the database

db , _:=setup.Open_DB()


crud.Create_Trans(db,trans)


//closing the database

setup.Close_DB(db)



json.NewEncoder(w).Encode("Transfer has been Done")



}


//Get Details using Transfer ID


func Get_Trans_Details(w http.ResponseWriter, r *http.Request){

id_str:=mux.Vars(r)["id"]

id,_ := strconv.Atoi(id_str)  // converting id string into integer

//opening the database

db , _:=setup.Open_DB()

//Getting the details of account

var details model.Transfers

details=crud.Get_Trans_Details(db,id)

//closing the database

setup.Close_DB(db)


if details.Id!=0 {

str:=fmt.Sprintf("Here is Details of Transfer with ID no. %d",id)


json.NewEncoder(w).Encode(str)

json.NewEncoder(w).Encode(details)

//fmt.Fprint(w,details)


}else{

   str:=fmt.Sprintf("Sorry there is no User exits with ID no. %d",id)
   
   json.NewEncoder(w).Encode(str)


}

}


//Get All Transfers Details


func Get_All_Trans(w http.ResponseWriter, r *http.Request){


//opening the database

db , _:=setup.Open_DB()

//Getting the details of account

var ALLTransfers []model.Transfers

ALLTransfers=crud.Get_All_Trans(db)

//closing the database

setup.Close_DB(db)


if len(ALLTransfers)!=0 {

str:=fmt.Sprintf("Here is Details of All Transfers")


json.NewEncoder(w).Encode(str)

json.NewEncoder(w).Encode(ALLTransfers)

//fmt.Fprint(w,details)


}else{

   str:=fmt.Sprintf("Sorry there is no Transfers exits")
   
   json.NewEncoder(w).Encode(str)


}

}