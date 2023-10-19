package grpc

import (
	"context"

	"github.com/emiliosheinz/imersao-full-cycle-challenge-01/application/grpc/pb"
	"github.com/emiliosheinz/imersao-full-cycle-challenge-01/application/usecase"
)

type ProductGrpcService struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func (productGrpcService *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := productGrpcService.ProductUseCase.CreateProduct(in.Name, in.Description, in.Price)

	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
		},
	}, nil
}

func (productGrpcService *ProductGrpcService) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	products, err := productGrpcService.ProductUseCase.FindProducts()

	if err != nil {
		return nil, err
	}

	var productsResponse []*pb.Product

	for _, product := range *products {
		productsResponse = append(productsResponse, &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
		})
	}

	return &pb.FindProductsResponse{
		Products: productsResponse,
	}, nil
}

func NewProductGrpcService(usecase usecase.ProductUseCase) *ProductGrpcService {
	return &ProductGrpcService{
		ProductUseCase: usecase,
	}
}
