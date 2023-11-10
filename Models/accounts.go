package model



type Accounts struct{
    Id int                 `json: "id"`
    Owner string           `json: "owner"`
    Balance int            `json: "balance"`
    Created_At string      `json: "createdAT"`

}