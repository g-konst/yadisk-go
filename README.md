Golang client for the [Yandex Disk API](https://yandex.com/dev/disk-api/doc/en/)

## Usage

```Golang
import (
    "os"
    "fmt"
    "context"

    yadisk "github.com/g-konst/yadisk-go/client"
)

func main() {
    token := os.Getenv("YANDEX_DISK_TOKEN")

    client := yadisk.NewYandexDiskClient(token)
    ctx := context.Background()

    // disk
    disk, resp, err := client.Disk.Get(ctx, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println("Response:", resp.Status)
    fmt.Println("User name:", disk.User.DisplayName)

    // resources
    resource, resp, err := client.Resources.Get(
        ctx, yadisk.ResourcesGetOpts{Path: "/"},
    )
    link, resp, err := client.Resources.Create(
        ctx, yadisk.ResourcesCreateOpts{Path: "/test"},
    )
    link, resp, err := client.Resources.Delete(
        ctx, yadisk.ResourcesDeleteOpts{Path: "/test"},
    )
    link, resp, err := client.Resources.Move(
        ctx, yadisk.ResourcesMoveOpts{From: "/From", Path: "/To"},
    )
    link, resp, err := client.Resources.Copy(
        ctx, yadisk.ResourcesMoveOpts{From: "/From", Path: "/To"},
    )
    filesRes, resp, err := client.Resources.GetFiles(
        ctx, &yadisk.ResourcesGetFilesOpts{Limit: 10},
    )
    link, resp, err := client.Resources.GetDownloadLink(
        ctx, yadisk.ResourcesGetDownloadLinkOpts{Path: "/test"},
    )
    link, resp, err := client.Resources.Publish(
        ctx, yadisk.ResourcesPublishOpts{Path: "/test"},
    )

    // trash
    trash, resp, err := client.Trash.Get(
        ctx, yadisk.TrashGetOpts{Path: "/",},
    )
    link, resp, err := client.Trash.Clear(ctx, nil)

    // operations
    status, resp, err := client.Operations.Get(
        ctx, yadisk.OperationsGetOpts{OperationId: "test"},
    )
}
```
