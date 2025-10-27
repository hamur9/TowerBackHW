package main

import (
	"errors"
	"fmt"
)

type Tree struct {
	Head *Node
}
type NodeKeyValue struct {
	Key   int
	Value string
}

type Node struct {
	Data  *NodeKeyValue
	Left  *Node
	Right *Node
}

func main() {
	var count int
	tree := Tree{nil}
	fmt.Print("Введите количесто вводимых элементов:\n> ")
	_, err := fmt.Scan(&count)
	if err != nil {
		errors.New("ошибка ввода количества элементов")
	}

	fmt.Print("Введите пары ключ-значение (например '5 Москва')\n> ")
	for i := 0; i < count; i++ {
		var newNode NodeKeyValue
		_, err = fmt.Scan(&newNode.Key, &newNode.Value)
		if err != nil {
			errors.New("ошибка ввода пары ключ-значение\n")
		} else if i != count-1 {
			fmt.Print("> ")
		}
		tree.Head = Insert(tree.Head, &newNode)
	}

	fmt.Printf("Максимальная глубина дерева: %d\n", Depth(tree.Head))

	var keyToFind int
	fmt.Print("Введите ключ, который хотите найти:\n> ")
	_, err = fmt.Scan(&keyToFind)
	if err != nil {
		errors.New("ошибка считывания ключа")
	}
	nodeAddr := Find(tree.Head, keyToFind)
	if nodeAddr == nil {
		fmt.Printf("Не существует записи с ключом %d.\n", keyToFind)
	} else {
		fmt.Printf("Запись, найденная по ключу %d: '%s'\n", keyToFind, nodeAddr.Data.Value)
	}

	var keyToDelete int
	fmt.Print("Введите ключ, который хотите удалить:\n> ")
	_, err = fmt.Scan(&keyToDelete)
	if err != nil {
		errors.New("Ошибка ввода искомого ключа")
	}
	if Find(tree.Head, keyToDelete) == nil {
		fmt.Print("Элемента нет в дереве.\n")
	} else {
		tree.Head = Remove(tree.Head, keyToDelete)
		if Find(tree.Head, keyToDelete) == nil {
			fmt.Print("Элемент успешно удален.\n")
			fmt.Printf("Новая максимальная глубина дерева: %d\n", Depth(tree.Head))
		}
	}
}

func Insert(head *Node, newNode *NodeKeyValue) *Node {
	if head == nil {
		return &Node{newNode, nil, nil}
	}

	if newNode.Key > head.Data.Key {
		head.Right = Insert(head.Right, newNode)
	} else if newNode.Key < head.Data.Key {
		head.Left = Insert(head.Left, newNode)
	} else {
		head.Data.Value = newNode.Value
	}
	return head
}

func Remove(head *Node, key int) *Node {
	if head == nil {
		return nil
	}
	if key > head.Data.Key {
		head.Right = Remove(head.Right, key)
	} else if key < head.Data.Key {
		head.Left = Remove(head.Left, key)
	} else {
		if head.Left == nil {
			return head.Right
		} else if head.Right == nil {
			return head.Left
		} else {
			minRight := findMinRight(head.Right)
			head.Data.Key = minRight.Data.Key
			head.Data.Value = minRight.Data.Value
			head.Right = Remove(head.Right, minRight.Data.Key)
		}
	}
	return head
}

func findMinRight(head *Node) *Node {
	if head == nil {
		return nil
	}
	for head.Left != nil {
		head = head.Left
	}
	return head
}

func Find(head *Node, key int) *Node {
	if head == nil {
		return nil
	}
	if key > head.Data.Key {
		return Find(head.Right, key)
	} else if key < head.Data.Key {
		return Find(head.Left, key)
	} else {
		return head
	}
}

func Depth(head *Node) int {
	if head == nil {
		return 0
	}
	return max(Depth(head.Left), Depth(head.Right)) + 1
}
