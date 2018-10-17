# OpenWeatherMap API Client for Go

```go
package main

import (
  "github.com/otiai10/openweathermap"
)
func main() {
  client := openweathermap.New(API_KEY)
  res, _ := client.ByCityName("Tokyo")
  fmt.Println(res.Forecasts[0].Weather[0].Main)
  // Clouds
}
```
