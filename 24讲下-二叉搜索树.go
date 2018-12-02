package main

func NewTreeNode(v int) *TreeNode {
	return &TreeNode{value: v}
}

func (this *TreeNode) AddNode(v int) {
	if nil == this {
		return
	}
	if v >= this.value {
		this.rchild = NewTreeNode(v)
	} else {
		this.lchild = NewTreeNode(v)
	}
}

func (this *TreeNode) Find(v int) {
	if nil == this {
		return
	}
	if v > this.value {
		return this.rchild.Find(v)
	} else if v < this.value {
		return this.lchild.Find(v)
	} else {
		return this.value
	}
}

func (this *TreeNode) Insert(v int) {
	if v >= this.value {
		if nil != this.rchild {
			return this.rchild.Insert(v)
		} else {
			this.rchild = NewTreeNode(v)
			return
		}
	} else {
		if nil != this.lchild {
			return this.lchild.Insert(v)
		} else {
			this.lchild = NewTreeNode(v)
			return
		}
	}
}

func (this *TreeNode) Delete(v int) {
	n := this
	var deleteNode *TreeNode
	var deleteNodeFather *TreeNode
	for nil != n && v != n.value {
		deleteNodeFather = n
		if v > n.value {
			n = n.rchild
		} else {
			n = n.lchild
		}
	}
	deleteNode = n
	if nil == deleteNode {
		return // not found node
	}

	if nil != deleteNode.lchild && nil != deleteNode.rchild {
		// exsit both left and right node, find min node from right
		minN := deleteNode.rchild
		minNP := deleteNode
		for nil != minN.lchild {
			minNP = minN
			minN = minN.lchild
		}
		deleteNode.value = minN.value // replace value of deleteNode
		deleteNode = minN // need to delete minN
		deleteNodeFather = minNP
	}
	// delete node has no children or one child
	var child *TreeNode
	if nil != deleteNode.lchild {
		child = deleteNode.lchild
	} else if nil != deleteNode.rchild {
		child = deleteNode.rchild
	} else {
		child = null;
	}

	if nil == deleteNodeFather {
		this = child
	} else if deleteNodeFather.lchild == p {
		deleteNodeFather.lchild = child
	} else {
		deleteNodeFather.rchild = child
	}

}
