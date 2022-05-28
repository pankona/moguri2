package main

import "fmt"

var route = map[string][]string{
	"a": {"b", "c", "d"},
	"b": {"e", "f"},
	"c": {"e", "f", "g"},
	"d": {"f", "g"},
	"e": {"h", "i"},
	"f": {"h", "i", "j"},
	"g": {"i", "j"},
	"h": {"k", "l"},
	"i": {"k", "l", "m"},
	"j": {"l", "m"},
	"k": {"n"},
	"l": {"n"},
	"m": {"n"},
	"n": {},
}

func main() {
	d, err := ConstructDungeon(route, "a")
	if err != nil {
		fmt.Printf("failed to construct dungeon: %v", err)
	}

	// 最初から最後までとりあえず出力してみる
	current := d.Start()
	for !current.IsDeadEnd() {
		fmt.Println(current)
		current = current.next[0]
	}
	fmt.Println(current)

	// 途中から再開するときはこのような
	fmt.Println(d.GetByName("i"))

	r := &SampleRoom{}
	r.Show(r.InitialState())
}
