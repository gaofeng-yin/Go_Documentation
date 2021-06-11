# Paginated FizzBuzz

## The Goal

The goal of this challenge is to provide a service that will give the following:
* a footer pagination to allow navigation in a website;
* a list of items to display.

The service should receive the following variables:
* page - the current page;
* per_page - the number of items to display;
* total_pages - total available pages;
* boundaries - how many pages we want to link in the beginning and in the end;
* around - how many pages we want to link before and after the actual page.

For the pages that we don’t want to link we want to show “...”.

Examples:
* current_page = 4; total_pages = 5; boundaries = 1; around = 0: Expected result should be: `1 ... 4 5`
* current_page = 4; total_pages = 10; boundaries = 2; around = 2: - Expected result should be: `1 2 3 4 5 6 ... 9 10`

Regarding the list of products to display, the service should return two fields:
* id - which should start from the current page + 1 and increment it;
* name - for id's multiples of 3, print the word "Fizz" instead of the "item id". For multiples of 5, the word "Buzz" instead of the "item id". For multiples of both 3 and 5, print "BuzzZazz".

The service should return a JSON in the following manner:

```
{
    "metadata": {
        "page": <page>,
        "per_page": <per_page>,
        "total": <total_pages>
    },
    "pagination": <the pagination result>,
    "items": [
        {
            "id": <generated id>,
            "name": <generated name>
        }
    ]
}
```

The amount of items is defined by the `per_page` parameter.
