# paging

Paging is a simple Go package that helps you manage pagination in your projects. It provides an easy-to-use interface to handle page indexing, page size, total items, and the items themselves.

## Features

- Calculate the total number of pages based on total items and page size.
- Store pagination details such as current page index, page size, total pages, and total items.
- Easily integrate with your data models.

## Installation

To install the package, you can use `go get`:

```sh
go get github.com/pupilcc/paging
```

## Usage

Below is an example of how to use the Paginator in your Go project:

```go
package main

import (
    "fmt"
    "github.com/pupilcc/paging"
)

func main() {
    // Example data
    items := []string{"item1", "item2", "item3", "item4", "item5"}
    totalItems := len(items)
    pageIndex := 1
    pageSize := 2

    // Create a new Paginator
    paginator := paging.NewPaginator(pageIndex, pageSize, totalItems, items)

    // Print pagination details
    fmt.Printf("Page Index: %d\n", paginator.PageIndex)
    fmt.Printf("Page Size: %d\n", paginator.PageSize)
    fmt.Printf("Total Pages: %d\n", paginator.Pages)
    fmt.Printf("Total Items: %d\n", paginator.Total)
    fmt.Printf("Items: %v\n", paginator.Items)
}
```

## Paginator Struct

The `Paginator` struct contains the following fields:

- `PageIndex`: The current page index.
- `PageSize`: The number of items per page.
- `Pages`: The total number of pages.
- `Total`: The total number of items.
- `Items`: The items for the current page.

## Functions

### NewPaginator

Creates a new Paginator instance.

```go
func NewPaginator(pageIndex, pageSize, total int, items interface{}) *Paginator
```

- `pageIndex`: The current page index.
- `pageSize`: The number of items per page.
- `total`: The total number of items.
- `items`: The items for the current page.

### getPages

Calculates the total number of pages.

```go
func getPages(total, pageSize int) int
```

- `total`: The total number of items.
- `pageSize`: The number of items per page.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Please feel free to submit issues, fork the repository and send pull requests!

## Acknowledgements

This project was inspired by the need for a simple and efficient pagination solution in Go.

---

Happy coding!