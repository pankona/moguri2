package main

/*
## ダンジョンの構造

   a (start)
 / | \
/  |  \
b  c  d
|\/|\/|
|/\|/\|
e  f  g
|\/|\/|
|/\|/\|
h  i  j
|\/|\/|
|/\|/\|
k  l  m
\  |  /
 \ | /
   n (goal)

各部屋を Room と呼ぶ
Room は 0 個以上の行き先を持つ (ただし、行き先が 0 個なのはゴールのみ)

Room の関連性は JSON で表すことができる

{
  a: [b, c, d],
  b: [e, f],
  c: [e, f, g],
  d: [f, g]},
  e: [h, i]},
  f: [h, i, j],
  g: [i, j],
  h: [k, l],
  i: [k, l, m],
  j: [l, m],
  k: [n],
  l: [n],
  m: [n],
  n: [],
}

*/

type Dungeon struct {
	start      *Room
	roomByName map[string]*Room
}

func ConstructDungeon(dungeon map[string][]string, start string) (*Dungeon, error) {
	// TODO: input のバリデーションをする
	// - ゴールがないといけないので、最後は next が nil になっている必要がある
	// - 他にもあるか…？

	m := map[string]*Room{}
	for k := range dungeon {
		m[k] = newRoom(k)
	}

	for k, v := range dungeon {
		for _, n := range v {
			m[k].AddNextRoom(m[n])
		}
	}

	return &Dungeon{
		start:      m[start],
		roomByName: m,
	}, nil
}

func (d *Dungeon) Start() *Room {
	return d.start
}

func (d *Dungeon) GetByName(name string) (*Room, bool) {
	r, ok := d.roomByName[name]
	return r, ok
}
