package crud

import (

    "fmt"
    "database/sql"
    "time"
    
    _ "github.com/go-sql-driver/mysql"
    "github.com/shubhammishra-1/bank-system/Models"
)


//time stamp after every UPDATE, CREATE


func formted_Time()(string){
      
    t := time.Now()
    timeFormtted := fmt.Sprintf("%d-%02d-%02d - %02d:%02d:%02d",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

    return timeFormtted

}




//create

func Create_acc(db *sql.DB,acc model.Accounts){

    
    timeFormtted := formted_Time()

    str:=fmt.Sprintf("INSERT INTO ACCOUNTS VALUES(%d,'%s',%d,'%s')",acc.Id,acc.Owner,acc.Balance,timeFormtted)


	insert, err := db.Query(str)

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }else{

		fmt.Println("inserted data into specified DB")

	}

    // be careful deferring Queries if you are using transactions
    defer insert.Close()
    

}


func Get_Acc_Details(db *sql.DB,id int)(details model.Accounts){


    str:=fmt.Sprintf("SELECT * FROM ACCOUNTS WHERE ID=%d",id)



	results, err := db.Query(str)
    if err != nil {
        panic(err.Error())
    }

	//there will be only one results so 
   // var details model.Accounts

   for results.Next() {

        // for each row, scan the result into our tag variable [basically names will be get]
        err = results.Scan(&details.Id,&details.Owner,&details.Balance,&details.Created_At)

        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
    }


    return details
    

}


func Get_All_Acc(db *sql.DB)(AllAccounts []model.Accounts){


    str:=fmt.Sprintf("SELECT * FROM ACCOUNTS")

    results, err := db.Query(str)

    if err != nil {
        panic(err.Error()) 
    }
   

   for results.Next() {
        
        var details model.Accounts
        
        err = results.Scan(&details.Id,&details.Owner,&details.Balance,&details.Created_At)

        if err != nil {
            panic(err.Error()) 
        }

        AllAccounts=append(AllAccounts,details)


    }
   
    
    return AllAccounts


}



func Update_Name(db *sql.DB,acc model.Accounts)(bool){
 
timeFormtted := formted_Time() 

str:=fmt.Sprintf("UPDATE ACCOUNTS SET Owner = '%s',Created_At = '%s' WHERE Id=%d",acc.Owner,timeFormtted,acc.Id)

 
res, err := db.Exec(str)

if err != nil {
    panic(err)
}

//this query returns no. of rows afftected after query defintaly for not insertion it will be 0

isUpdated, err := res.RowsAffected()
if err != nil {
    panic(err)
}


if isUpdated ==0{
    return false
}

return true



}


func Update_Balance(db *sql.DB,acc model.Accounts)(bool){

timeFormtted := formted_Time()

str:=fmt.Sprintf("UPDATE ACCOUNTS SET BALANCE=%d , Created_At = '%s'  WHERE ID=%d",acc.Balance,timeFormtted,acc.Id)


res, err := db.Exec(str)

if err != nil {
    panic(err)
}

//this query returns no. of rows afftected after query defintaly for not insertion it will be 0

isUpdated, err := res.RowsAffected()
if err != nil {
    panic(err)
}


if isUpdated ==0{
    return false
}

return true



}

//Deleting the account


func Delete_Acc(db *sql.DB, id int)(bool){
   
str:=fmt.Sprintf("DELETE FROM ACCOUNTS WHERE ID = %d",id)

res,err:= db.Exec(str)

if err!=nil{
    panic(err)
}


//checking wheather that perticular row deleted or not

isDeleted, err := res.RowsAffected()

if err != nil {
    panic(err)
}


if isDeleted==0{
    return false
}

return true



}

