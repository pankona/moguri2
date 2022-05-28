package main

import (
	"fmt"
	"strings"
)

type Room struct {
	id   string
	next []*Room
}

func newRoom(id string) *Room {
	return &Room{id: id}
}

// 次の行き先の配列を返す。配列は何らかでソートされている。
func (r *Room) GetNextRooms() ([]*Room, error) {
	return r.next, nil
}

// 次の行き先を足す。足せる数には上限があるかもしれない。
func (r *Room) AddNextRoom(next *Room) error {
	r.next = append(r.next, next)
	return nil
}

// 行き止まりかどうかを返す。典型的にはゴールの判定に使う？
func (r *Room) IsDeadEnd() bool {
	return len(r.next) == 0
}

func (r *Room) String() string {
	return fmt.Sprintf("%s: %s", r.id, strings.Join(names(r.next), ","))
}

func names(in []*Room) []string {
	ret := make([]string, 0, len(in))
	for _, v := range in {
		ret = append(ret, v.id)
	}
	return ret
}
