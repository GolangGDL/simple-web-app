
var data = {}
var queryName = ""


function fetch() {

}

function add(name, price, quantity) {
  axios.post('/add', {
    id: queryName,
    name: name,
    quantity: parseInt(price),
    price: parseInt(quantity)
  })
    .then(function (response) {
      data = response.data
      $('#addName').val('');
      $('#addQuantity').val('');
      $('#addPrice').val('');
      $('#add').modal('toggle');
      console.log(data)
      state(data)
    })
    .catch(function (error) {
      console.log(error);
    })
  queryName = ""
}

function edit(name, price, quantity) {
  axios.post('/edit', {
    name: name,
    quantity: parseInt(price),
    price: parseInt(quantity),
    id: queryName
  })
    .then(function (response) {
      data = response.data
      $('#editName').val('');
      $('#editQuantity').val('');
      $('#editPrice').val('');
      $('#edit').modal('toggle');
      console.log(data)
      state(data)
    })
    .catch(function (error) {
      console.log(error);
    })
  queryName = ""
}

function deleteItem() {
  axios.post('/delete', {
    id: queryName
  })
    .then(function (response) {
      data = response.data
      $('#delete').modal('toggle');
      console.log(data)
      state(data)
    })
    .catch(function (error) {
      console.log(error);
    })
  queryName = ""
}


function state(elementList) {
  $("#tableBody").empty();
  var bodyTable = document.getElementById("tableBody");
  var items = elementList.item
  for (var key in items) {
    var pEdit = paragraphEdit(items[key].name,items[key].price,items[key].quantity)
    var pDelete = paragraphDelete(items[key].name)
    var tableRow = document.createElement("tr");
    var dataNameTable = document.createElement("td");
    var dataPriceTable = document.createElement("td");
    var dataQuantityTable = document.createElement("td");
    var dataEditTable = document.createElement("td");
    var dataDeleteTable = document.createElement("td");
    var nameValue = document.createTextNode(items[key].name);
    dataEditTable.setAttribute("id", items[key].name);
    var priceValue = document.createTextNode(items[key].price);
    var quantityValue = document.createTextNode(items[key].quantity);
    dataNameTable.appendChild(nameValue);
    dataNameTable.setAttribute("id", items[key].name+"-name");
    dataPriceTable.setAttribute("id", items[key].price+"-price");
    dataQuantityTable.setAttribute("id", items[key].quantity+"-quantity");
    dataPriceTable.appendChild(priceValue);
    dataQuantityTable.appendChild(quantityValue);
    dataEditTable.appendChild(pEdit)
    dataDeleteTable.appendChild(pDelete)
    tableRow.appendChild(dataNameTable)
    tableRow.appendChild(dataQuantityTable)
    tableRow.appendChild(dataPriceTable)
    tableRow.appendChild(dataEditTable)
    tableRow.appendChild(dataDeleteTable)
    bodyTable.appendChild(tableRow);
  }
  var dataAddTable = document.createElement("td")
  var pAdd = paragraphAdd()
  dataAddTable.appendChild(pAdd)
  var addNewTableRow = document.createElement("tr");
  addNewTableRow.appendChild(dataAddTable)
  bodyTable.appendChild(addNewTableRow);
  document.getElementById("total").textContent = "Total: " + elementList.total;
}

function paragraphEdit(name,price,quantity) {
  var pEdit = document.createElement("p");
  pEdit.title = "Edit"
  var dataPlacement = document.createAttribute("data-placement");
  dataPlacement.value = "top";
  pEdit.setAttributeNode(dataPlacement);
  var dataToggle = document.createAttribute("data-toggle");
  dataToggle.value = "tooltip";
  pEdit.setAttributeNode(dataToggle);
  var buttonEdit = document.createElement("button");
  buttonEdit.classList.add("btn");
  buttonEdit.classList.add("btn-primary");
  buttonEdit.classList.add("btn-xs");
  var dataTitle = document.createAttribute("data-title");
  dataTitle.value = "Edit";
  buttonEdit.setAttributeNode(dataTitle);
  var idEdit = name+"-"+price+"-"+quantity
  buttonEdit.setAttribute("id", idEdit);
  var onclickbtn = document.createAttribute("onclick");
  onclickbtn.value = "setQuery(this.id); setValueModal(this.id);";
  buttonEdit.setAttributeNode(onclickbtn);
  var dataTogglebtn = document.createAttribute("data-toggle");
  dataTogglebtn.value = "modal";
  buttonEdit.setAttributeNode(dataTogglebtn);
  var dataTarget = document.createAttribute("data-target");
  dataTarget.value = "#edit";
  buttonEdit.setAttributeNode(dataTarget);
  var spanEdit = document.createElement("span");
  spanEdit.classList.add("glyphicon");
  spanEdit.classList.add("glyphicon-pencil");
  buttonEdit.appendChild(spanEdit)
  pEdit.appendChild(buttonEdit)
  return pEdit
}

function paragraphDelete(name) {
  var pDelete = document.createElement("p");
  pDelete.title = "Delete"
  var dataPlacement = document.createAttribute("data-placement");
  dataPlacement.value = "top";
  pDelete.setAttributeNode(dataPlacement);
  var dataToggle = document.createAttribute("data-toggle");
  dataToggle.value = "tooltip";
  pDelete.setAttributeNode(dataToggle);
  var buttonDelete = document.createElement("button");
  buttonDelete.classList.add("btn");
  buttonDelete.classList.add("btn-danger");
  buttonDelete.classList.add("btn-xs");
  var dataTitle = document.createAttribute("data-title");
  dataTitle.value = "Delete";
  buttonDelete.setAttributeNode(dataTitle);
  buttonDelete.setAttribute("id", name);
  var onclickbtn = document.createAttribute("onclick");
  onclickbtn.value = "setQuery(this.id)";
  buttonDelete.setAttributeNode(onclickbtn);
  var dataTogglebtn = document.createAttribute("data-toggle");
  dataTogglebtn.value = "modal";
  buttonDelete.setAttributeNode(dataTogglebtn);
  var dataTarget = document.createAttribute("data-target");
  dataTarget.value = "#delete";
  buttonDelete.setAttributeNode(dataTarget);
  var spanDelete = document.createElement("span");
  spanDelete.classList.add("glyphicon");
  spanDelete.classList.add("glyphicon-trash");
  buttonDelete.appendChild(spanDelete)
  pDelete.appendChild(buttonDelete)
  return pDelete
}


function paragraphAdd() {
  var pAdd = document.createElement("p");
  pAdd.title = "Add"
  var dataPlacement = document.createAttribute("data-placement");
  dataPlacement.value = "top";
  pAdd.setAttributeNode(dataPlacement);
  var dataToggle = document.createAttribute("data-toggle");
  dataToggle.value = "tooltip";
  pAdd.setAttributeNode(dataToggle);
  var buttonAdd = document.createElement("button");
  buttonAdd.classList.add("btn");
  buttonAdd.classList.add("btn-secondary");
  buttonAdd.classList.add("btn-xs");
  var dataTitle = document.createAttribute("data-title");
  dataTitle.value = "Add";
  buttonAdd.setAttributeNode(dataTitle);
  var dataTogglebtn = document.createAttribute("data-toggle");
  dataTogglebtn.value = "modal";
  buttonAdd.setAttributeNode(dataTogglebtn);
  var dataTarget = document.createAttribute("data-target");
  dataTarget.value = "#add";
  buttonAdd.setAttributeNode(dataTarget);
  var spanAdd = document.createElement("span");
  spanAdd.classList.add("glyphicon");
  spanAdd.classList.add("glyphicon-plus");
  buttonAdd.appendChild(spanAdd)
  pAdd.appendChild(buttonAdd)
  return pAdd
}

function setQuery(id) {
 var name =  id.split("-"); 
  queryName = name[0]
}

function setValueModal(id){
  var all =  id.split("-");
  var name = all[0]
  var price = all[1]
  var quantity = all[2]
  var nameRef = name+"-name"
  var quantRef = quantity+"-quantity"
  var priceRef = price+"-price"
  var nameItem = document.getElementById(nameRef).textContent; 
  var quantityItem = document.getElementById(quantRef).textContent; 
  var priceItem = document.getElementById(priceRef).textContent; 
  $('#editName').val(nameItem);
  $('#editQuantity').val(quantityItem);
  $('#editPrice').val(priceItem);
}












