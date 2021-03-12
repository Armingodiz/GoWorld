package futures

/*
Sometimes you know you need to compute a value before you need to actually use the value.
In this case, you can potentially start computing the value on another processor and have it ready when you need it.
This is the idea behind futures.
*/
type Matrix [][]int

func InverseProduct(a Matrix, b Matrix) Matrix {
	a_inv_future := InverseFuture(a)
	b_inv_future := InverseFuture(b)
	// using channels in this part is for stopping product to be executed before inversion of a & b .
	a_inv := <-a_inv_future
	b_inv := <-b_inv_future
	return Product(a_inv, b_inv)
}

func InverseFuture(a Matrix) chan Matrix {
	future := make(chan Matrix)
	go func() { future <- Inverse(a) }()
	return future
}

// dummy functions :
func Product(a, b Matrix) Matrix {
	return a
}
func Inverse(a Matrix) Matrix {
	return a
}
