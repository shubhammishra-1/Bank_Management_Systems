package crud

import (

    "fmt"
    "database/sql"
    
    _ "github.com/go-sql-driver/mysql"
    "github.com/shubhammishra-1/bank-system/Models"
)




//create

func Create_Trans(db *sql.DB,trans model.Transfers){

    
    timeFormtted := formted_Time()

    str:=fmt.Sprintf("INSERT INTO TRANSFERS VALUES(%d,%d,%d,%d,'%s')",trans.Id,trans.From_Account_Id,trans.To_Account_Id,trans.Amount,timeFormtted)


	insert, err := db.Query(str)

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }else{

		fmt.Println("inserted data into specified DB")

	}

    //Creating two entiries

    //1 for "from_acc" as (-)

    var neg_ent model.Entiries
    
    neg_ent.Id=trans.Id
    neg_ent.Account_id=trans.From_Account_Id
    neg_ent.Amount=-(trans.Amount)
    
    Create_Ent(db,neg_ent)


    //2 for "to_acc" as (+)


    var pos_ent model.Entiries

    pos_ent.Id=trans.Id+1
    pos_ent.Account_id=trans.To_Account_Id
    pos_ent.Amount=trans.Amount

    Create_Ent(db,pos_ent)
 


    // be careful deferring Queries if you are using transactions
    defer insert.Close()
    

}


func Get_Trans_Details(db *sql.DB,id int)(details model.Transfers){


    str:=fmt.Sprintf("SELECT * FROM TRANSFERS WHERE ID=%d",id)



	results, err := db.Query(str)
    if err != nil {
        panic(err.Error())
    }

	//there will be only one results so 
   // var details model.Accounts

   for results.Next() {

        // for each row, scan the result into our tag variable [basically names will be get]
        err = results.Scan(&details.Id,&details.From_Account_Id,&details.To_Account_Id,&details.Amount,&details.Created_At)

        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
    }


    return details
    

}


func Get_All_Trans(db *sql.DB)(ALLTransfers []model.Transfers){


    str:=fmt.Sprintf("SELECT * FROM Transfers")

    results, err := db.Query(str)

    if err != nil {
        panic(err.Error()) 
    }
   

   for results.Next() {
        
        var details model.Transfers
        
        err = results.Scan(&details.Id,&details.From_Account_Id,&details.To_Account_Id,&details.Amount,&details.Created_At)

        if err != nil {
            panic(err.Error()) 
        }

        ALLTransfers=append(ALLTransfers,details)


    }
   
    
    return ALLTransfers


}



