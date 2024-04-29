import {httpRequest} from "$lib/http/httpRequest";

export type Todo = {
  userId: number,
  id: number,
  title: string,
  completed: boolean
}

export async function fetchTodos() {
  return httpRequest<Todo[], undefined>("https://jsonplaceholder.typicode.com/todos").then(response => {
    if (response.ok) {
      console.log(response.data)
      return response.data as Todo[];
    }
    throw new Error("Failed to fetch todos");
  });
}