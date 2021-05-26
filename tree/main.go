package main

import (
	"encoding/json"
	"fmt"
)

// Entry
type Entry struct {
	ID       int     // 主键
	PID      int     // 上级主键
	Children []Entry // 子列表
}

func ToTree(items []Entry, pid int) []Entry {
	// 将所有pid相同的放入map的数组中
	mp := make(map[int][]Entry)
	for _, v := range items {
		mp[v.PID] = append(mp[v.PID], v)
	}

	// 确定顶级
	root, ok := mp[pid]
	if !ok {
		return []Entry{}
	}
	for i := 0; i < len(root); i++ {
		root[i] = deepChildren(root[i], mp)
	}
	return root
}

func deepChildren(item Entry, mp map[int][]Entry) Entry {
	children, ok := mp[item.ID]
	if ok { // 确定此父级是否有子项,有则递归子项
		for i := 0; i < len(children); i++ {
			children[i] = deepChildren(children[i], mp)
		}
	} else {
		children = []Entry{}
	}
	item.Children = children
	return item
}

func main() {
	var depts = []Entry{
		{ID: 1, PID: 0},
		{ID: 2, PID: 1},
		{ID: 3, PID: 1},
		{ID: 4, PID: 2},
		{ID: 5, PID: 2},
		{ID: 6, PID: 3},
		{ID: 7, PID: 3},
		{ID: 8, PID: 0},
		{ID: 9, PID: 8},
		{ID: 10, PID: 8},
		{ID: 11, PID: 9},
		{ID: 12, PID: 9},
		{ID: 13, PID: 10},
		{ID: 14, PID: 10},
	}
	v, _ := json.MarshalIndent(ToTree(depts, 0), "  ", "  ")
	fmt.Printf("%+v", string(v))
}
