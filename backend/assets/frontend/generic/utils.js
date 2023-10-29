function newOption() {
    let input = document.createElement("input");
    input.setAttribute("type", "text");
    input.setAttribute("class", "manyRelation");

    let div = document.getElementById("inputs");
    div.appendChild(input);
}

function deleteOption() {
    let inputs = document.getElementsByClassName("manyRelation");
    let lastInput = inputs.item(inputs.length - 1);
    lastInput.remove();
}