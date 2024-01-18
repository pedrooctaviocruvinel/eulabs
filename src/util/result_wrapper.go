package util

type ResultWrapper[T any] struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
	Data    T        `json:"data"`
}

func NewResultWrapperFailed[T any](e []string) ResultWrapper[T] {
	return ResultWrapper[T]{
		Success: false,
		Errors:  e,
	}
}

func NewResultWrapperSucceded[T any](d T) ResultWrapper[T] {
	return ResultWrapper[T]{
		Success: true,
		Errors:  nil,
		Data:    d,
	}
}

func NewResultWrapperSuccededEmpty[T any]() ResultWrapper[T] {
	return ResultWrapper[T]{
		Success: true,
		Errors:  nil,
	}
}
