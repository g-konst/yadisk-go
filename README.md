Golang client for the [Yandex Disk API](https://yandex.com/dev/disk-api/doc/en/)

## Usage

```Golang
import (
    "os"
    "fmt"

    "github.com/g-konst/yadisk-go"
)

func main() {
	token := os.Getenv("YANDEX_DISK_TOKEN")

	client := yadisk.NewYandexDiskClient(token)
	ctx := context.Background()

    disk, resp, err := client.Disk.Get(ctx, nil)
    fmt.Println("User name:", disk.User.DisplayName)
}
```
