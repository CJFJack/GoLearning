package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const TimeLayout = "2006-01-02 15:04:05"

type User struct {
	ID       int    `inputJson:"id"`
	Name     string `inputJson:"name"`
	Password string `inputJson:"-"`
	Sex      bool   `inputJson:"sex,string"`
	Birthday Birthday
}

type Birthday time.Time

func (b *Birthday) UnmarshalText(txt []byte) error {
	if txt == nil {
		*b = Birthday(time.Unix(0, 0))
	}
	if t, err := time.Parse(TimeLayout, string(txt)); err != nil {
		return err
	} else {
		*b = Birthday(t)
		return nil
	}
}

func (b Birthday) MarshalText() ([]byte, error) {
	return []byte(time.Time(b).Format(TimeLayout)), nil
}

func (b Birthday) String() string {
	return time.Time(b).Format(TimeLayout)
}

func main() {
	now := Birthday(time.Now())
	users := []User{
		{1, "lulu", "123", false, now},
		{2, "cjf", "321", true, now},
		{3, "yueyue", "456", false, now},
	}

	bytes, _ := json.MarshalIndent(users, "", "\t")
	fmt.Println(string(bytes))

	var birthday Birthday = Birthday(now)
	ctx, _ := json.Marshal(birthday)
	fmt.Println(string(ctx))

	var bt Birthday
	fmt.Println(json.Unmarshal([]byte(`"2020-11-25 11:13:22"`), &bt))
	fmt.Println(time.Time(bt))
	fmt.Println(bt)
}
