import {httpRequest} from "$lib/http/httpRequest";

export async function login() {
  return httpRequest<any, undefined>("http://localhost:8080/auth/google?provider=google" ).then(response => {
    if (response.ok) {
      return response.data as any;
    }
    throw new Error("Failed to log in");
  });
}