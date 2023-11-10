package model



type Entiries struct{
    Id int                  `json: "id"` 
    Account_id int          `json: "accountId"`
    Amount int              `json: "amount"`
    Created_At string       `json: "createdAT"`
}