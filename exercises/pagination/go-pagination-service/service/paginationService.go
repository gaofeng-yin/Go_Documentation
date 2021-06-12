package service

import (
	"encoding/json"
	"fmt"
	"log"
	"service/pagination/model"
	"strconv"
	"strings"
)

//Pagination service will print out Json containing all the information requested: metadata, pagination and items.
func Paginationservice(page, per_page, total_page, boundaries, around uint) {
	//Data variable to store product information.
	//Metadata is information about product.
	productData := model.Data{
		Metadata: model.Metadata{
			Page:     page,
			Per_page: per_page,
			Total:    total_page,
		},
	}

	//function to get pagination.
	productData.Pagination = pagination(page, total_page, boundaries, around)

	//Function to get items.
	productData.Items = productList(page, per_page)

	//Encode pageData in Json.
	//uncomment next line for formated json
	//dataToJson, err := json.MarshalIndent(data, "", " ")
	dataToJson, err := json.Marshal(productData)
	//In case of error, program ends here and logs the error
	if err != nil {
		log.Fatal(err)
	}

	//Print output in string because Marshal returns in type []byte.
	fmt.Println(string(dataToJson))
}

//Pagination function will create proper pagination with pages and boundaries.
func pagination(current_page, total_page, boundaries, around uint) string {
	//pagination is type strings.builder because we are going to write a variable number of strings.
	//strings.Join for joining and concatenate string when you know exact number of things you are
	//going to concatenate. Link:https://www.reddit.com/r/golang/comments/8vsa54/stringbuilder_is_slower_than_basic_string/
	var pagination strings.Builder
	//Counter to count beginning boundary.
	//For example: boundaries = 3.
	//In this case 1, 2 and 3 will be appended to slice and in each append, counter will be decremented until it reaches to 0.
	countBeginBound := boundaries

	//Counter to count ending boundary.
	//For example: total_page = 14 and boundaries = 3, countEndBound will be 11.
	//In this case 12, 13 and 14 will be appended to slice and in each append, counter will be incremented until it reaches the end.
	countEndBound := total_page - boundaries

	//Set limit around left current page.
	//For example: around = 2 and current_page = 8.
	//AroundLeftLimit will be 6 and append to slice values around these limit.
	AroundLeftLimit := current_page - around

	//Set limit around right current page.
	//For example: around = 2 and current_page = 8.
	//AroundRightLimit will be 10 and append to slice values around these limit.
	AroundRightLimit := current_page + around

	//i is type uint because total page is uint
	for i := uint(1); i <= total_page; i++ {
		switch {
		//while countBeginBound is bigger than 0, append value to slice and decrement counter.
		case countBeginBound > 0:
			//I opt strconv because is much more faster than fmt.Sprint(i), but if you
			//want more readability use Sprint().
			//sprintf vs stringConv: https://gist.github.com/montanaflynn/58df9a48e40baca8dcdc
			pagination.WriteString(strconv.FormatUint(uint64(i), 10))
			pagination.WriteString(" ")
			countBeginBound--
			//check if beginning boundary and around left have intersection.
			//If true, increment countAroundLeft.
			//This will prevent appending same number twice.
			if i == AroundLeftLimit {
				AroundLeftLimit++
			}
			if i > current_page {
				AroundRightLimit--
			}
			if countBeginBound == 0 && i != AroundLeftLimit-1 {
				i = AroundLeftLimit - 1
				pagination.WriteString("... ")

			}
		//when i is at ending boundary, append these values and increment countEndBound.
		case i-1 == countEndBound:
			pagination.WriteString(strconv.FormatUint(uint64(i), 10))
			pagination.WriteString(" ")

			countEndBound++
		//Append current page to slice.
		case i == current_page:
			pagination.WriteString(strconv.FormatUint(uint64(i), 10))
			pagination.WriteString(" ")

		//Append number from left around current page and incremment left limit.
		case AroundLeftLimit != current_page && AroundLeftLimit-i == 0:
			pagination.WriteString(strconv.FormatUint(uint64(i), 10))
			pagination.WriteString(" ")

			AroundLeftLimit++
		//Appende number from right around current page and decrement right limit.
		case AroundRightLimit != current_page && i > current_page:
			pagination.WriteString(strconv.FormatUint(uint64(i), 10))
			pagination.WriteString(" ")

			AroundRightLimit--

			//Jump directly to the end bound.
			if i == current_page+around {
				pagination.WriteString("... ")
				i = countEndBound
			}
		}
	}
	//Return pagination in string.
	return pagination.String()
}

//ProductList generate list of items according to page information.
func productList(current_page, per_page uint) []model.Item {
	//Sclice of item to store items.
	items := []model.Item{}

	//Id start from current page plus 1.
	current_page = current_page + 1
	for i := uint(1); i <= per_page; i++ {
		//Id equal current page.
		id := current_page
		//Name is generated using id by generateName function.
		name := generateName(id)
		//Append new item to slice of items.
		items = append(items, model.Item{Id: id, Name: name})
		//Increment id.
		current_page++
	}
	//Return slice o item.
	return items
}

//GenerateName creates name by checking id number.
func generateName(id uint) string {
	switch {
	//Commom number between multiple of 3 and 5 are multiples of 15.
	case id%15 == 0:
		return "BuzzZazz"
	//Multiple of 3.
	case id%3 == 0:
		return "Fizz"
	//Multiple of 5.
	case id%5 == 0:
		return "Buzz"
	//Return empty string if neither of above condition is valid.
	default:
		return ""
	}
}

