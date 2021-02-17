# Text Ticker

For ticking text like the [marquee](https://en.wikipedia.org/wiki/Marquee_element) tag.

## Usage

```go
package main

import (
        "log"

        tt "github.com/meinside/text-ticker"
)

func main() {
        ticker := tt.NewTextTicker("동해물과 백두산이 마르고 닳도록!", 7)

        for i:=0; i < 100; i++ {
                log.Printf("%d: [%s]", i+1, ticker.NextText())
        }
}

```

Output:

```
// TODO
```

## License

MIT

