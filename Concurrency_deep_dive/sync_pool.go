package mac

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age  int
}

func SyncPool() {
	var pool = sync.Pool{
		// New: func() interface{} {
		// 	fmt.Println("Creating a new Person.")
		// 	return &person{}
		// },
	}
	pool.Put(&person{name: "John", age: 81})

	// Get an object from the pool
	person1 := pool.Get().(*person)
	// person1.name = "John"
	// person1.age = 18
	// fmt.Println("Got Person:", person1)

	fmt.Printf("Person1 - Name: %s, Age: %d\n", person1.name, person1.age)

	pool.Put(person1)
	fmt.Println("Returned Person to pool")

	person2 := pool.Get().(*person)
	fmt.Println("Got person2:", person2)

	// person3 := pool.Get().(*person)
	person3 := pool.Get()
	if person3 != nil {
		fmt.Println("Got person3:", person3)
		person3.(*person).name = "Jane"
		pool.Put(person3)
	} else {
		fmt.Println("Sync Pool is empty")
	}

	// Returning object to the pool again
	pool.Put(person2)
	fmt.Println("Returned Persons to pool")

	person4 := pool.Get().(*person)
	fmt.Println("Got person4: ", person4)

	person5 := pool.Get().(*person)
	fmt.Println("Got person5:", person5)

}
