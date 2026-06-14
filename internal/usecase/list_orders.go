package usecase

import "github.com/JacksomGuilherme/Desafio-CleanArch/internal/entity"

type ListOrdersOutputDTO struct {
	Orders []entity.Order `json:"oreders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() (ListOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}

	dto := ListOrdersOutputDTO{
		Orders: orders,
	}

	return dto, nil
}
