## Binary Tree 前序非递归

``` go

type Node struct {
    Val int
    Left *Node
    Right *Node
}



func Tranverse(root *Node) []int {
    if root == nil {
        return nil
    }

    stack := make([]*Node,0)
    result := make([[]int,0])

    for len(stack) != 0 || root != nil {
        for root != nil {
            result = append(result, root.Val)
            stack = append(stack, root)
            root = root.Left
        }
        // reach the bot left
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        root = root.Right

    }
    return result
}

```


## Binary Tree 层次遍历


```go

func Tranverse(root *Node) []int {

    if root == nil {
        return nil
    }

    queue := make([]*Node,0)
    queue = append(queue, root)

    result := make([][]int,0)

    for len(queue) != 0 {
        
        // len of the layer
        l := len(queue)
        list := make([]int, l)
        // enqueue

        for i:=0; i < l; i++ {
            node := queue[i]
            queue = queue[1:]
            list[i] = node.Val
           
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result := append(result, list)
    
    }
    return result
}

```


## Nearest Common Ancestor



```go 

type Node struct {
    Val int
    Left *Node
    Right *Node
}

type Result struct {
    Have bool
    Ances *Node
}



func NearestAncestor(target1, target2 int, root *Node) *Node {







}


func find (target1, target2 int, root *Node) Result {
    
    if root.Val == target1 || root.Val == target2 {
        return Result{
            Have: true,
        }
    }

    lresult := find(target1, target2, root.Left)
    rresult := find(target1, target2, root.Right)

    if lresult.Have && rresult.Have {
        return Result {
            Have: true,
            Ances: root,
        }
    }

    if lresult.Have || rresult.Have {
        if lresult.Ances != nil {
            return Result {
                Have: true,
                Ances: lresult.Ances
            }
        }

        if rresult.Ances != nil {
            return Result {
                Have: true,
                Ances: rresult.Ances
            }

        }

    }


}


```

