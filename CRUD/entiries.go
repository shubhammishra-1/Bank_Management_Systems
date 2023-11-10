package crud

import (

    "fmt"
    "database/sql"

    _ "github.com/go-sql-driver/mysql"
    "github.com/shubhammishra-1/bank-system/Models"
)



//create

func Create_Ent(db *sql.DB,ent model.Entiries){

    
    timeFormtted := formted_Time()

    str:=fmt.Sprintf("INSERT INTO ENTIRIES VALUES(%d,%d,%d,'%s')",ent.Id,ent.Account_id,ent.Amount,timeFormtted)
    

	insert, err := db.Query(str)

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }else{

		fmt.Println("inserted data into specified DB")

	}

    //updating the balance

    str=fmt.Sprintf("UPDATE ACCOUNTS SET BALANCE = BALANCE + %d WHERE ID = %d",ent.Amount,ent.Account_id)

    insert, err = db.Query(str)

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }else{

		fmt.Println("updated the balance")

	}



    // be careful deferring Queries if you are using transactions
    defer insert.Close()
    

}


func Get_Ent_Details(db *sql.DB,id int)(details model.Entiries){


    str:=fmt.Sprintf("SELECT * FROM Entiries WHERE ID=%d",id)



	results, err := db.Query(str)
    if err != nil {
        panic(err.Error())
    }

	//there will be only one results so 
   // var details model.Accounts

   for results.Next() {

        // for each row, scan the result into our tag variable [basically names will be get]
        err = results.Scan(&details.Id,&details.Account_id,&details.Amount,&details.Created_At)

        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
    }


    return details
    

}


func Get_All_Ent(db *sql.DB)(AllEntiries []model.Entiries){


    str:=fmt.Sprintf("SELECT * FROM Entiries")

    results, err := db.Query(str)

    if err != nil {
        panic(err.Error()) 
    }
   

   for results.Next() {
        
        var details model.Entiries
        
        err = results.Scan(&details.Id,&details.Account_id,&details.Amount,&details.Created_At)

        if err != nil {
            panic(err.Error()) 
        }

        AllEntiries=append(AllEntiries,details)


    }
   
    
    return AllEntiries


}



