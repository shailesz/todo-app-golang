const MARK_COMPLETE_URI = "/mark-complete";
const UNMARK_COMPLETE_URI = "/unmark-complete";
const ADD_ITEM_URI = "/add";
const DELETE_ITEM_URI = "/delete";

const form = document.querySelector("#add-item-form");
form.addEventListener("submit", submitForm);

const postData = async (url = "", data = {}) => {
  const response = await fetch(url, {
    method: "POST",
    mode: "cors",
    cache: "no-cache",
    credentials: "same-origin",
    headers: {
      "Content-Type": "application/json",
    },
    redirect: "follow",
    referrerPolicy: "no-referrer",
    body: JSON.stringify(data),
  });
  return response.json();
};

const handleCheckbox = (id) => {
  const item = document.getElementById(id);
  if (item.checked === true) {
    postData(MARK_COMPLETE_URI, { markComplete: [id] }).then((data) => {
      generateToast("Success!", "Your todo item has been marked complete.");
    });
  } else {
    postData(UNMARK_COMPLETE_URI, { unmarkComplete: [id] }).then((data) => {
      generateToast("Success!", "Your todo item has been marked incomplete.");
    });
  }
};

function submitForm(e) {
  e.preventDefault();

  const input = document.querySelector("#item-input-box");
  const value = input.value;

  postData(ADD_ITEM_URI, { description: value }).then((res) => {
    const newItem = `
    <div id="${res.id}" class="input-group w-100">
    <div class="input-group-text rounded-0">
        <input onclick="handleCheckbox("${
          res.id
        }")" class="form-check-input mt-0" type="checkbox" ${
      res.isComplete ? checked : null
    } aria-label="Checkbox for following text input">
    </div>
    
    <li class="list-group-item col">${res.description}</li>
    <span onclick='deleteItem("${res.id}")' class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-danger" style="cursor: pointer;">
    *
    <span class="visually-hidden">delete item</span>
    </span>
    </div>
`;

    const itemList = document.querySelector("#item-list");

    itemList.innerHTML += newItem;

    form.reset();

    generateToast(
      "Success!",
      "Your todo item has been added."
    );
  });
}

const deleteItem = (id) => {
  postData(DELETE_ITEM_URI, { id }).then((res) => {
    const div = document.getElementById(id);

    div.remove();

    generateToast("Success!", "Your todo item has been deleted.");
  });
};

function generateToast(title, message) {
  const toastSection = document.getElementById("toastSection");

  const toastHtml = `
    <div class="toast-container position-fixed top-0 start-0 p-3">
      <div id="toastBox" class="toast" role="alert" aria-live="assertive" aria-atomic="true">
        <div class="toast-header">
          <strong class="me-auto">${title}</strong>
          <button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close"></button>
        </div>
        <div class="toast-body">
          ${message}
        </div>
      </div>
    </div>
    `;

  toastSection.innerHTML = toastHtml;

  const toast = new bootstrap.Toast(document.getElementById("toastBox"));
  toast.show();
}
