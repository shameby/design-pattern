## 3．行为型模式

---

行为型模式涉及到算法和对象间职责的分配，行为模式描述了对象和类的模式，以及它们之间的通信模式，行为模式刻划了在程序运行时难以跟踪的复杂的控制流可分为行为类模式和行为对象模式。1.
行为类模式使用继承机制在类间分派行为。2. 行为对象模式使用对象聚合来分配行为。一些行为对象模式描述了一组对等的对象怎样相互协作以完成其中任何一个对象都无法单独完成的任务。

在对象的结构和对象的创建问题都解决了之后，就剩下对象的行为问题了，如果对象的行为设计的好，那么对象的行为就会更清晰，它们之间的协作效率就会提高，这里有11个具体的行为型模式可供研究，它们分别是：

- 模板方法模式（Template Method）
    - 在模板对象中定义算法骨架，子对象可以不改变模板对象算法结构的情况下，重新定义该算法的某些方法
- 观察者模式（Observer）
    - 也叫发布订阅模式，当一个主题对象状态发生改变时，所有依赖它的观察者对象都会被通知到
- 状态模式（State）
    - 对于有状态的对象，把复杂的逻辑判断提取到不同的状态对象中，这些状态对象是无状态的，只会有逻辑，在状态对象内部完成状态改变，从而改变对象行为
- 策略模式（Strategy）
    - 定义多个策略算法，将每个算法封装起来，调用方可以选择策略，与上下文对象结合起来完成特定的功能
- 职责链模式（Chain of Responsibility）
    - 为了避免请求发送者和多个处理者耦合在一起，所以将所有处理者通过前一个对象记住下一个处理者这种引用的方式连城一条链，请求沿着这条链传输，直到有处理者处理为止
- 命令模式（Command）
    - 使请求和执行分离，两者通过命令对象沟通
- 访问者模式（Visitor）
    - 将作用于某个对象结构元素中的操作分离成独立的类，在不改变对象结构的前提下，可以对这些元素添加新的操作，提供多种访问方式，将数据操作和数据结构对象分离，提高扩展性和灵活性
- 调停者模式（Mediator）
    - 也叫中介者模式，通过定义一个中介对象来封装多个同事对象之间的交互，使同事对象间松耦合
- 备忘录模式（Memento）
    - 也叫快照模式，允许调用者捕获一个对象内部状态，并在该对象之外进行保存，在以后需要时可以恢复该状态
- 迭代器模式（Iterator）
    - 提供一个统一的迭代器对象来顺序访问聚合对象中的一系列数据，不用暴露聚合对象内部细节
- 解释器模式（Interpreter）
    - 对于重复多次出现的类似问题，可以将它们抽象为一种简单的语言，并定义该语言的规则，再设计一个解释器来解释语言中的问题
  