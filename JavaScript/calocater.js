 function appendValue(value) {
        document.getElementById("display").value += value;
    }

    function clearDisplay() {
        document.getElementById("display").value = "";
    }

    function calculate() {
        let result = document.getElementById("display").value;
        document.getElementById("display").value = eval(result);
    }