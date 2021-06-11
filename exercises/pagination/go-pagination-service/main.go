package main

import (
	"service/pagination/service"
)

func main() {
	//service.Paginationservice(4, 5, 5, 1, 0)
	//service.Paginationservice(4, 5, 1000, 15, 100)
	service.Paginationservice(12, 5, 20, 2, 2)
	service.Paginationservice(4, 5, 10, 2, 2)

}
