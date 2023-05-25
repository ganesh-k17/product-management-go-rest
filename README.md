# product-management-go-rest

Rest API for product management in GO language

## API Endpoints

| HTTP Verb | End Point   | Description                                                                                     |
| --------- | ----------- | ----------------------------------------------------------------------------------------------- |
| GET       | /products   | Gets a list of all the products                                                                 |
| GET       | /product/id | Gets the information about the respective product.                                              |
| POST      | /product    | Creates a new product based on the given information from the user and saves it to the database |
| PUT       | /product/id | Updates the respective product with the given information from the user.                        |
| DELETE    | /product/id | Deletes the respective product                                                                  |
