## 1．创建型模式

---

创建型模式，就是创建对象的模式，抽象了实例化的过程。它帮助一个系统独立于如何创建、组合和表示它的那些对象。关注的是对象的创建，创建型模式将创建对象的过程进行了抽象，也可以理解为将创建对象的过程进行了封装，作为客户程序仅仅需要去使用对象，而不再关系创建对象过程中的逻辑。

社会化的分工越来越细，自然在软件设计方面也是如此，因此对象的创建和对象的使用分开也就成为了必然趋势。因为对象的创建会消耗掉系统的很多资源，所以单独对对象的创建进行研究，从而能够高效地创建对象就是创建型模式要探讨的问题。这里有6个具体的创建型模式可供研究，它们分别是：

- 简单工厂模式（Simple Factory）
    - 有一个具体的工厂方法，可以根据传入的参数，生产不同的产品。
- 工厂方法模式（Factory Method）
    - 调用方只需知道具体工厂名就可以生产对应的产品。
- 抽象工厂模式（Abstract Factory）
    - 可以利用抽象工厂生产多个产品族，不同等级的产品。 
    - 抽象工厂解决的是横向的产品族，工厂方法解决的是纵向的产品等级
- 创建者模式（Builder）
    - 将一个复杂对象的构造和表示分离，使同样的构建过程可以创建不同的产品，产品组成不变，但是每个部分可以灵活选择。
- 原型模式（Prototype）
    - 用一个已经创建的实例作为原型，通过复制该对象来创建一个与原型相同的新对象。
- 单例模式（Singleton）
    - 一个进程唯一一个实例，提供全局访问点。