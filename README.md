# repro-client

## Repro client library in Go

ref. http://docs.repro.io/ja/dev/user-profile-api/index.html#user-profiles-payload

## Example

```
package main

import (
	"time"

	"github.com/kentokento/repro-client"
)

func init() {
	repro.Setup("API-TOKEN")
}

func main() {
	userProfiles := repro.NewUserProfiles(123456789)

	userProfiles.AddString("国籍", "日本")
	userProfiles.AddString("会員ステータス", "通常会員")
	userProfiles.AddString("課金経験の有無", "有り")
	userProfiles.AddInt("生まれた年", 1989)
	userProfiles.AddInt("生まれた月", 7)
	userProfiles.AddDecimal("アクティブ率", 60.345)
	userProfiles.AddDatetime("登録日時", time.Now())

	resp, err := userProfiles.Send()
	if err != nil {
		// an error has occurred
	}
}
```
