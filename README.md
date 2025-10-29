# Go Design Patterns

This repository contains practical implementations of classic design patterns in Go. Below you will find detailed explanations and motivations for each pattern implemented so far.

---

## 🧩 Generator

### What is the Generator Pattern?

The Generator Pattern is a behavioral pattern with the following purpose:

👉 “Produce a sequence of values on demand, without needing to store them all in memory, allowing the consumer to process values as they are generated.”

🧠 **In practical terms:**

- Encapsulates a process for generating data (numbers, events, file lines, messages, etc.).
- The consumer reads these values incrementally, usually in a loop.
- Can be synchronous (one value at a time) or asynchronous (values produced in parallel).

📘 **Short (formal) definition:**

The Generator Pattern defines an object or function capable of producing an iterable sequence of values, one at a time, without needing to compute the entire sequence in advance.

💬 **Why is it useful?**

- Memory efficiency: does not load everything at once.
- Performance: processing can start before the entire sequence is generated.
- Natural concurrency (in Go): with goroutines and channels, we can generate and consume data simultaneously.
- Decoupling: producer and consumer are independent.

---

## 🎨 Decorator

### What is the Decorator Pattern?

The Decorator Pattern is a structural pattern with the following purpose:

👉 “Add new responsibilities or behaviors to an object dynamically, without altering its structure or requiring changes to its code.”

🧠 **In practical terms:**

- Allows you to wrap an object with additional functionality (e.g., logging, timing, validation) transparently.
- You can stack multiple decorators to compose behaviors flexibly.
- The original object remains unchanged and unaware of the extra features.

📘 **Short (formal) definition:**

The Decorator Pattern defines a way to attach new responsibilities to objects dynamically by placing them inside special wrapper objects called decorators.

💬 **Why is it useful?**

- Open/Closed Principle: extend behavior without modifying existing code.
- Flexible composition: combine multiple features as needed.
- Reusability: decorators can be reused across different objects.
- Separation of concerns: keeps core logic clean and focused.

---

## 🏭 Simple Factory

### What is the Simple Factory Pattern?

The Simple Factory Pattern is a creational pattern with the following purpose:

👉 “Centralize the creation logic of objects, allowing the client to request an instance by specifying a type, without knowing the details of how each object is constructed.”

🧠 **In practical terms:**

- The client requests an object by passing a type or key (e.g., "email" or "sms") to the factory.
- The factory function decides which concrete type to instantiate and returns it as a common interface.
- The client uses the returned object through the interface, without worrying about its internal construction.
- The client still needs to know the available types to request.

📘 **Short (formal) definition:**

The Simple Factory Pattern centralizes object creation in a single function or type, which returns instances of different concrete types based on input parameters, all conforming to a common interface.

💬 **Why is it useful?**

- Encapsulates instantiation logic, keeping client code cleaner.
- Reduces duplication: object creation code is in one place.
- Makes it easier to add new types: just update the factory.
- Promotes use of interfaces, increasing flexibility and testability.

---

# 🏭 Factory Method

### What is the Factory Method Pattern?

The Factory Method Pattern is a creational pattern with the following purpose:

👉 “Define an interface for creating an object, but let subclasses decide which class to instantiate. Factory Method lets a class defer instantiation to subclasses.”

🧠 **In practical terms:**

- The client works only with factory interfaces and product interfaces, not with concrete types.
- Each concrete factory knows how to create a specific product (or family of products).
- To change the product, the client just switches the factory, not the creation logic or parameters.
- Promotes extensibility: new product types require only new factories, not changes to client code.

📘 **Short (formal) definition:**

The Factory Method Pattern defines an interface for creating objects, allowing subclasses or concrete factories to decide which class to instantiate, thus decoupling object creation from usage.

💬 **Why is it useful?**

- Decouples object creation from usage, following the Dependency Inversion Principle.
- Makes it easy to introduce new product types without changing client code.
- Encourages use of interfaces and abstraction, increasing flexibility and testability.
- Supports the Open/Closed Principle: add new factories/products without modifying existing code.

---

# 🏢 Abstract Factory

### What is the Abstract Factory Pattern?

The Abstract Factory Pattern is a creational pattern with the following purpose:

👉 “Provide an interface for creating families of related or dependent objects without specifying their concrete classes.”

🧠 **In practical terms:**

- Allows the creation of multiple related products (e.g., Notifier and Logger) that are designed to work together.
- The client uses only the abstract factory and product interfaces, never the concrete types.
- Each concrete factory produces a family of products that are compatible with each other.
- Makes it easy to switch entire families of products by changing the factory.

📘 **Short (formal) definition:**

The Abstract Factory Pattern defines an interface for creating families of related objects, letting concrete factories produce objects that belong together, without exposing their concrete classes to the client.

💬 **Why is it useful?**

- Ensures consistency among products that are designed to be used together.
- Decouples client code from concrete classes, increasing flexibility and maintainability.
- Makes it easy to introduce new families of products without changing client code.
- Supports the Open/Closed Principle: add new factories and product families without modifying existing code.

---

# 🏗️ Builder

### What is the Builder Pattern?

The Builder Pattern is a creational pattern with the following purpose:

👉 “Separate the construction of a complex object from its representation, allowing the same construction process to create different representations.”

🧠 **In practical terms:**

- Used when an object has many optional fields or requires step-by-step construction.
- The builder provides methods to set each field or part, often returning itself for method chaining (fluent interface).
- The client can configure only the desired fields, in any order, and then call `Build()` to get the final object.
- Improves code readability and maintainability, especially for objects with many parameters.

📘 **Short (formal) definition:**

The Builder Pattern defines a separate builder object that constructs a complex object step by step, allowing different representations and configurations without a telescoping constructor.

💬 **Why is it useful?**

- Avoids constructors with many parameters (telescoping constructors).
- Makes object creation flexible and readable.
- Allows validation or custom logic during the build process.
- Supports immutability by returning fully constructed objects.

---
