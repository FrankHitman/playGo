# Custom Generic


## Constraints and Type Parameters
A constraint means a type constraint, it is used to constrain some type parameters. 
We could view constraints as types of types. 类型约束，类型的类型

The relation between a constraint and a type parameter is like the relation between a type and a value. 

- 约束 <--> 类型： constraints are type templates (and types are constraint instances)
- 类型 <--> 值： types are value templates (and values are type instances)

Go 1.18内置了any和comparable两个类型约束。any表示任意类型，comparable表示类型的值可以使用==和!=比较大小。

## Tilde form and type union
The ~T form, where T is a type literal or type name. 
T must denote a non-interface type whose underlying type is itself 
(so T may not be a type parameter, which is explained below). 
The form denotes a type set, which include all types whose underlying type is T

T 是非接口类型， 加波浪线之后代表类型的集合。

The T1 | T2 | ... | Tn form, which is called a union of terms

## The difference between the version before and after 1.18

### interface
We know that, before Go 1.18, an interface type may embed
我们知道，在 Go 1.18 之前，接口类型可以嵌入

- arbitrary number of method specifications (method elements, one kind of interface elements);
任意数量的方法规范（方法元素，一种接口元素）;
- arbitrary number of type names (type elements, the other kind of interface elements), but the type names must 
  denote interface types.
任意数量的类型名称（类型元素，另一种接口元素），但类型名称必须表示接口类型。


Go 1.18 relaxed the limitations of type elements, so that now an interface type may embed the following type elements:
Go 1.18 放宽了类型元素的限制，因此现在接口类型可以嵌入以下类型元素：

- any type literals or type names, whether or not they denote interface types, but they must not denote type parameters.
任何类型文本或类型名称，无论它们是否表示接口类型，但不得表示类型参数。
- tilde forms. 波浪号形式。
- term unions. 术语联合。


### Type sets and type implementations
类型集和类型实现

- Before Go 1.18, an interface type is defined as a method set. 接口类型被定义为方法集。
- Since Go 1.18, an interface type is defined as a type set. A type set consists of only non-interface types.
  接口类型被定义为类型集。类型集仅包含非接口类型。

Go 1.18对接口定义语法进行了扩展。 在接口定义中既可以定义接口的方法集(Method Set)，
也可以声明可以被用作泛型类型参数的类型实参的类型集(Type Set)。

```go
type Constraint1 interface {
	T1 // 约束限制为T1类型
}

type Constraint2 interface {
	~T1 // 约束限制为所有底层类型为T1的类型
}

type Constraint3 interface {
	T1 | T2 | T3 // 约束限制为T1, T2, T3中的一个类型
}
```

refer to [enhanced_interface_syntax](enhanced_interface_syntax.go)

```go
type Bytes []byte  // underlying type is []byte
type Letters Bytes // underlying type is []byte
type Blank struct{}
type MyString string // underlying type is string

func (MyString) M() {}
func (Bytes) M() {}
func (Blank) M() {}

// The type set of P only contains one type:
// []byte.
type P interface {[]byte}

// The type set of Q contains
// []byte, Bytes, and Letters.
type Q interface {~[]byte}

// The type set of R contains only two types:
// []byte and string.
type R interface {[]byte | string}

// The type set of S is empty.
type S interface {R; M()}

// The type set of T contains:
// []byte, Bytes, Letters, string, and MyString.
type T interface {~[]byte | ~string}

// The type set of U contains:
// MyString, Bytes, and Blank.
type U interface {M()}

// V <=> P
type V interface {[]byte; any}

// The type set of W contains:
// Bytes and MyString.
type W interface {T; U}

// Z <=> any. Z is a blank interface. Its
// type set contains all non-interface types.
type Z interface {~[]byte | ~string | any}

```

### 从 go1.19开始 可以给非基本类型的接口声明别名，例如：
```go
type C[T any] interface{~int; M() T}
type C1 = C[bool]
type C2 = comparable
type C3 = interface {~[]byte | ~string}
```
以上在 1.19 之前是不支持的

### interface 内部约束之间取交集，可以过滤
```go
type C interface {
	comparable
	[]byte | func() | map[int]bool | string | [2]any
}
```
以上 comparable 接口是严格可比较类型， 过滤掉容器类型：[]byte, map[int]bool, [2]any。 
最终声明的约束 C 只包括类型 string。

现在 1.22 版本 comparable 还被当作非基本interface类型。是 any 的子集。

### strictly comparable vs comparable
strictly comparable 在 Go1.20 版本中被引入，比较 "可以严格比较类型" 可以避免运行时的 Panic 错误。
- 普通类型和 comparable 是 "可严格比较类型"
- interface 和 interface 的组合不是 "可严格比较类型"

The following value types are not strictly comparable:

- (basic) interface types.
- struct types which fields contain interfaces.
- array types which elements contain interfaces.
- type parameters which type set contains at least one type of the above two cases (structs and arrays).


comparable类型，但不是 strictly comparable类型：
- any
- [2]any
- struct{x any}

### type implementation vs type satisfaction
在 Go1.20 之前，"type implementation" 和 "type satisfaction" 是概念等价的。

在 Go1.20之后， type implementation 意思没变。
但是有时候 类型 X 可能满足接口类型 Y，但是却没有实现接口类型 Y。
如果类型 X 实现了接口类型 Y，那么 X 一定能满足接口类型 Y。

例如 any 没有实现类型约束 comparable，但是当它被用作普通值类型时候，它满足 comparable。

```go
type A [2]any

func (a *A) M() {}

type B struct{
	A
	x any
}

type C interface {
	comparable
	M() // interface{ M() }
}
```
以上 *A 和 *B 满足 接口 C，但是它们都没有实现接口 C。因为 *A 和 *B 实现了 interface{ M() }

### Each type parameter is a distinct named type
每一个类型参数都是独一无二的类型，即使声明的内容一样。

a type parameter is just a placeholder for the types in its type set.
类型参数只是这个类型集的一个占位符。
```go
[A int, B ~A]                       // error
type Cx[T int,] interface {         // error
T
}
```


## Reference

- [go-1.18 notes](https://blog.frognew.com/2022/03/go-1.18-notes-04.html)
- [go101](https://go101.org/generics/555-type-constraints-and-parameters.html)
