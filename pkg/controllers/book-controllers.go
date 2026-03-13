package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TechYash-bit/proj1Go/pkg/models"
	// "github.com/TechYash-bit/proj1Go/pkg/routes"
	"github.com/TechYash-bit/proj1Go/pkg/utils"
	"github.com/gorilla/mux"
)

var newBook models.Book

func GetAllBook(w http.ResponseWriter,r *http.Request){
	newBook:=models.GetAllBook()
	res,_:=json.Marshal(newBook)
	w.Header().Set("Content-type","pkglicaton/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	bookId:=vars["bookId"]
	Id,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("errors while parsing")
	}
	booketails,_:=models.GetBookById(Id)
	res,_:=json.Marshal(booketails)
	w.Header().Set("Content-type","pkglicaton/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter,r *http.Request){
	createBook:=&models.Book{}
	utils.ParseBody(r,createBook)
	b:=createBook.CreateBook()
	res,_:=json.Marshal(b)
	// w.Header().Set("Content-type","pkglicaton/json") 
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	bookId:=vars["bookId"]
	Id,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("errors while parsing")
	}

	books:=models.DeleteBook(Id)
	res,_:=json.Marshal(books)
	w.Header().Set("Content-type","pkglicaton/json") 
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


func UpdateBooks(w http.ResponseWriter, r *http.Request){
	UpdateBook:=&models.Book{}
	utils.ParseBody(r,UpdateBook)
	vars:=mux.Vars(r)
	bookId:=vars["bookId"]
	Id,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Print("Parsing fail of id ")
	}
	bookDetails,db:=models.GetBookById(Id)
	if UpdateBook.Name!=""{
		bookDetails.Name=UpdateBook.Name
	}
	if UpdateBook.Author!=""{
		bookDetails.Author=UpdateBook.Author
	}
	if UpdateBook.Publication!=""{
		bookDetails.Publication=UpdateBook.Publication
	}
	db.Save(&bookDetails)
	res,_:=json.Marshal(bookDetails)
	w.Header().Set("Content-type","pkglicaton/json") 
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
