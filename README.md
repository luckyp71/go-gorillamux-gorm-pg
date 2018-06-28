<h1>RESTful web service using Golang, Gorilla Mux, Gorm, and PostgreSQL with one to many relationship</h1>

<div>
<img src="./resources/go.png" height="360pt" width="350pt"/>
<img src="./resources/golang.png"  height="360pt" width="350pt"/>
</div>


<p>This repository is an example on how to create RESTful web service using Golang, Gorrila Mux, Gorm, and PostgreSQL. The app contains two models i.e. customer and contact. Each customer may have one or more contacts, so that this example will simulate one to many relationship using one of Golang's Object Relational Mappin (ORM) frameworks i.e. Gorm. 

<h3>Prerequisites</h3>
<p>1. <a href="https://golang.org">Go</a></p>
<p>2. <a href="http://www.gorillatoolkit.org/pkg/mux">Gorilla Mux</a></p>
<p>3. <a href="http://gorm.io">GORM</a></p>
<p>4. <a href="https://www.postgresql.org">PostgreSQL</a></p>

<h3>Sample Payload</h3>
<p>1. <a href="./resources/getCustomers.png">Get customers</a></p>
<p>2. <a href="./resources/getCustomersByName.png">Get customer list by name</a></p>
<p>3. <a href="./resources/getCustomerById.png">Get customer by</a></p>
<p>4. <a href="./resources/insertCustomer.png">Insert new customer</a></p>
<p>5. <a href="./resources/updateCustomer.png">Update existing customer</a></p>
<p>6. <a href="./resources/deleteCustomerById.png">Delete customer by id</a></p>
