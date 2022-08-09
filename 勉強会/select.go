func main()  {
	gen1, gen2 := make(chan, int), make(chan int)

	select {
	case num := <-gen1:
		fmt.Println
		
	}
}