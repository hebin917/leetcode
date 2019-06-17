package main



func SearchHash(hash []int,hashlen int,key int) int {
	address := key % hashlen

	for hash[address] !=0 && hash[address] !=key {
		address ++
		address = address % hashlen
	}
	if hash[address] == 0 {
		return -1
	}
	return address
}


func InsertHash(hash []int ,hashlen int ,data int) {
	address := data %hashlen

	for hash[address] !=0 {
		address ++
		address= address % hashlen
	}
	hash[address] = data
}