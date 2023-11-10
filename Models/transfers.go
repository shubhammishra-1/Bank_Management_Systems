package model



type Transfers struct{
    Id int                  `json: "id"`
    From_Account_Id int     `json: "from"`
    To_Account_Id   int     `json:  "to"`
    Amount  int             `json:  "amount"`
    Created_At string       `json:  "createdAT"`
}