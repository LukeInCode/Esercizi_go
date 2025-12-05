package main
import "fmt"


//Doppia lista concatenata
type DListNode struct {
    Value string
    Prev *DListNode
    Next *DListNode
}

func (first *DListNode) AddFirst(v string) *DListNode {
  var newNode *DListNode = new(DListNode);

  newNode.Value = v;
  newNode.Next = first;
  
  first.Prev = newNode;

  return newNode;
}


func (first *DListNode) AddLast(v string) *DListNode {
  newNode := new(DListNode);
  newNode.Value = v;
  
  if first == nil {
    return newNode;
  }
  
  var cursor *DListNode;
  for cursor = first; cursor.Next != nil; cursor = cursor.Next {
  }
  cursor.Next = newNode;
  newNode.Prev = cursor;

  return first;
}

func (first *DListNode) String() (list string) {
  var cursor *DListNode;
  for cursor = first; cursor.Next != nil; cursor = cursor.Next {
    list += fmt.Sprintf("%s <-> ", cursor.Value);
  }
  list += fmt.Sprintf("%s", cursor.Value);
  return;
}


func (first *DListNode) Remove(v string) *DListNode {
  newFirst := first;
  for cursor := first; cursor != nil; cursor = cursor.Next {
    if cursor.Value == v {
      if cursor.Prev == nil {
        newFirst = cursor.Next;
        break;
      }
      if cursor.Next == nil {
        (cursor.Prev).Next = nil;
        break;
      }
      (cursor.Prev).Next = cursor.Next;
      (cursor.Next).Prev = cursor.Prev;
      break;
    }
  }
  return newFirst;
}

func (first *DListNode) Len() (len int) {
  for cursor := first; cursor != nil; cursor = cursor.Next {
    len++;
  }
  return len;
}


//Stack (Implementato come insieme)

type UniqueStack struct {
  Value string;
  Next *UniqueStack;
}

func (top *UniqueStack) Push(v string)  *UniqueStack {
  newTop := new(UniqueStack);
  if top == nil {
    newTop.Value = v;
    return newTop;
  }
  var flag bool = true;
  for cursor := top; cursor != nil; cursor = cursor.Next {
    if cursor.Value == v {
      flag = false;
      break;
    }
  }
  if flag {
    newTop.Value = v;
    top.Next = newTop;
  }
  return top;
}

func (top *UniqueStack) pop() *UniqueStack {
  return top.Next;
}

func (top *UniqueStack) Len() (len int) {
  for cursor := top; cursor != nil; cursor = cursor.Next {
    len++;
  }
  return len;
}

func (top *UniqueStack) String() (stack string) {
  var cursor *UniqueStack;
  for cursor = top; cursor.Next != nil; cursor = cursor.Next {
    stack += fmt.Sprintf("%s -> ", cursor.Value);
  }
  stack += cursor.Value;
  return;
}


//Interfaccia per implementazione del polimorfismo
type ImplementsLen interface {
  Len() int;
}

func AverageLen(lists []ImplementsLen) int {
  var items, sum int = len(lists), 0;
  for i := 0; i < items; i++ {
      sum += lists[i].Len();
  }

  return sum / items;
}


func main() {
  var myList *DListNode;
  fmt.Println("\n---------------- Double Linked  List ----------------");
  myList = myList.AddLast("ciao");
  myList = myList.AddLast("come va");
  myList = myList.AddLast("oioi");
  fmt.Println(myList)
  myList = myList.Remove("come va");
  fmt.Println(myList);
  fmt.Println("Len", myList.Len());

  var myStack *UniqueStack;
  fmt.Println("---------------- Stack ----------------");
  myStack = myStack.Push("ciao");
  myStack = myStack.Push("come va");
  fmt.Println(myStack);
  myStack = myStack.pop();
  fmt.Println(myStack);
  fmt.Println("Len", myStack.Len());

  fmt.Println("---------------- Polymorphism ----------------");
  var collection []ImplementsLen =  []ImplementsLen{myList, myStack};
  fmt.Println("Average lenght", AverageLen(collection));
}
