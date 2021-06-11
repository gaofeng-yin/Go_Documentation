<h3>Solution</h3>
I structured the project in 2 packages, model and service, and the main function. 
<h3>Model</h3>
In the model package You can find the product model with Data, Metadata, and Item struct. Data struct contains all product information, Metadata are page information and Items is a list of product.

<h3>Service</h3>
The service package handles all the business logic.<br>
<br>

The pagination service will print out Json containing all the information requested: metadata, pagination, and items.<br>
`func Paginationservice(page, per_page, total_page, boundaries, around int)`<br>
 - Create productData as product data type and set metadata;
 - Get pagination with pagination function;
 - Get items with productList function;
 - Convert productData to Json ([]byte);
 - Print out in string.

Pagination function will create proper pagination with pages and boundaries.<br>
`func pagination(current_page, total_page, boundaries, around int) string `
 - Create slice of string to store pagination value and initialize counter and limit variables;
 - For loop to get all the values;
 - Convert pagination slice to string and return.

ProductList generate list of items according to page information.<br>
`func productList(current_page, per_page int) []model.Item `
 - Create slice of item to store items;
 - For loop to get all the items;
 - Return list of items.

GenerateName creates name by checking id number.<br>
`func generateName(id int) string `
 - Switch statement to return proper name for each id value. 