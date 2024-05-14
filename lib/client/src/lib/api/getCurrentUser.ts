import {httpRequest} from "$lib/http/httpRequest";
import {getBaseUrl} from "$lib/http/utils";

type User = {
  ID: string;
  Email: string;
  Firstname: string;
  Lastname: string;
  CreatedAt: Date;
  UpdatedAt: Date;
  Provider: string | null;
  Image: string | null;
}

export async function getCurrentUser() {
  console.log("getting current user")
  const baseUrl = getBaseUrl()
  return httpRequest<User, undefined>(`${baseUrl}/users/session`);
}