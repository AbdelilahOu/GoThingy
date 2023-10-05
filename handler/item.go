package handler

import (
	"fmt"
	"net/http"

	"github.com/AbdelilahOu/GoThingy/repository"
)

type OrderItem struct {
	Repo *repository.OrderItemRepo
}

func (o *OrderItem) Create(w http.ResponseWriter, r *http.Request) {

}

func (o *OrderItem) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get orderItems")
}

func (o *OrderItem) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get orderItem")
}

func (o *OrderItem) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update orderItem")
}

func (o *OrderItem) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete orderItem")
}