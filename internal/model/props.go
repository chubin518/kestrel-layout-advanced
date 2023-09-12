package model

type TreeNode struct {
	Value    string      `json:"value"`
	Label    string      `json:"label"`
	Children []*TreeNode `json:"children"`
}
