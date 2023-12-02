package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/leobelini-studies/full-cycle_apis-mensageria-golang/internal/infra/akafka"
	"github.com/leobelini-studies/full-cycle_apis-mensageria-golang/internal/infra/repository"
	"github.com/leobelini-studies/full-cycle_apis-mensageria-golang/internal/infra/web"
	"github.com/leobelini-studies/full-cycle_apis-mensageria-golang/internal/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")
	if err != nil {
		fmt.Println("Erro ao conectar no MySQL")
		fmt.Println(err.Error())
		panic(err)
	}

	defer db.Close()

	repository := repository.NewProductRepositoryMySQL(db)
	createProductUsecase := usecase.NewCreateProductUseCase(repository)
	listProductsUsecase := usecase.NewListProductsUseCase(repository)

	productHandlers := web.NewProductHandlers(createProductUsecase, listProductsUsecase)

	r := chi.NewRouter()
	r.Post("/product", productHandlers.CreateProduct)
	r.Get("/products", productHandlers.ListProducts)

	go http.ListenAndServe(":8081", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"product"}, "kafka:29092", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			fmt.Println("Erro ao deserializar o product")
			fmt.Println(err.Error())
			continue
		}

		_, errCreate := createProductUsecase.Execute(dto)

		if errCreate != nil {
			fmt.Println("Erro ao criar o product")
			fmt.Println(err.Error())
			continue
		}

	}
}
