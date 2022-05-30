package main

func (t *TreeNode) insertIntoBST(value int) *TreeNode {
	if t == nil {
		return CreateTreeNode(value)
	}

	if t.Value == value {
		return t
	}

	//根结点比插入值大
	if t.Value > value {
		t.Left = t.Left.insertIntoBST(value)
	}else if t.Value < value {
		t.Right = t.Right.insertIntoBST(value)
	}
	return t
}