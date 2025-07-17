// cmd/blockchainfabriccomposer/main.go
package main

import (
"flag"
"log"
"os"

"blockchainfabriccomposer/internal/blockchainfabriccomposer"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainfabriccomposer.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
