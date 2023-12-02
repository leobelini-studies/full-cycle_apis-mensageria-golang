package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/leobelini-studies/full-cycle_apis-mensageria-golang/internal/usecase"
)

type ProductHandlers struct {
	CreateProductUseCase *usecase.CreateProductUseCase
	ListProductsUseCase  *usecase.ListProductsUseCase
}

func NewProductHandlers(
	createProductUseCase *usecase.CreateProductUseCase,
	listProductsUseCase *usecase.ListProductsUseCase,
) *ProductHandlers {
	return &ProductHandlers{
		CreateProductUseCase: createProductUseCase,
		ListProductsUseCase:  listProductsUseCase,
	}
}

func (h *ProductHandlers) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	output, err := h.CreateProductUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(output)
}

func (h *ProductHandlers) ListProducts(w http.ResponseWriter, r *http.Request) {
	output, err := h.ListProductsUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(output)
}
