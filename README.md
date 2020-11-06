## DDD


### Domain Layer
We will consider the domain first.
The domain has several patterns. Some of which are:
*Entity*, *Value*, *Repository*, *Service*, and so on.
 - Defined the API business logic in our domain.
 - Currently we are going to use Entity and Repository

#### Entity
Here we define the Schema of things, the models
#### Repository
The repository defines a collection of methods that the infrastructure implements. 
The methods are defined in an interface. These methods will later be implemented in the infrastructure layer.

---

### Infrastructure Layer
This layer implements the methods defined in the repository. 
The methods interact with the database or a third-party API.
This article will only consider database interaction.

---
### Application Layer
The application connects the domain and the interfaces layers.
- P.e: The UserApp struct has the UserRepository interface, which made it possible to call the user repository methods.

---

### Interfaces Layer
The interfaces is the layer that handles HTTP requests and responses. 
This is where we get incoming requests for authentication, user-related stuff, and food-related stuff.
We can also create middleware in that side.


---

Jose Ramón Mañes