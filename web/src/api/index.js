const BASE_URL = "http://10.0.0.147:8080";

const getTodos = () =>
  fetch(BASE_URL + "/api/v1.0/todos/", {
    method: "GET",
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
    }
  }).then(response => {
    return response.json();
  });

const createTodo = () =>
  fetch(BASE_URL + "/api/v1.0/todos/", {
    method: "POST",
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
    }
  }).then(response => {
    return response.json();
  });

const updateTodo = id =>
  fetch(BASE_URL + `/api/v1.0/todos/${id}`, {
    method: "PUT",
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
    }
  }).then(response => {
    return response.json();
  });

const deleteTodo = id =>
  fetch(BASE_URL + `/api/v1.0/todos/${id}`, {
    method: "DELETE",
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
    }
  }).then(response => {
    return response.json();
  });

export { getTodos, createTodo, updateTodo, deleteTodo };
