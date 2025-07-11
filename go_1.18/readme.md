# Go 1.18前后接口（interface）的变化

Go 1.18 是 Go 语言发展中的一个重大版本，引入了 **泛型（Generics）**，这也直接影响了 `interface` 的定义与使用方式。以下是 Go 1.18 之前与之后在 `interface` 方面的主要变化：

---

### 🔹Go 1.18 **之前** 的 `interface` 特性

1. **传统接口定义**
   接口只用于定义一组方法集合：

   ```go
   type Reader interface {
       Read(p []byte) (n int, err error)
   }
   ```

   实现某个接口的类型必须具有接口中声明的所有方法。

2. **空接口 `interface{}`**
   空接口可以代表任意类型，因为所有类型都实现了空接口：

   ```go
   var x interface{} = 123
   ```

3. **接口不能用于类型约束**（泛型未引入）
   只能通过 `interface{}` + 类型断言实现“泛型”功能，非常不类型安全，代码笨重。

---

### 🔸Go 1.18 **之后** 的 `interface` 变化

1. ✅ **引入了类型参数（泛型）语法**
   使用 `interface` 定义**类型集合（type sets）**，用于约束泛型参数：

   ```go
   type Number interface {
       ~int | ~float64
   }

   func Add[T Number](a, b T) T {
       return a + b
   }
   ```

2. ✅ **新关键字 `any` 替代 `interface{}`**
   `any` 是 `interface{}` 的别名，使代码更语义化：

   ```go
   var v any = "hello" // 等价于 var v interface{}
   ```

3. ✅ **类型约束 interface 可含有类型表达式**
   Go 1.18 的接口不仅可以定义方法，还可以通过 `|` 运算符组合多个类型（即 **联合类型**）：

   ```go
   type IntOrString interface {
       int | string
   }
   ```

4. ✅ **支持类型推导和约束组合**
   类型集合接口可与方法签名并存，用于复杂的泛型约束。

5. ✅ **类型集语义增强**
   接口现在不仅代表“行为（方法）”的集合，还可以代表“类型”的集合。这让 `interface` 成为构建泛型时的核心。

---

### 📌 总结对比

| 特性             | Go 1.17 及以前 | Go 1.18 起                    |              |
| -------------- | ----------- | ---------------------------- | ------------ |
| 接口定义是否支持类型集合   | 否           | ✅ 是                          |              |
| 是否支持泛型         | 否           | ✅ 是                          |              |
| 是否可以限制类型参数     | 否           | ✅ 是，使用 `interface` 类型约束      |              |
| 是否支持 \`        | \` 类型合并     | 否                            | ✅ 是，接口支持联合类型 |
| 是否有 `any` 类型别名 | 否           | ✅ 是 (`any` == `interface{}`) |              |

---

如需示例代码或进一步说明类型推导、约束嵌套等高级用法，也可以告诉我。
