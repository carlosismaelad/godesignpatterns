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
