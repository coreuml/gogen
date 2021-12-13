# Gogen (Clean Architecture Code Generator)
Provide code structure based on clean architecure


## Introduction
CLEAN ARCHITECTURE (CA) is a concept of "composing and organizing folder and code files in the very maintainable ways" which has the benefit of separating code logic and infrastructure very neatly so that it is easy to test, easy to mock and easy to switch between technologies with very few changes.

This concept is agnostic against the technology. Means it does not depend on specific programming language. 

If we are googling looking for how the CA structure, then there will be many code examples that implement CA in many programming language. But in general it is just a project template. If we want to apply it, we must imitate it.

I'm trying to go one step further. Instead of copying the project template, why not create a code generator that applies this CA concept. These tools will help bootstrap the code so that developers don't have to struggle with imitating previous projects that aren't necessarily proven.

The process of drafting the concept and making it is also not easy and requires time and coding experience. I had to do some research by reading and studying dozens of articles on the internet about Clean Architecture, Clean Code, Solid Design, Domain Driven Design.

I try to empathize with programmers. Try to feel what they think when they want to write code. For example, "what should I create first? where should I put the controller? what is the proper name for this file?" I try to pour all these feelings and thoughts into this tool. So that it can guide programmers in coding activities.

Some principles I apply are
1. These tools should not be "know-it-all" tools. The programmer should still be the master of design. Because I don't want these tools to drive logic programmers instead. This tool only helps to guide to write standard code templates with clear names and conventions. The rest we still give the programmer space to work.
2. This tool has several alternatives to choose the technology. So if the programmer has better technology or is more familiar, the programmer can easily replace it.
3. I apply the Scream Architecture concept in it so that the generated code can speak for itself to the developers about what their role is and what they are doing (helping the learning process).

Some benefits that can be obtained if you apply this tool are:
1. These tools can become standard in a team. I love innovation and improvisation. However, if innovation and improvisation do not have a clear concept, it is feared that it will mislead the development process and complicate the process of changing or adding requirements in the future.
2. Because it has become a standard, this tool help the communication process between developers QA, project manager, annd product owner, 
3. Help the handover process and knowledge transfer with new programmers because it is easy to learn and imitated.
4. Speed up the code review process and minimize code conflicts during code merges.
5. The code generated results in a readable, simple structure with few directories and a minimum depth that has been calculated very carefully.
6. Facilitate the creation of story cards. a standard structure will help shape the mindset of project managers when making stories. For example every member of developer team can have task per usecase. 

However, this is just a tools. The important things to remember is we must follow the basic principles of clean architecture itself. You may copy and paste the existing code if you think it is more easy. But remember you have to be careful anytime you do that. 

## Video Tutorial how to use it
https://youtu.be/ZqZQGllfbbs

## Sample Apps
- https://github.com/mirzaakhena/userprofile
- https://github.com/mirzaakhena/danarisan
- https://github.com/mirzaakhena/kraicklist
- https://github.com/mirzaakhena/mywallet
- https://github.com/mirzaakhena/yourwishes
- https://github.com/mirzaakhena/gogen (this code itself, just for gimmick purpose)

## Structure
This generator has basic structure like this

```
application/
controller/
domain/
gateway/
infrastructure/
usecase/
main.go
```

## Clean Architecture Concept
The main purpose of this architecture is :
* Separation concern between INFRASTRUCTURE part and LOGIC Part
* Independent of Framework. Free to swap to any framework
* Independent of UI. Free to swap UI. For ex: from Web UI to Console UI
* Independent of Database. Free to swap to any database (data storage)
* Independent of any external agency. Business rule doesn’t know anything about outside world
* Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element


## How to use the gogen?
You always start from creating an usecase, then you continue to create the gateway, then create the controller, and the last step is bind those three (usecase + gateway + controller) in registry part. That's it.

To create the usecase, you need to understand the concept of usecase according to Uncle Bob's Clean Architecture article. Usecase has 3 main part.
* Input Port (Inport)
* Interactor
* Output Port (Outport)

```
Controller is using an Inport.
Inport is implemented by Interactor
Interactor is using an Outport
Outport is implemented by Gateway
```

*Inport* is an interface that has only one method (named `Execute`) that will be called by *Controller*. The method in an interface define all the required (request and response) parameter to run the specific usecase. *Inport* will implemented by *Interactor*. 

*Interactor* is the place that you can define your business process flow by involving entity, valueobject, service and repository. When an usecase logic need a data, it will ask the *Outport* to provide it. *Interactor* also use *Outport* to send data, store data or do some action to other service. *Interactor* only have one *Outport* field. We must not adding new *Outport* field to *Interactor* to keep a simplicity and consistency.

*Outport* is a data and action provider for *Interactor*. *Outport* never know how it is implemented. The implementor (in this case a *Gateway*) will decide how to provide a data or do an action. This *Outport* is very exclusive for specific usecase (in this case *Interactor*) and must not shared to other usecase. By having exclusive *Outport* it will isolate the testing for usecase.

By organize this usecase in a such structure, we can easily change the *Controller*, or the *Gateway* in very flexible way without worry to change the logic part. This is how the logic and infrastructure separation is working.

## Comparison with three layer architecture (Controller -> Service -> Repository) pattern
* *Controller* is the same controller for both architecture.
* *Service* is similar like *Interactor* with additional strict rule. *Service* allowed to have many repositories. In gogen Clean Architecture, *Interactor* only have one *Outport*.
* *Service* have many method grouped by the domain. In Clean Architecture, we focus per usecase. One usecase for One Class to achieve *Single Responsibility Principle*.
* In *Repository* you often see the CRUD pattern. Every developer can added new method if they think they need it. In reality this *Repository* is shared to different *Service* that may not use that method. In *Outport* you will strictly to adding method that guarantee used. Even adding new method or updating existing method will not interfere another usecase.
* *Repository* is an *Outport* with *Gateway* as its implementation.

## Gogen Convention
* Usecase is first class citizen. 
* We always start the code from the usecase.
* Usecase doesn't care about what the technology will be used
* Usecase basically manage interaction between entity, value object, repository and service
* As interface, inport has one and only one method which handle one usecase
* Interactor is a manager for entity and outport
* Interactor must not decide to have any technology. All technology is provided by Outport.
* Outport is a servant for an interactor. It will provide any data from external source
* As interface, Outport at least have one method and can have multiple method
* All method in Outport is guarantee used by interactor
* Inport and Outport must not shared to other usecase. They are exclusive for specific usecase
* Interactor have one and only one outport. Mutiple outport is prohibited
* Entity is mandatory to make a validation for any input. Controller optionaly handle the validation
* Interactor is the one who made a decision mostly based on Entity consideration.
* Entity must not produce unpredictible value like current time, or UUID value. Unpredictible value is provided by Inport or Outport (external Entity)
* Since a log is technology, log only found in Controller and Gateway not in Interactor or Entity
* To avoid a log polution, Log only printing the coming request, leaving response or error response
* Error code can be produced by anyone and will printed in log
* If somehow Gateway produce an error it may log the error, forward back the error,
  forward back the error with new error message or all the possibility
* Error code at least have messaged and code (imitate the http protocol response code)
* Error enum must accessed by the developer and Error code can read by end user
* Interactor and Entity is prioritized to be tested first rather than Controller and Gateway
* Controller name can be an actor name who is using the system
* Registry name can be an application name. You can utilize it if you want to develop microservice with mono repo

## Why you (will) need gogen?
- we want to separate logic code and infrastructure code
- save time because we think less for naming (one of "hardest" think in programming), conventions and structure
- Increase readability, scream architecture
- built for lazy developer
- consistent structure
- gogen is zero dependency. Your code will not have dependency to gogen at all
- gogen is not engine it just a "well written code" so there is no performance issue
- gogen is good templating tools, because deleting is easier than creating right
- gogen support multiple application in one repo
- gogen already implement trace id in every usecase Request
- support lazy documentation. Your interactor is telling everything about how it's work. no need to work twice only to write/update doc
- suitable for new project and revamp existing project per service
- allow you to do code modification for experimental purpose without changing the current implementation
- there is no automagically in gogen. 

## Download it
```
$ go get github.com/mirzaakhena/gogen
```
Install it into your local system (make sure you are in gogen directory)
```
$ go install
```

## Step by step to working with gogen

## Create your basic usecase structure

So you will create your first usecase. Let say the usecase name is  a `CreateOrder`. We will always create our usecase name with `PascalCase`. Now let's try our gogen code generator to create this usecase for us.
```
$ gogen usecase CreateOrder
```

Usecase name will be used as a package name under usecase folder by lowercasing the usecase name.

- `usecase/createorder/inport.go` is an interface with one method that will implement by your usecase. The standart method name is a `Execute`.
- `usecase/createorder/outport.go` is an interface which has many methods that will be used by your usecase. It must not shared to another usecase.
- `usecase/createorder/interactor.go` is the core implementation of the usecase (handle your bussiness application). It implements the method from inport and call the method from outport.

## Create your usecase test file
```
$ gogen test normal CreateOrder
```
normal is the test name and CreateOrder is the usecase name.
This command will help you
- create a test file under `usecase/createorder/testcase_normal_test.go`

## Create a repository 
```
$ gogen repository SaveOrder Order CreateOrder
```
This command will help you 
- create a Repository named `SaveOrderRepo` under `repository/repository.go` (if it not exists yet)
- create an Entity with name `Order` under `domain/entity/order.go` (if it not exists yet)
- inject `repository.SaveOrderRepo` into `usecase/createorder/outport.go` 
- and inject code template into `usecase/createorder/interactor.go`. Injected code will be appear if `//!` is found in interactor's file.

Usecase name in Create a Repository command is optional so you can call it too without injecting it to the usecase
```
$ gogen repository SaveOrder Order
```

## Create a service
```
$ gogen service PublishMessage CreateOrder
```
This command will help you to
- create a Service named `PublishMessageService` under `domain/service/service.go`
- inject `service.PublishMessageService` into `usecase/createorder/outport.go`
- Inject code template into `usecase/createorder/interactor.go`. Injected code will be appear if `//!` is found in interactor's file.

## Create a gateway for your usecase

Gateway is the struct to implement your outport interface. You need to set a name for your gateway. 
In this example we will set name : inmemory
```
$ gogen gateway inmemory CreateOrder
```
This command will read the Outport of `CreateOrder` usecase and implement all the method needed in `gateway/inmemory/gateway.go`

```
$ gogen gateway inmemory
```
This command will read all the usecase under `usecase/` and create all the default implementation needed in one struct


## Create a controller for your usecase

In gogen, we define a controller as trigger of the usecase. 
It can be rest api, grpc, consumer for event handling, or any other source input. 
By default, it only uses gin/gonic web framework. 
Call this command for create a controller. 
`restapi` is your controller name. Controller name can be grouped by client who use the API
```
$ gogen controller restapi CreateOrder gin
```

## Glue your controller, usecase, and gateway together

After generate the usecase, gateway and controller, we need to bind them all by calling this command.
```
$ gogen registry appone restapi
```
appone is the registry name. registry name is an application name. After calling the command, some of those file generated will generate for you in `application/registry`

## Create entity
entity is a mutable object that has an identifier. This command will create new entity struct under `domain/entity/` folder
```
$ gogen entity Order
```

## Create valueobject
valueobject is an immutable object that has no identifier. This command will create new valueobject under `domain/vo/` folder
```
$ gogen valueobject FullName FirstName LastName 
```

## Create valuestring
valuestring is a valueobject simple string type. This command will create new valuestring struct under `domain/vo/` folder
```
$ gogen valuestring OrderID
```

## Create enum
enum is a single immutable value. This command will create new enum struct under `domain/vo/` folder
```
$ gogen enum PaymentMethod DANA Gopay Ovo LinkAja
```

## Create error enum
error enum is a shared error collection. This command will add new error enum line in `application/apperror/error_enum.go` file
```
$ gogen error SomethingGoesWrongError
```

Some Information and FAQ
```
Fact
- all repository is interface
- almost all the repository only have one method
- controller may call multi usecase via inport
- service can be a simple function or an interface
- All method in repository and service have a context in its first params

ideally repository interface has only one method on it
then in what condition an interface can have more than one method?
- when the client of that interface guarantee will call one or more from all the method on it
- sample : repository for transaction that have begin commit and rollback

violation of interface method
- when one of method is guarantee will never called

benefit of simple interface
- easy to mock and test simulation

what is handled by entity
- collection data as object
- action related to that object only
- use other entity
- use value object

what should you only call/do in interactor?
- any method available in Outport (repository/service)
- entity
- value object
- do the simple logic like looping condition checking

what can you put in InportRequest?
- simple builtin data type as field
- we can also expose the func for advanced usage

what can you put in InportResponse?
- we can expose simple builtin data type (more recommended)
- we can also expose our Entity or Value Object

can we use entity as database structure?
- i recommend not to use it to database structure
- create the other structure in gateway is better but you must create a mapper for this

what you can declare in Outport?
- (mostly) repository
- service
- direct method

how to decide whether to use service instead of direct method?
- it is depend on the question "are we will sharing it to other usecase?"
- if we think that we want to share the method to other usecase then put it to the service
- otherwise put it as direct method in outport
- we can promoted direct method to service if it is used by other usecase

what is entity
- mutable object that has an id
- two entity is a same object if the id is similar even all of the field value is different
- entity can be a aggregate for other entity
- entity has its own factory method as constructor

what is value object?
- imutable object that all the field combination as an id
- two value object is a same object if all of the field value is different
- changing single value on any field will make it treated as different object

why we need value object?
- to objectify the value instead of using primitive datatype, for ex:
- address has a street, number, city,
- price has a amount, currency and must not negative (depend on bussiness process)

what will handled by repository?
- anything that related to data storage
- Find, Save, Delete, or Update data
- always passing the entity or primitve type
- make sure not use a loot of parameter

if you find a lot of parameter inf a function/method
then it can be wrapped to the struct

what is handled by service (when we use service?)
- anything that is not related to data storage
- anything that canot be handled by method entity
- when it is hard to find to which Entity that behavior belongs
- contains business invariants that are too complex to be stored in a single Entity or Value Object
- generating id which used in interactor
- publishing message to message broker
- basically it is an interface. But it can be a function that has a repository and entity
- does not hold any state

how to naming the controller?
- we are naming it by the actor who is accessing it
- for example: mobileapi, userapi, restapi, backoffice, consumer, webhook

how to naming the registry?
- we are naming it by your application/service name
- for ex: paymentservice, shippingservice, backofficeservice

how to naming the gateway?
- we are naming it by your enviroment
- you can using different gateway if you want to make a simulation without changing the existing implementation
- for ex: prod (for production), local (for local db), experimental, mockdb, inmemory

how to naming the usecase?
- use simple 2 or 3 word that describe the action

there are two types of usecase
- command (or called an action)
- query (displaying information only without changing the data)

who is using infrastructure?
- controller: log, server, consumer, util, token
- gateway: log, database, publisher, util, token, cache
- registry: log,

make sure to naming bool variable as positif word like : isRunning instead of isNotRunning, isFound instead of isNotFound

when to use panic?
- when you first time initialize the system
- never call it in runtime method otherwise your system will stop and crash suddenly

any repository, service or method in entity at least always return the error
any repository and service method must have context.Context in its first params

You are not allowed to call util directly from entity or service

service may include the repository

avoid init method

follow the rule even you are not understand yet

what is benefit of using the repository/service interface instead of make your own method?
- you can reuse the existing implementation without writing the new one

what if you want to create your own implementation but still using existing repository/service interface?
- you can create your own method on outport
- or you can create your own implementation in gateway by create new gateway

sample case:
- upload image to Google Content Store

We always start from
- Data structure using class diagram database table
- According to DDD


In domain, (entity, vo, service, repository) we will never find dependency to usecase, gateway and controller (outside layer)
entity never know about the technology it used.
so you never mention a sql or json or gorm or something else here
except you adopt external technology that wrapped in your code and never expose it directly

In usecase you mostly mention about entity, vo, service, repository and domerror only
even, no log found here

in gateway and controller, you may mention about entity, vo, 







```

# Reference

https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

https://sparxsystems.com/enterprise_architect_user_guide/14.0/guidebooks/meet_the_business_analysis_tools.html

https://emacsway.github.io/en/service-layer/

https://lostechies.com/jimmybogard/2008/08/21/services-in-domain-driven-design/

https://itqna.net/questions/5017/what-service-layer-ddd

http://gorodinski.com/blog/2012/04/14/services-in-domain-driven-design-ddd/

https://developer20.com/services-in-ddd-finally-explained/

https://github.com/alesimoes/hexagonal-clean-architecture

https://github.com/GSabadini/go-bank-transfer

https://levelup.gitconnected.com/practical-ddd-in-golang-domain-service-4418a1650274

http://web.archive.org/web/20180422210157/http://alistair.cockburn.us/Hexagonal+Architecture

https://dev.to/pereiren/clean-architecture-series-part-3-2795

http://www.plainionist.net/Implementing-Clean-Architecture-Controller-Presenter/

https://blog.ploeh.dk/2013/12/03/layers-onions-ports-adapters-its-all-the-same/

https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/

https://medium.com/@jfeng45/go-microservice-with-clean-architecture-transaction-support-61eb0f886a36

https://medium.com/onfido-tech/designing-gateways-for-greater-good-b6d8340465c7

https://altkomsoftware.pl/en/blog/create-better-code-using-domain-driven-design/

https://www.thereformedprogrammer.net/three-approaches-to-domain-driven-design-with-entity-framework-core/

http://cleancodejava.com/uncle-bob-payroll-case-study-full-implementation/

https://craftsmanshipcounts.com/clean-architecture-a-tale-of-two-stories/
	
https://stackoverflow.com/questions/56700996/if-business-logic-wants-to-send-an-email

https://softwareengineering.stackexchange.com/questions/357052/clean-architecture-use-case-containing-the-presenter-or-returning-data

https://opus.ch/ddd-concepts-and-patterns-service-and-repository/

https://blog.flexiple.com/clean-architecture-build-software-like-an-artisan/

https://enterprisecraftsmanship.com/posts/domain-centric-vs-data-centric-approaches/

https://kevinlawry.wordpress.com/2012/08/07/why-i-avoid-stored-procedures-and-you-should-too/

https://wkrzywiec.medium.com/ports-adapters-architecture-on-example-19cab9e93be7

https://zogface.blog/2018/01/13/thinking-about-microservices-start-with-your-database/

http://www.7loops.com/software-architect-interview-questions/

https://crosp.net/blog/software-architecture/clean-architecture-part-2-the-clean-architecture/

https://vaadin.com/learn/tutorials/ddd/tactical_domain_driven_design

https://www.tutorialspoint.com/software_architecture_design/introduction.htm

https://softwareengineering.stackexchange.com/questions/362071/clean-architecture-too-many-use-case-classes

https://stackoverflow.com/questions/50871171/how-do-you-use-transactions-in-the-clean-architecture

https://github.com/mattia-battiston/clean-architecture-example/issues/1

http://blog.cleancoder.com/uncle-bob/2016/01/04/ALittleArchitecture.html

https://culttt.com/2014/09/29/creating-domain-services/

https://towardsdatascience.com/how-to-implement-domain-driven-design-ddd-in-golang-2e2139beb09d

https://pusher.com/tutorials/clean-architecture-introduction/

https://khalilstemmler.com/articles/software-design-architecture/domain-driven-design-vs-clean-architecture/

https://softwareengineering.stackexchange.com/questions/371966/is-clean-architecture-by-bob-martin-a-rule-of-thumb-for-all-architectures-or-i

http://www.pinte.ro/Blog/DesignPatterns/Clean-Architecture-An-alternative-to-traditional-three-layer-database-centric-applications/37

https://go2goplay.golang.org/p/zfTIGXO02tZ

https://www.youtube.com/watch?v=WcU0EImGI_g

https://www.youtube.com/watch?v=92ZJcxJgmmE