// Maps

// let fruit = new Map()
// fruit.set("", 900)
// fruit.get()


// WeakMaps
let weakMap = new WeakMap()

let weak_player = {fname: "John", lname: "Smith"}

weakMap.set(weak_player, "player")

let get_weak_player = weakMap.get(weak_player)
console.log(get_weak_player)