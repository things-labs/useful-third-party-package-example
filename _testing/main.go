package main

import (
	"fmt"
)

// Dept 部门
type Dept struct {
	ID       int    // 主键
	PID      int    // 上级主键
	Children []Dept // 子列表
}

func BuildDeptTree(ds []Dept) []Dept {
	mp := make(map[int][]Dept)
	for _, v := range ds {
		mp[v.PID] = append(mp[v.PID], v)
	}
	tree, ok := mp[0]
	if !ok {
		return tree
	}
	deepChildrenDept(tree, mp)
	return tree
}

func deepChildrenDept(item []Dept, mp map[int][]Dept) {
	for i := 0; i < len(item); i++ {
		children, ok := mp[item[i].ID]
		if !ok {
			continue
		}
		deepChildrenDept(children, mp)
		item[i].Children = children
	}
}

func toDeptTree1(items []Dept) []Dept {
	tree := make([]Dept, 0)
	for _, itm := range items {
		if itm.PID == 0 {
			tree = append(tree, deepChildrenDept1(items, itm))
		}
	}
	return tree
}

func deepChildrenDept1(items []Dept, item Dept) Dept {
	item.Children = make([]Dept, 0)
	for _, itm := range items {
		if itm.PID == item.ID {
			item.Children = append(item.Children, deepChildrenDept1(items, itm))
		}
	}
	return item
}

func main() {
	var depts = []Dept{
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

	fmt.Printf("%+v", BuildDeptTree(depts))
}
