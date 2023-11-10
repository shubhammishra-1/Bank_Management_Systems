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



//::CREATE Accounts

func Create_Acc(w http.ResponseWriter, r *http.Request){

//decoding

decoder := json.NewDecoder(r.Body)

var acc model.Accounts

//note only that fields of Payload will be filled whose values are passed into body

err := decoder.Decode(&acc)

if err != nil {
    panic(err)
}


//opening the database

db , _:=setup.Open_DB()

//creating the account

crud.Create_acc(db,acc)

//closing the database

setup.Close_DB(db)



json.NewEncoder(w).Encode("Account has been created successfully")



}


//Get Details using Account ID


func Get_Acc_Details(w http.ResponseWriter, r *http.Request){

id_str:=mux.Vars(r)["id"]

id,_ := strconv.Atoi(id_str)  // converting id string into integer

//opening the database

db , _:=setup.Open_DB()

//Getting the details of account

var details model.Accounts

details=crud.Get_Acc_Details(db,id)

//closing the database

setup.Close_DB(db)


if details.Owner!="" {

str:=fmt.Sprintf("Here is Details of Account ID no. %d",id)


json.NewEncoder(w).Encode(str)

json.NewEncoder(w).Encode(details)

//fmt.Fprint(w,details)


}else{

   str:=fmt.Sprintf("Sorry there is no User exits with ID no. %d",id)
   
   json.NewEncoder(w).Encode(str)


}

}


//Get All Accounts Details


func Get_All_Acc(w http.ResponseWriter, r *http.Request){


//opening the database

db , _:=setup.Open_DB()

//Getting the details of account

var AllAccounts []model.Accounts

AllAccounts=crud.Get_All_Acc(db)

//closing the database

setup.Close_DB(db)


if len(AllAccounts)!=0 {

str:=fmt.Sprintf("Here is Details of All Accounts")


json.NewEncoder(w).Encode(str)

json.NewEncoder(w).Encode(AllAccounts)

//fmt.Fprint(w,details)


}else{

   str:=fmt.Sprintf("Sorry there is no Account exits")
   
   json.NewEncoder(w).Encode(str)


}

}


//UPDATE THE ACCOUNT' NAME   using ID   {name must be passed in json body}

func Update_Name(w http.ResponseWriter , r *http.Request){
   
//decoding

decoder := json.NewDecoder(r.Body)

var acc model.Accounts


err := decoder.Decode(&acc)

if err != nil {
    panic(err)
}



//opening the database

db , _:=setup.Open_DB()

//updating the name of account

var isUpdate bool

isUpdate=crud.Update_Name(db,acc)

//closing the database

setup.Close_DB(db)



if isUpdate {

    str:=fmt.Sprintf("Successfully Name of Id = %d has been Updated!",acc.Id)

    json.NewEncoder(w).Encode(str)

}else{

    str:=fmt.Sprintf("Sorry Can't Update the name, either ID = %d doesnot exits OR Same Name passed.",acc.Id)

    json.NewEncoder(w).Encode(str)


}



}





//UPDATE THE ACCOUNT' BALANCE using ID  {again it also must be in the form json}

func Update_Balance(w http.ResponseWriter , r *http.Request){
   
//decoding

decoder := json.NewDecoder(r.Body)

var acc model.Accounts


err := decoder.Decode(&acc)

if err != nil {
    panic(err)
}

fmt.Println("SSSSS")

//opening the database

db , _:=setup.Open_DB()

//updating the name of account

var isUpdate bool

isUpdate=crud.Update_Balance(db,acc)

//closing the database

setup.Close_DB(db)



if isUpdate {

    str:=fmt.Sprintf("Successfully Balance of Id = %d has been Updated!",acc.Id)

    json.NewEncoder(w).Encode(str)

}else{

    str:=fmt.Sprintf("Sorry Can't Update the Balanace of Account, either ID = %d doesnot exits OR Same Name passed.",acc.Id)

    json.NewEncoder(w).Encode(str)


}



}


//Deleting account by Id

func Delete_Acc(w http.ResponseWriter, r *http.Request){

//getting id     

id_str:=mux.Vars(r)["id"]

id,_ := strconv.Atoi(id_str)  // converting id string into integer

//opening the database

db , _:=setup.Open_DB()

//performing Delete Operation

var isDeleted bool

isDeleted=crud.Delete_Acc(db,id)

//closing the database

setup.Close_DB(db)



if isDeleted {

    str:=fmt.Sprintf("Successfully Deleted the Account whose Id = %d ",id)

    json.NewEncoder(w).Encode(str)

}else{

    str:=fmt.Sprintf("Sorry Can't Delete the Account, Account for ID = %d doesnot exits.",id)

    json.NewEncoder(w).Encode(str)


}





}







