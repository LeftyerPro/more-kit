# more-kit

A tiny Go kit for **lossless** image compression (PNG/JPEG → WebP) and utility helpers.

## Install

```bash
go install github.com/LeftyerPro/more-kit/cmd/morekit@latest

# version
morekit -v

# compress 1.png → 1.webp (lossless)
morekit -i 1.png -o 1.png -t 1
morekit -i 1.jpg -o 1.jpg -t 2
morekit -i 1.png -o 1.webp -t 3

import "github.com/LeftyerPro/more-kit/pkg/morekit"

err := morekit.CompImage("in.png", "out.png", 1)
err := morekit.CompImage("in.jpg", "out.jpg", 2)
err := morekit.CompImage("in.png", "out.webp", 3)

| Op      | Time      | Mem      |
| ------- | --------- | -------- |
| CompImg | ~44 ms/op | ~2 MB/op |

MIT – see [LICENSE](LICENSE) for details.
