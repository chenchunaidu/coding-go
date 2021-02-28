
**Interfaces are implemented implicitly**

A type implements an interface by implementing its method. There  is no explicit declaration of intent, no "implements" keyword



**Interface values**

Under the hood, interface values can thought of as a tuple of a value and a concrete type
(value,type)

**Interface values with nil underlying values**

if the concrete value inside the interface itself is nil the method will be called with a nil receiver.