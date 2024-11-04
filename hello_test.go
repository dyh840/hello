// hello_test.go
package hello_test

import (
        "testing"
        "fmt"
        "github.com/dyh840/hello"
)

func TestHello(t *testing.T) {
        data := "jack"
        expected := fmt.Sprintf("hello %s!", data)
        result := hello.Hello(data)

        if result != expected {
                t.Fatalf("expected result %s, but got %s", expected, result)
        }

}
