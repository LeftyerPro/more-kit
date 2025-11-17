 # more-kit 
 A tiny Go kit for lossless image compression (PNG/JPEG -> WebP) and common dev-ops utilities. 

 ## Install 
 ```bash 
 go install github.com/LeftyerPro/more-kit/cmd/morekit@latest 
 ``` 

 ## Quick Start 
 ```bash 
 # version 
 morekit -v 

 # compress 1.png -> 1.webp (lossless) 
 morekit comp -i 1.png -o 1.webp -t 3 

 # device info (defaults to -info) 
 morekit device 
 morekit device -id -name 
 morekit device -cpu -mem -disk 
 morekit device -ips -ip 
 ``` 

 ## CLI Reference 
 | Command  | Flags                                                                 | Description                       | 
 |----------|-----------------------------------------------------------------------|-----------------------------------| 
 | device   | -id -name -info -boot -uuid -cpu -cores -mem -disk -ips -macs -ip    | hardware / network info           | 
 | comp     | -i <in> -o <out> -t <1|2|3>                                           | 1=PNG->PNG 2=JPEG->JPEG 3->WebP   | 
 | folder   | -c copy|exist -s <src> [-d <dst>] [-clear 0|1]                        | folder copy or exist check        | 
 | file     | -c copy|exist|read|write -s <src> [-d <dst>] [-t <text>] [-p <path>]  | file utilities                    | 
 | json     | -c read|write -p <path>                                               | read / write json                 | 

 ## Library Usage 
 ```go 
 import "github.com/LeftyerPro/more-kit/pkg/morekit" 

 // lossless compress 
 err := morekit.CompImage("in.png", "out.webp", 3) 

 // device info 
 id   := morekit.DeviceGetId() 
 name := morekit.DeviceGetName() 
 info := morekit.DeviceGetInfo() 
 ``` 

 ## Benchmark 
 | Op      | Time      | Mem      | 
 |---------|-----------|----------| 
 | CompImg | ~44 ms/op | ~2 MB/op | 

 ## License 
 MIT – see [LICENSE](LICENSE) for details. 