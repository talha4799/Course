// Array Spread Method Iteration
let quarter1 = ["January", "February", "March"];
let quarter2 = ["April", "May", "June"];
let quarter3 = ["July", "August", "September"];
let quarter4 = ["October", "November", "December"];
let full_year = [...quarter1, ...quarter2, ...quarter3, ...quarter4];
console.log(full_year);
// Array with Methods
// let fruits = ["apple", "banana", "orange", "grape", "mango"]
// let fruits_with = fruits.with(2, "kiwi");
// console.log(fruits_with);
// console.log(fruits);
// Array Entries Method
// let fruits = ["apple", "banana", "orange", "grape", "mango"];
// let fruits_entries = fruits.entries();

// for (x of fruits_entries) {
//     txt = "The fruit index and name are: ";
//     print_entry = txt + x;
//     console.log(print_entry);
// }

// console.log(fruits);
// Arrays Keys Method
// let fruits = ["apple", "banana", "orange", "grape", "mango"];
// let fruits_keys = fruits.keys();
// for (let x of fruits_keys) {
//     txt = "The fruits indexes are:";
//     print_index =txt + x;
//     console.log(print_index);
// }

// Array Filter Method
// let numbers = [67, 50, 80, 76, 10, 25, 30, 45, 60, 75];
// let filter_num = numbers.filter(filter_func)

// function filter_func(value, array) {
//     return value > 18;
// }

// console.log(filter_num.sort());

// Array FlatMap Method
// let numbers = [1, 2, 3, 4, 5];
// let num_result = numbers.flatMap((x) => [x * 2]);
// console.log(num_result);
// Array ForEach Method

// let numbers = [10, 20, 30, 40, 50];
// let map = numbers.map(mapFunc);

// function mapFunc(value) {
//     return value * 2;
// }
// console.log(map);
// document.getElementById("map").innerHTML = map;
// let num_find = numbers.forEach(forEachFunc);
// function forEachFunc(value, index, array) {
//     if (value > 35) {
//         return;
//     } 
//     text = "The value is: ";
//     num = text+ value 
//     console.log(num);
//
// }

// console.log(num_find)

// Array Search Methods


// let fruits = ["apple", "banana", "orange", "grape", "mango", "orange"];
// let search_by_index = fruits.lastIndexOf("orange") + 1;
// let search_by_includes = fruits.includes("orange");
// console.log(search_by_includes)

// find
// let numbers = [10, 25, 30, 45, 60, 75];
// let find_num = numbers.findLast(findNumFunc);

// function findNumFunc(value) {
//     return value > 50;
// }

// console.log(find_num); // Output: 60


// sort

// let vegetables = ["carrot", "broccoli", "asparagus", "cauliflower", "corn"];

// vegetables.sort();

// vegetables.reverse();

// console.log(vegetables);

// sort for Number
// let numbers = [40, 100, 1, 5, 25, 10];
// sort_num = compare_num(numbers);

// function compare_num(array) {
//     // return Math.min.apply(null, array)
//     return Math.max.apply(null, array)
//     // return b - a;
//     // return a - b;
// }

// console.log(sort_num);
// Slice

// let fruits = ["apple", "banana", "orange", "grape", "mango"];
// let khata = fruits.slice(2);
// console.log(khata); // Output: ["orange", "grape", "mango"]
// console.log(khata.slice(1)); // Output: ["apple", "banana", "orange", "grape", "mango"]
// console.log(fruits)

// Splice
// fruits = ["apple", "banana", "orange", "grape", "mango"];
// let khata = fruits.toSpliced(2, 1, "kiwi", "pear");
// console.log(khata);