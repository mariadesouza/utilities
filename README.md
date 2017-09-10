# utilities

# Get Package

go get github.com/mariadesouza/utilities/structutil

# Usage
type testStruct struct {
	Name       string
	Email      []string
	Occupation string
}

	testStruct := testStruct{"Ethan", []string{"emdesouza@gmail.com"}, "engineer"}
  m := structutil.StructToMap(&testStruct)
	for key, val := range *m {
		fmt.Println(key, val)
	}
