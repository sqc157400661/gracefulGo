package must

type ForTest struct{
	a string
}
func retrunWithVal() (a ForTest){
	a = ForTest{a:"sdafsafsafsadfsafsadfsadfasdfsdafafa"}
	return a
}

func retrunWithPoint() (a *ForTest){
	a = &ForTest{a:"sdafsafsafsadfsafsadfsadfasdfsdafafa"}
	return a
}