// Array ForEach Method

let numbers = [10, 20, 30, 40, 50];
let map = numbers.map(mapFunc);

function mapFunc(value) {
    return value * 2;
}
console.log(map);
document.getElementById("map").innerHTML = map;
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