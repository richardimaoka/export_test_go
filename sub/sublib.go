package sub

import "fmt"

func ExportedAdd(a, b int) {
	fmt.Printf("ExportedAdd: %d + %d = %d\n", a, b, a+b)
}

func unExportedAdd(a, b int) {
	fmt.Printf("unExportedAdd: %d + %d = %d\n", a, b, a+b)
}
