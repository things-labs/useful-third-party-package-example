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
	// 递归顶级确定所有子项
	deepChildren(root, mp)
	return root
}

func deepChildren(items []Entry, mp map[int][]Entry) {
	for i := 0; i < len(items); i++ { // 确定此父级的子项,从map中查找
		children, ok := mp[items[i].ID]
		if ok { // 确定此父级是否有子项,有则递归子项
			deepChildren(children, mp)
		} else {
			children = []Entry{}
		}
		items[i].Children = children
	}
}

func toDeptTree1(items []Entry) []Entry {
	tree := make([]Entry, 0)
	for _, itm := range items {
		if itm.PID == 0 {
			tree = append(tree, deepChildrenDept1(items, itm))
		}
	}
	return tree
}

func deepChildrenDept1(items []Entry, item Entry) Entry {
	item.Children = make([]Entry, 0)
	for _, itm := range items {
		if itm.PID == item.ID {
			item.Children = append(item.Children, deepChildrenDept1(items, itm))
		}
	}
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
