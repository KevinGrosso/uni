mkdir %1
cd %1
(
    echo package main
    echo.
    echo import ^(
    echo     "fmt"
    echo ^)
    echo.
    echo func main^(^) ^{
    echo ^}
) > %1.go
