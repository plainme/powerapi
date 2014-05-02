package filewalker

import (
  "path/filepath"
  "os"
  "flag"
  "fmt"
)

func visit(path string, f os.FileInfo, err error) error {
  fmt.Printf("Visited: %s\n", path)
  return nil
} 

func Doit(path string) {
  //flag.Parse()
  //root := flag.Arg(0)
  root := path
  fmt.Println("testtest")
  fmt.Println("und root ist " + root)
  if (root == "") {
  	root = "/home/sascha/"
  }
  
  err := filepath.Walk(root, visit)
  fmt.Printf("filepath.Walk() returned %v\n", err)
}

func main() {
  flag.Parse()
  root := flag.Arg(0)
  fmt.Println("testtest")
  fmt.Println("und root ist " + root)
  if (root == "") {
  	root = "/home/sascha/"
  }
  
  err := filepath.Walk(root, visit)
  fmt.Printf("filepath.Walk() returned %v\n", err)
}