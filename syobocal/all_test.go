package syobocal

import (
	"fmt"
	"testing"
	"time"

	. "github.com/otiai10/mint"
)

func TestClient(t *testing.T) {
	Expect(t, true).ToBe(true)
	client := NewClient()
	Expect(t, client).TypeOf("*syobocal.Client")
}

func TestClient_TitleLookup(t *testing.T) {
	Expect(t, true).ToBe(true)
	client := NewClient()
	res, err := client.TitleLookup(time.Now(), time.Now())

	fmt.Printf("%+v\n", res)

	Expect(t, res.Result.Code).ToBe(200)
	Expect(t, err).ToBe(nil)
}
