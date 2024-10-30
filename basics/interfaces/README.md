# Dependency Injection Using Interfaces
Based on https://www.youtube.com/watch?v=McRq-uBAa9I

Lets say there is a function that handles user creation. As the step of this we also need to do some additional things - like send user a notification.

We could write a function that sends a notification and be done with it - but what if we have a new notification service. All this can be handled with dependency injection.

# What is dependency injection? 
When using dependency injection, objects are given their dependencies *at run time rather than compile time
[source](http://blog.gtiwari333.com/2011/05/understanding-dependency-injection-and.html)

OR Dependency injection means giving an object its instance variables. Really. That’s it. [source](https://www.jamesshore.com/v2/blog/2006/dependency-injection-demystified)

# How to do it in GO?

First we need to create structure that will hold our dependencies. It can hold as many depenencies as it needs. 

```go
type UserHandler struct {
	UserNotifier UserNotifier
}
```
Here basically what we say is that UserHandler struct will have access to whatever functionality UserNotifier defines.

Then the funcionality of UserNotifier can be accessed within the methods that are defined on UserHandler struct.

```go
func (u UserHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
    ...
	u.UserNotifier.NotifyUserCreated(user)
    ...
}
```

## Why is this useful? 
So useful parts comes when we want to have a new notifier service. So without this pattern we would need to change 
`u.UserNotifier.NotifyUserCreated(user)` this line over and over.

but now we can just decide on the handler creation which notifer it can use

```go
UserHandler := UserHandler{
    UserNotifier: BetterNotifier{},
}
```
This is useful if we want to have different handlers in different cases. It also makse `handleCreateUser` more flexible and future proof in a sense that if we need a new Notifier definition we know what it will need to do without worrying too much about handleCreateUser. 

