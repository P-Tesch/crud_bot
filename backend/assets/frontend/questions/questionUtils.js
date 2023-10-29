function newOptionCheckbox() {
    let input = document.createElement("input");
    input.setAttribute("type", "checkbox");
    input.setAttribute("class", "manyRelationCheckbox");

    let div = document.getElementById("inputs");
    div.appendChild(input);
}

function deleteOptionCheckbox() {
    let inputs = document.getElementsByClassName("manyRelationCheckbox");
    let lastInput = inputs.item(inputs.length - 1);
    lastInput.remove();
}